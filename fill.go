package tacvs

import "math/rand"

// Zeros sets each tensor element to its zero value.
func (t *Tensor) Zeros() *Tensor {
	// TODO: tests
	return nil
}

// Ones fills all elements of the tensor with ones.
func (t *Tensor) Ones() *Tensor {
	// TODO: tests
	return nil
}

// Full replaces all tensor elements with a given constant.
func (t *Tensor) Full(val complex128) *Tensor {
	// TODO: tests
	return nil
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
	return nil
}

// Random sets each element's real part value to a pseudo-random number in
// [0.0,1.0) range using provided random source. If the provided argument is
// nil, default random source will be used.
func (t *Tensor) Random(source rand.Source) *Tensor {
	// TODO: tests
	return nil
}

// Re removes imaginary part value from each element of a tensor.
func (t *Tensor) Re() *Tensor {
	// TODO: tests
	return nil
}

// Im removes real part value from each element of a tensor.
func (t *Tensor) Im() *Tensor {
	// TODO: tests
	return nil
}

// Apply iterates over all tensor elements and calls f. The returned value will
// be set at given tensor index.
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
