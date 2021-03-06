package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

	now := time.Now()

	ins, err := loadInstances(*insPath)
	if err != nil {
		log.Fatalf("cannot load test instances: %v", err)
	}

	match, err := filepath.Glob(*mtsPattern)
	if err != nil {
		log.Fatalf("cannot get methods directory: %v", err)
	}

	const parallelFilesN = 4
	ticketC := make(chan struct{}, parallelFilesN)

	for i := 0; i < parallelFilesN; i++ {
		ticketC <- struct{}{}
	}

	var wg sync.WaitGroup
	wg.Add(len(match))

	var nAllPass, nAllPanic int64
	for _, m := range match {
		m := m // Capture range variables.
		go func() {
			defer wg.Done()

			<-ticketC

			name := strings.TrimSuffix(filepath.Base(m), ".json")
			outFilepath := filepath.Join(*outPath, name+"_test.go")

			mts, err := loadMethods(m)
			if err != nil {
				log.Fatalf("cannot load methods from %q: %v", m, err)
			}

			nPass, nPanic, err := generate(outFilepath, mts, ins)
			if err != nil {
				log.Fatalf("cannot generate tests from %v: %v", m, err)
			}

			ticketC <- struct{}{}

			atomic.AddInt64(&nAllPass, nPass)
			atomic.AddInt64(&nAllPanic, nPanic)

			log.Printf("generated tests from %v were saved in %v", m, outFilepath)
		}()
	}

	wg.Wait()

	log.Println("done!")
	log.Printf("generated %d tests in %v, %d of them are panic handlers",
		nAllPass+nAllPanic, time.Since(now), nAllPanic)
}

func generate(outFilepath string, mts []*method, ins []*instance) (nPass, nPanic int64, err error) {
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
		nPass += int64(len(test.t.Pass))
		nPanic += int64(len(test.t.Panic))
	}

	close(testC)

	f, err := os.Create(outFilepath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	return nPass, nPanic, tmpl.Execute(f, tests)
}
