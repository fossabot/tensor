package stale_test

import (
	"fmt"
	"math/cmplx"

	"github.com/ppknap/tensor/stale"
)

func ExampleTensorSoftmax() {
	softmax := func(t *stale.Tensor) *stale.Tensor {
		// Compute exponential of each tensor element.
		sm := t.Clone().Each(cmplx.Exp)

		for _, row := range sm.Split(0) {
			sum := row.Sum()
			row.Each(func(v complex128) complex128 {
				return v / sum
			})
		}

		return sm
	}

	// Sample values.
	vals := []complex128{3.0, 1.0, 0.2}

	fmt.Println(softmax(stale.NewTensor(1, 3).Fill(vals)))
	// Output: [ 0.8360188027814407 0.11314284146556011 0.05083835575299916]
}