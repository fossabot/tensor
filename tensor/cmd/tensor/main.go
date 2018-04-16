package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
)

func init() {
	// Check if numpy is available in given version.
	const numpyVer = "1.14"

	switch ver, err := runPyCmd("np.version.version"); {
	case ver == "":
		panic("cannot check numpy version")
	case err != nil:
		panic(err)
	case strings.HasPrefix(ver, numpyVer):
		log.Printf("numpy %s will be used to generate test cases.\n", ver)
	default:
		panic(fmt.Sprintf("invalid numpy version %s (want: %s)", ver, numpyVer))
	}
}

type Instance struct {
	Name    string // Tensor name.
	Tensor  string // Go tensor initialization.
	NDArray string // Python's ndarray initialization.
}

var instances = []*Instance{
	&Instance{
		Name:    "matrix one element",
		Tensor:  "tensor.New(1, 1)",
		NDArray: "np.zeros((1, 1))",
	},
}

type Method struct {
	Name  string  // Method name.
	RTyp  string  // Function return type.
	Calls []*Call // Go to python method call.
}

var methods = map[string][]*Method{
	"layout": {{
		Name: "NDim",
		RTyp: "int",
		Calls: []*Call{
			&Call{
				Go: "$inst$.NDim()",
				Py: "$inst$.ndim",
			},
			&Call{
				Dsc: "two",
				Go:  "$inst$.NDim()",
				Py:  "$inst$.ndim",
			},
		},
	}, {
		Name: "Size",
		RTyp: "int",
		Calls: []*Call{
			&Call{
				Go: "$inst$.Size()",
				Py: "$inst$.size",
			},
		},
	}},
}

// Call describes corresponding Go and Python method calls.
type Call struct {
	Dsc string // Additional description.
	Go  string // Call in Go.
	Py  string // Call in Python.
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
	methods, ok := methods["layout"]
	if !ok {
		panic("unknown test group: " + "group")
	}

	tests := makeTests(methods, instances)
	if err := tmpl.Execute(os.Stdout, tests); err != nil {
		panic("template error: " + err.Error())
	}
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
		panic("unknown type: " + typ)
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
	"int": skipEmpty(func(output string) string {
		if _, err := strconv.Atoi(output); err == nil {
			return output
		}
		panic("invalid integer: " + output)
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

var tmpl = template.Must(template.New("").Funcs(funcMap).Parse(`// TODO

package tensor_test

import (
	"testing"
){{ range . }}{{ if .Pass }}

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
			if {{ makeComparator .RTyp }} {
				t.Errorf("want %v; got %v", want, got)
			}
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
	// makeComparator creates a testing expression in test.
	"makeComparator": func(t string) string {
		switch t {
		case "int":
			return "want != got"
		default:
			return "reflect.DeepEqual(want, got)"
		}
	},
}
