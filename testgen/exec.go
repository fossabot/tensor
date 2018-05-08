package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func execPythonCmd(typ, op string) (string, error) {
	op = prepareOp(typ, op)

	cmd := exec.Command("python3", "-c", "import numpy as np\nnp.seterr(divide='ignore', invalid='ignore')\n"+op)

	out, err := cmd.CombinedOutput()
	if err == nil {
		return string(bytes.TrimSpace(out)), nil
	}

	_, ok := err.(*exec.ExitError)
	if ok && !bytes.Contains(out, []byte("SyntaxError")) {
		return "", nil
	}

	return "", fmt.Errorf("cannot run %q: %v (output %q)", op, err, out)
}

func prepareOp(typ, op string) string {
	switch typ {
	case "*tensor.Tensor":
		return "tmp = " + op + "; print('nil') if tmp is None else print(tmp.shape)"
	default:
		// If op does not print anything, print the whole statement result.
		return "print(" + op + ")"
	}
}
