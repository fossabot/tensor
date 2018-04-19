package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func execPythonCmd(op string) (string, error) {
	// If op does not print anything, print the whole statement result.
	if !strings.Contains(op, "print(") {
		op = "print(" + op + ")"
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
