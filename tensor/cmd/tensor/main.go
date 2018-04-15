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

	switch ver, err := runPyCmd("version.version"); {
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

var instances = []struct {
	Name    string // Tensor name.
	Tensor  string // Go tensor initialization.
	NDArray string // Python's ndarray initialization.
}{
	{Name: "matrix one element", Tensor: "New(1, 1)", NDArray: "zeros((1, 1))"},
}

var methods = map[string][]struct {
	Name  string // Method name.
	RTyp  string // Function return type.
	Calls []Call // Go to python method call.
}{
	"layout": {{
		Name: "NDim",
		RTyp: "int",
		Calls: []Call{
			{
				Go: "NDim()",
				Py: "ndim",
			},
		},
	}, {
		Name: "Size",
		RTyp: "int",
		Calls: []Call{
			{
				Go: "Size()",
				Py: "size",
			},
		},
	}},
}

// Call describes corresponding Go and Python method calls.
type Call struct {
	Desc string // Additional description.
	Go   string // Call in Go.
	Py   string // Call in Python.
}

// TestCase describes a single test case.
type TestCase struct {
	Name string // Subtest name.
	Expr string // Go expression to call.
	Want string // Wanted result.
}

// TestFull describes all tests for a single method.
type TestFull struct {
	Name  string      // Tested method.
	RTyp  string      // Output type.
	Pass  []*TestCase // Test cases that should pass.
	Panic []*TestCase // Test cases that should panic.
}

func generateTests(group string) (res []TestFull) {
	methods, ok := methods[group]
	if !ok {
		panic("unknown test group: " + group)
	}

	for _, method := range methods {
		var tf = TestFull{
			Name: method.Name,
			RTyp: method.RTyp,
		}

		for _, call := range method.Calls {
			for _, inst := range instances {
				var expr = inst.NDArray + "." + call.Py

				output, err := runPyCmd(expr)
				if err != nil {
					log.Printf("cannot execute %q: %v", expr, err)
					continue
				}

				if _, ok := typeToExpr[method.RTyp]; !ok {
					panic("unknown type: " + method.RTyp)
				}

				tc := &TestCase{
					Name: strings.TrimSpace(inst.Name + " " + call.Desc),
					Expr: inst.Tensor + "." + call.Go,
					Want: typeToExpr[method.RTyp](output),
				}

				if output != "" {
					tf.Pass = append(tf.Pass, tc)
					log.Printf("generated tests for %q, output %q: OK", tc.Expr, tc.Want)
				} else {
					tf.Panic = append(tf.Panic, tc)
					log.Printf("generated tests for %q: PANIC", tc.Expr)
				}
			}
		}

		res = append(res, tf)
	}

	return res
}

func main() {
	tests := generateTests("layout")
	if err := tmpl.Execute(os.Stdout, tests); err != nil {
		panic("template error: " + err.Error())
	}
}

func runPyCmd(op string) (string, error) {
	var (
		code = fmt.Sprintf("import numpy as np\nprint(np.%s)", op)
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

// typeToExpr i a set of functions that format and validate given output to
// Go expression based on expected type.
var typeToExpr = map[string]func(string) string{
	"int": skipEmpty(cmpInt),
}

func skipEmpty(f func(string) string) func(string) string {
	return func(output string) string {
		if output == "" {
			return ""
		}
		return f(output)
	}
}

func cmpInt(output string) string {
	if _, err := strconv.Atoi(output); err == nil {
		return fmt.Sprintf("int(%s)", output)
	}
	panic("invalid integer: " + output)
}

var tmpl = template.Must(template.New("").Parse(`// TODO

package tensor_test

import (
	"testing"
)

{{ range . }}
func TestTensor{{ .Name }}(t *testing.T) {
	tests := map[string]struct {
		Got  {{ .RTyp }}
		Want {{ .RTyp }}
	}{ {{- range .Pass }}
		"{{ .Name }}": {
			Tensor: {{ .Expr }},
			Want: {{ .Want }},
		} {{ end }}
	}

	for name, test := range tests {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v; got %v", want, got)
		}
	}
}
{{ end }}
`))
