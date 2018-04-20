package main

import (
	"fmt"
	"log"
	"strings"
)

// test describes all test cases for a single method.
type test struct {
	Name  string // Tested method.
	RTyp  string // Output type.
	Pass  []*cas // Test cases that should pass.
	Panic []*cas // Test cases that should panic.
}

func newTest(method *method, instances []*instance) *test {
	tf := &test{
		Name: method.Name,
		RTyp: method.RTyp,
	}

	for _, call := range method.Calls {
		for _, inst := range instances {
			tc, isPanic, err := newTestCas(method.RTyp, inst, call)
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

// cas describes a single test case.
type cas struct {
	Name string // Subtest name.
	Expr string // Go expression to call.
	Want string // Wanted result.
}

func newTestCas(typ string, inst *instance, call *call) (*cas, bool, error) {
	var expr = strings.Replace(call.Py, "$inst$", inst.NDArray, -1)

	output, err := execPythonCmd(typ, expr)
	if err != nil {
		return nil, false, fmt.Errorf("cannot execute %q: %v", expr, err)
	}

	if _, ok := typeToExpr[typ]; !ok {
		log.Fatalf("unknown type: " + typ)
	}

	return &cas{
		Name: strings.TrimSpace(inst.Name + " " + call.Dsc),
		Expr: strings.Replace(call.Go, "$inst$", inst.Tensor, -1),
		Want: typeToExpr[typ](output),
	}, output == "", nil
}
