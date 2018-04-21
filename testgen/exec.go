package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func execPythonCmd(typ, op string) (string, error) {
	op, err := prepareOp(typ, op)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("python3", "-c", "import numpy as np\n"+op)

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

func prepareOp(typ, op string) (string, error) {
	switch typ {
	case "*tensor.Tensor":
		return "tmp = " + op + "; print('nil') if tmp is None else print(tmp.shape)", nil
	default:
		// If op does not print anything, print the whole statement result.
		return "print(" + op + ")", nil
	}
}
