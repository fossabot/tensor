package tacvs

import "fmt"

type Tensor struct {
	data  []complex128
	shape []int
}

func NewTensor(shape ...int) *Tensor {
	var size int
	if len(shape) > 0 {
		size = 1
	}

	// Shrink redundant dimmensions.
	shape = shrinkRight(shape, 1, 1)

	for _, n := range shape {
		if n == 0 {
			panic(fmt.Sprintf("zero rank in shape: %v", shape))
		}

		size *= n
	}

	return &Tensor{
		data:  make([]complex128, size),
		shape: shape,
	}
}

func (t *Tensor) At(idx ...int) complex128 {
	// Shrink redundant indexes.
	idx = shrinkRight(idx, 0, len(t.shape))

	if len(idx) != len(t.shape) {
		panic(fmt.Sprintf("invalid tensor index %v for shape %v", idx, t.shape))
	}

	if len(idx) == 0 {
		panic("cannot index empty tensor")
	}

	var pos int
	for k := 0; k < len(t.shape); k++ {
		stride := 1
		for j := 0; j < k; j++ {
			stride *= t.shape[j]
		}

		pos += stride * idx[k]
	}

	return t.data[pos]
}

func (t *Tensor) Slice(dim, from, to int) []*Tensor {
	return nil
}

// Size returns the total number of elements stored in tensor. It is equal to
// the product of shape elements.
func (t *Tensor) Size() int {
	if len(t.shape) == 0 {
		return 0
	}

	size := 1
	for i := range t.shape {
		size *= t.shape[i]
	}

	return size
}

// NDim returns the number of tensor dimensions.
func (t *Tensor) NDim() int {
	return len(t.shape)
}

// Shape returns the size of the tensor in each dimension.
func (t *Tensor) Shape() []int {
	cp := make([]int, len(t.shape))
	copy(cp, t.shape)

	return cp
}

// Data returns the internal root buffer of the tensor. Thus, the returned slice
// may not point to the data when called on views.
func (t *Tensor) Data() []complex128 {
	return t.data
}

func shrinkRight(slice []int, val, till int) []int {
	for i := len(slice) - 1; i >= till; i-- {
		if slice[i] == val {
			slice = slice[:i]
		}
	}

	return slice
}
