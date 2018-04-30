package tensor_test

import (
	"github.com/ppknap/tensor"
)

// enumerate is a simplified Arange method which assigns unique and predictable
// value to each element in a called tensor.
func enumerate(t *tensor.Tensor) *tensor.Tensor {
	var i int
	t.Each(func(pos []int, _ *tensor.Tensor) {
		t.ItemSet(tensor.NewScalar(i), pos...)
		i++
	})

	return t
}
