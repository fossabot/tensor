package tacvs

import "fmt"

type Tensor struct {
	// FmtMaxElems defines the number of elements returned when pretty printed.
	// It defaults to DefaultMaxFmtElements when less than two.
	FmtMaxElems int `json:"-"`

	data  []complex128
	shape []int
}

func NewTensor(shape ...int) *Tensor {
	var size int
	if len(shape) > 0 {
		size = 1
	}

	// Shrink redundant dimmensions.
	shape = mustGe(1, shrinkRight(shape, 1, 1))

	for _, n := range shape {
		size *= n
	}

	return &Tensor{
		data:  make([]complex128, size),
		shape: shape,
	}
}

func (t *Tensor) At(idx ...int) complex128 {
	// Shrink redundant indexes.
	idx = mustGe(0, shrinkRight(idx, 0, len(t.shape)))

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

func (t *Tensor) Split(dim int) []*Tensor {
	return nil
}

func (t *Tensor) Slice(dim, from int, to ...int) *Tensor {
	return nil
}

func (t *Tensor) T(perms ...int) *Tensor {
	return nil
}

func (t *Tensor) ConjT(perms ...int) *Tensor {
	return nil
}

// Resize changes shape and size of the tensor. Elements from returned array
// will be at the same positions as they were in the old shape. Empty space will
// be filled with zero values.
//
// This function changes the underlying tensor but, it returns a new one
// when called on views.
func (t *Tensor) Resize(shape ...int) *Tensor {
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

// Shape returns the size of the tensor in each of its dimensions.
func (t *Tensor) Shape() []int {
	cp := make([]int, len(t.shape))
	copy(cp, t.shape)

	return cp
}

// Data returns the internal root buffer of the tensor. The returned slice may
// not point to the data when called on views.
func (t *Tensor) Data() []complex128 {
	return t.data
}

func shrinkRight(slice []int, val, till int) []int {
	for i := len(slice) - 1; i >= till; i-- {
		if slice[i] == val {
			slice = slice[:i]
		} else {
			break
		}
	}

	return slice
}

func mustGe(min int, slice []int) []int {
	for _, n := range slice {
		if n < min {
			panic(fmt.Sprintf("invalid value in: %v (min:%d)", slice, min))
		}
	}

	return slice
}
