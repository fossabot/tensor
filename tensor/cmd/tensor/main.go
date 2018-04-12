package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	generator()
	fmt.Println(">>", runPyCmd("version.version"))
}

func generator() {
	instances := []struct {
		GoT string // Go tensor initialization.
		PyA string // Python's ndarray initialization.
	}{
		{
			GoT: "NewTensor(1, 1)", PyA: "zeros((1, 1))",
		},
	}

	methods := map[string]struct {
		Name string                     // Test name.
		GoM  string                     // Go method to call.
		PyM  string                     // Python method to call.
		Cmp  func(output string) string // How to compare outputs.
	}{
		"layout": {
			GoM: "NDim()", PyM: "ndim", Cmp: nil,
		},
	}

	_ = instances
	_ = methods
}

func runPyCmd(op string) string {
	var (
		code = fmt.Sprintf("import numpy as np\nprint(np.%s)", op)
		cmd  = exec.Command("python3", "-c", code)
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok && !bytes.Contains(out, []byte("SyntaxError")) {
			return ""
		}

		panic(fmt.Sprintf("cannot run %q: %v (output %q)", code, err, out))
	}

	return string(out)
}

func cmpInt(output string) string {
	if _, err := strconv.Atoi(output); err == nil {
		return fmt.Sprintf("int(%s)", output)
	}
	panic("invalid integer: " + output)
}
