package tensor_test

import (
	"fmt"
	"reflect"

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

// checkTensor compares two tensors against their properties. This function
// returns an error if any of the checks fail.
func checkTensor(a, b *tensor.Tensor) error {
	if as, bs := a.Shape(), b.Shape(); !reflect.DeepEqual(as, bs) {
		return fmt.Errorf("shapes do not match: (%v!=%v)", as, bs)
	}

	return nil
}
