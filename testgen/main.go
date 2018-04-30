package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// Check if numpy is available in given version.
	const numpyVer = "1.14"

	switch ver, err := execPythonCmd("", "np.version.version"); {
	case ver == "":
		log.Fatal("cannot check numpy version")
	case err != nil:
		log.Fatal(err)
	case strings.HasPrefix(ver, numpyVer):
		log.Printf("numpy %s will be used to generate test cases", ver)
	default:
		log.Fatalf("invalid numpy version %s (want: %s)", ver, numpyVer)
	}
}

func main() {
	mtsPattern := flag.String("methods", "*.json", "pattern for methods definition JSON files")
	insPath := flag.String("instances", "instance.json", "tensor instances definition file")
	outPath := flag.String("out", ".", "generated files destination directory")

	flag.Parse()

	ins, err := loadInstances(*insPath)
	if err != nil {
		log.Fatalf("cannot load test instances: %v", err)
	}

	match, err := filepath.Glob(*mtsPattern)
	if err != nil {
		log.Fatalf("cannot get methods directory: %v", err)
	}

	for _, m := range match {
		name := strings.TrimSuffix(filepath.Base(m), ".json")
		outFilepath := filepath.Join(*outPath, name+"_test.go")

		mts, err := loadMethods(m)
		if err != nil {
			log.Fatalf("cannot load methods from %q: %v", m, err)
		}

		if err := generate(outFilepath, mts, ins); err != nil {
			log.Fatalf("cannot generate tests from %v: %v", m, err)
		}

		log.Printf("generated tests from %v were saved in %v", m, outFilepath)
	}

	log.Println("done!")
}

func generate(outFilepath string, mts []*method, ins []*instance) error {
	type idxTest struct {
		i int
		t *test
	}

	testC := make(chan idxTest)
	for i, mt := range mts {
		i, mt := i, mt // Capture range variables.
		go func() {
			testC <- idxTest{
				i: i,
				t: newTest(mt, ins),
			}
		}()
	}

	tests := make([]*test, len(mts))
	for i := 0; i < len(mts); i++ {
		test := <-testC
		tests[test.i] = test.t
	}

	close(testC)

	f, err := os.Create(outFilepath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, tests)
}
