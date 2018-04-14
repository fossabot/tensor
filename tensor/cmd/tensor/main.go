package main

import (
	"bytes"
	"fmt"
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
		fmt.Printf("Numpy %s will be used to generate test cases.\n", ver)
	default:
		panic(fmt.Sprintf("invalid numpy version %s (want: %s)", ver, numpyVer))
	}
}

var instances = []struct {
	Name string // Tensor name.
	GoT  string // Go tensor initialization.
	PyA  string // Python's ndarray initialization.
}{
	{Name: "matrix one element", GoT: "New(1, 1)", PyA: "zeros((1, 1))"},
}

var methods = map[string][]struct {
	Name  string            // Method name.
	RTyp  string            // Function return type.
	Cases map[string]string // Go to python method call.
}{
	"layout": {{
		Name: "NDim",
		RTyp: "int",
		Cases: map[string]string{
			"NDim()": "ndim",
		},
	}, {
		Name: "Size",
		RTyp: "int",
		Cases: map[string]string{
			"Size()": "size",
		},
	}},
}

// TestCase describes a single test case.
type TestCase struct {
	Name   string // Subtest name.
	Tensor string // Tensor built from instances.
	Expr   string // Go expression to call.
	Want   string // Wanted result.
}

// TestFull describes all tests for a single method.
type TestFull struct {
	Name  string // Tested method.
	RTyp  string
	Pass  []TestCase // Test cases that should pass.
	Panic []TestCase // Test cases that should panic.
}

func generateTests(group string) (res []TestFull) {
	// methods, ok := methods[group]
	// if !ok {
	// 	panic("unknown test group: " + group)
	// }

	// for i, method := range methods {
	// 	tf := TestFull{Name: method.Name}

	// }

	return res
}

func main() {
	for _, methods := range methods {
		var data = struct{ Is, Ms interface{} }{
			Is: instances, Ms: methods,
		}

		fmt.Println(tmpl.Execute(os.Stdout, data))
	}
}

func generator() {

	_ = instances
	_ = methods
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

{{ range .Ms }}
func TestTensor{{ .Name }}(t *testing.T) {
	tests := map[string]struct {
		Tensor *tensor.Tensor
		Want   {{ .RTyp }}
	}{ {{- range $.Is }}
		"{{ .Name }}": {
			Tensor: {{ .GoT }}
			Want:
		} {{ end }}
	}

	for name, test := range tests {
		if got := test.{{ .GoM }};
	}
}
{{ end }}
`))
