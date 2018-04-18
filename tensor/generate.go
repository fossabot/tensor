// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func init() {
	// Check if numpy is available in given version.
	const numpyVer = "1.14"

	switch ver, err := runPyCmd("np.version.version"); {
	case ver == "":
		log.Fatal("cannot check numpy version")
	case err != nil:
		log.Fatal(err)
	case strings.HasPrefix(ver, numpyVer):
		log.Printf("numpy %s will be used to generate test cases.\n", ver)
	default:
		log.Fatalf("invalid numpy version %s (want: %s)", ver, numpyVer)
	}
}

// Instance represents an object on which tests are called.
type Instance struct {
	Name    string `json:"name"`    // Tensor name.
	Tensor  string `json:"tensor"`  // Go tensor initialization.
	NDArray string `json:"ndarray"` // Python's ndarray initialization.
}

// Method describes the call on a single instance.
type Method struct {
	Name  string  `json:"name"`  // Method name.
	RTyp  string  `json:"rtyp"`  // Function return type.
	Calls []*Call `json:"calls"` // Go to python method call.
}

// Call describes corresponding Go and Python method calls.
type Call struct {
	Dsc string `json:"dsc"` // Additional description.
	Go  string `json:"go"`  // Call in Go.
	Py  string `json:"py"`  // Call in Python.
}

// TestFull describes all tests for a single method.
type TestFull struct {
	Name  string      // Tested method.
	RTyp  string      // Output type.
	Pass  []*TestCase // Test cases that should pass.
	Panic []*TestCase // Test cases that should panic.
}

// TestCase describes a single test case.
type TestCase struct {
	Name string // Subtest name.
	Expr string // Go expression to call.
	Want string // Wanted result.
}

func main() {
	match, err := filepath.Glob("testdata/*.json")
	if err != nil {
		log.Fatal(err)
	}

	instances := getInstances()
	for _, m := range match {
		tfs := makeTests(getMethods(m), instances)

		name := strings.TrimSuffix(filepath.Base(m), ".json") + "_test.go"
		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}

		if err := tmpl.Execute(f, tfs); err != nil {
			f.Close()
			log.Fatalf("template error: %v", err)
		}
		f.Close()

		log.Printf("created tests for %v and saved in %v", m, name)
	}

	log.Println("done!")
}

func getInstances() (instances []*Instance) {
	f, err := os.Open(filepath.Join("testdata", "instances", "instances.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&instances); err != nil {
		log.Fatal(err)
	}

	return instances
}

func getMethods(fullpath string) (methods []*Method) {
	f, err := os.Open(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&methods); err != nil {
		log.Fatal(err)
	}

	return methods
}

func makeTests(methods []*Method, instances []*Instance) (res []*TestFull) {
	for _, method := range methods {
		res = append(res, newTestFull(method, instances))
	}

	return res
}

func newTestFull(method *Method, instances []*Instance) *TestFull {
	tf := &TestFull{
		Name: method.Name,
		RTyp: method.RTyp,
	}

	for _, call := range method.Calls {
		for _, inst := range instances {
			tc, isPanic, err := newTestCase(method.RTyp, inst, call)
			if err != nil {
				log.Println(err)
				continue
			}

			if !isPanic {
				tf.Pass = append(tf.Pass, tc)
				log.Printf("generated tests for %q, output %q: OK", tc.Expr, tc.Want)
			} else {
				tf.Panic = append(tf.Panic, tc)
				log.Printf("generated tests for %q: PANIC", tc.Expr)
			}
		}
	}

	return tf
}

func newTestCase(typ string, inst *Instance, call *Call) (*TestCase, bool, error) {
	var expr = strings.Replace(call.Py, "$inst$", inst.NDArray, -1)

	output, err := runPyCmd(expr)
	if err != nil {
		return nil, false, fmt.Errorf("cannot execute %q: %v", expr, err)
	}

	if _, ok := typeToExpr[typ]; !ok {
		log.Fatalf("unknown type: " + typ)
	}

	return &TestCase{
		Name: strings.TrimSpace(inst.Name + " " + call.Dsc),
		Expr: strings.Replace(call.Go, "$inst$", inst.Tensor, -1),
		Want: typeToExpr[typ](output),
	}, output == "", nil
}

func runPyCmd(op string) (string, error) {
	var (
		code = fmt.Sprintf("import numpy as np\nprint(%s)", op)
		cmd  = exec.Command("python3", "-c", code)
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok && !bytes.Contains(out, []byte("SyntaxError")) {
			return "", nil
		}

		return "", fmt.Errorf("cannot run %q: %v (output %q)", code, err, out)
	}

	return string(bytes.TrimSpace(out)), nil
}

var typeToExpr = map[string]func(string) string{
	"bool": skipEmpty(func(output string) string {
		if b, err := strconv.ParseBool(output); err == nil {
			return fmt.Sprintf("%t", b)
		}
		panic("invalid boolean: " + output)
	}),
	"int": skipEmpty(func(output string) string {
		if _, err := strconv.Atoi(output); err == nil {
			return output
		}
		panic("invalid integer: " + output)
	}),
	"[]int": skipEmpty(func(output string) string {
		toks := strings.Split(strings.Trim(output, "(,)"), ", ")
		if len(toks) == 1 && toks[0] == "" {
			return "nil"
		}

		for _, tok := range toks {
			if _, err := strconv.Atoi(tok); err != nil {
				panic("invalid integer: " + output)
			}
		}

		return "[]int{" + strings.Join(toks, ", ") + "}"
	}),
}

// typeToExpr i a set of functions that format and validate given output to
// Go expression based on expected type.
func skipEmpty(f func(string) string) func(string) string {
	return func(output string) string {
		if output == "" {
			return ""
		}
		return f(output)
	}
}

var tmpl = template.Must(template.New("").Funcs(funcMap).Parse(`// Code generated by generate.go. DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs/tensor"
)

// Use reflect package in case it isn't used in tests.
var _ = reflect.TypeOf(&tensor.Tensor{}){{ range . }}{{ if .Pass }}

func TestTensor{{ .Name }}(t *testing.T) {
	tests := map[string]struct {
		Got, Want {{ .RTyp }}
	}{ {{- range .Pass }}
		"{{ .Name }}": {
			Got:  {{ .Expr }},
			Want: {{ .Want }},
		},{{ end }}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			{{ range $expr, $msg := makeComparators .RTyp }}if {{ $expr }} {
				t.Errorf({{ $msg }})
			}{{ end }}
		})
	}
}{{ end }}{{ if .Panic }}

func TestTensorPanic{{ .Name }}(t *testing.T) {
	tests := map[string]func() {
		{{- range .Panic }}
		"{{ .Name }}": func() { _ = {{ .Expr }} },{{ end }}
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			fn()
		})
	}
}{{ end }}{{ end }}
`))

var funcMap = template.FuncMap{
	// makeComparators creates a testing expressions based on tested type.
	"makeComparators": func(t string) map[string]string {
		switch t {
		case "bool", "int":
			return map[string]string{
				"test.Want != test.Got": `"want %v; got %v", test.Want, test.Got`,
			}
		default:
			return map[string]string{
				"reflect.DeepEqual(test.Want, test.Got)": `"want %v; got %v", test.Want, test.Got`,
			}
		}
	},
}
