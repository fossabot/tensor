package tacvs

import (
	"math/rand"
	"time"
)

// Zeros sets each tensor element to its zero value.
func (t *Tensor) Zeros() *Tensor {
	return t.Full(0)
}

// Ones fills all elements of the tensor with ones.
func (t *Tensor) Ones() *Tensor {
	return t.Full(1)
}

// Full replaces all tensor elements with a given constant.
func (t *Tensor) Full(val complex128) *Tensor {
	return t.Each(func(complex128) complex128 { return val })
}

// Each walks trough each tensor element and applies a provided function on it.
// The argument passed to f will be a current element's value.
func (t *Tensor) Each(f func(complex128) complex128) *Tensor {
	if f == nil {
		panic("tensor: each with nil predicate")
	}

	if t.parent == nil {
		for i := range t.data {
			t.data[i] = f(t.data[i])
		}
		return t
	}

	t.Apply(func(t *Tensor, idx []int) {
		t.Set(f(t.At(idx...)), idx...)
	})

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
	return t.Apply(func(t *Tensor, idx []int) {
		// Special case for vectors.
		if len(idx) == 1 {
			if idx[0] == 0 {
				t.Set(1, idx...)
				return
			}

			t.Set(0, idx...)
			return
		}

		var first = idx[0] // len(idx) is always > 0.
		for i := 1; i < len(idx); i++ {
			if first != idx[i] {
				t.Set(0, idx...)
				return
			}
		}

		t.Set(1, idx...)
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
	t.Each(func(val complex128) complex128 {
		return complex(real(val), 0)
	})

	return t
}

// Im removes real part value from each element of a tensor.
func (t *Tensor) Im() *Tensor {
	t.Each(func(val complex128) complex128 {
		return complex(0, imag(val))
	})

	return t
}

// Apply iterates over all tensor elements and calls f. The index order is
// preserved.
func (t *Tensor) Apply(f func(*Tensor, []int)) *Tensor {
	if f == nil {
		panic("tensor: apply with nil predicate")
	}

	if len(t.shape) == 0 {
		return t
	}

	mul, size := make([]int, len(t.shape)), 1

	for i := len(t.shape) - 1; i >= 0; i-- {
		size *= t.shape[i]
		mul[i] = size
	}
	mul = append(mul, 1)

	cur := make([]int, len(t.shape))
	for i := 0; i < size; i++ {
		for j := range cur {
			cur[j] = (i / mul[j+1]) % t.shape[j]
		}
		f(t, cur)
	}

	return t
}

// Fill copies provided buffer to tensor in column-wise order. The size of
// provided slice must be identical as the size of the tensor. It panics if
// called on views.
func (t *Tensor) Fill(vs []complex128) *Tensor {
	if t.parent != nil {
		panic("fill called on view tensor")
	}

	if len(vs) != len(t.data) {
		panic("fill with a buffer of invalid size")
	}

	for i := range t.data {
		t.data[i] = vs[i]
	}

	return t
}
