package tacvs

import (
	"math/rand"
	"time"
)

// Zeros sets each tensor element to its zero value.
func (t *Tensor) Zeros() *Tensor {
	// TODO: tests
	return t.Full(0)
}

// Ones fills all elements of the tensor with ones.
func (t *Tensor) Ones() *Tensor {
	// TODO: tests
	return t.Full(1)
}

// Full replaces all tensor elements with a given constant.
func (t *Tensor) Full(val complex128) *Tensor {
	// TODO: tests
	for i := range t.data {
		t.data[i] = val
	}

	return t
}

// Arrange creates evenly spaced values with respect to tensor index. Start
// element will be set on zero index and the value will be increased by step
// until last index is reached.
func (t *Tensor) Arrange(start, step complex128) *Tensor {
	// TODO: tests
	return nil
}

// Linspace creates evenly spaced values with respect to tensor index. The first
// element of the tensor will have start value and the last element will be set
// to end value.
func (t *Tensor) Linspace(start, end complex128) *Tensor {
	// TODO: tests
	return nil
}

// Eye sets one to tensor elements with equal indexes on each axe. The remaining
// elements will be set to zero.
func (t *Tensor) Eye() *Tensor {
	// TODO: tests
	return t.Apply(func(t *Tensor, idx []int) complex128 {
		var first = idx[0] // len(idx) is always > 0.
		for i := 1; i < len(idx); i++ {
			if first != idx[i] {
				return 0
			}
		}

		return 1
	})
}

var defaultRandSource = rand.NewSource(time.Now().UnixNano())

// Random sets each element's real part value to a pseudo-random number in
// [0.0,1.0) range using provided random source. If the provided argument is
// nil, a new random source will be created and used.
func (t *Tensor) Random(source rand.Source) *Tensor {
	// TODO: tests
	if source == nil {
		source = defaultRandSource
	}

	var r = rand.New(source)
	for i := range t.data {
		t.data[i] = complex(r.Float64(), 0)
	}

	return t
}

// Re removes imaginary part value from each element of a tensor.
func (t *Tensor) Re() *Tensor {
	// TODO: tests
	for i := range t.data {
		t.data[i] = complex(real(t.data[i]), 0)
	}

	return t
}

// Im removes real part value from each element of a tensor.
func (t *Tensor) Im() *Tensor {
	// TODO: tests
	for i := range t.data {
		t.data[i] = complex(0, imag(t.data[i]))
	}

	return t
}

// Apply iterates over all tensor elements and calls f. The returned value will
// be set at given tensor index. The index order is preserved.
func (t *Tensor) Apply(f func(t *Tensor, idx []int) complex128) *Tensor {
	// TODO: tests
	return nil
}

// Fill copies provided buffer to tensor in column-wise order. The size of
// provided slice must be identical as the size of the tensor.
func (t *Tensor) Fill(vs []complex128) *Tensor {
	if len(vs) != len(t.data) {
		panic("fill with a buffer of invalid size")
	}

	for i := range t.data {
		t.data[i] = vs[i]
	}

	return t
}
