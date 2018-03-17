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

// At returns a tensor value at a given position. It panics when called on empty
// tensor or when the index is out of tensor shape range.
func (t *Tensor) At(idx ...int) complex128 {
	// Shrink redundant indexes.
	idx = mustGe(0, shrinkRight(idx, 0, len(t.shape)))

	t.checkIdxConst(idx)
	return t.data[t.position(idx)]
}

// Set inserts a value on a given position. It panics when called on empty
// tensor or when the index is out of tensor shape range.
func (t *Tensor) Set(val complex128, idx ...int) *Tensor {
	// Shrink redundant indexes.
	idx = mustGe(0, shrinkRight(idx, 0, len(t.shape)))

	t.checkIdxConst(idx)
	t.data[t.position(idx)] = val
	return t
}

// checkIdxConst checks if index is valid in terms of its shape.
func (t *Tensor) checkIdxConst(idx []int) {
	if len(idx) != len(t.shape) {
		panic(fmt.Sprintf("invalid tensor index %v for shape %v", idx, t.shape))
	}

	if len(idx) == 0 {
		panic("cannot index empty tensor")
	}
}

// position computes the index of value described by column-wise coordinates.
func (t *Tensor) position(idx []int) (pos int) {
	for k := 0; k < len(t.shape); k++ {
		stride := 1
		for j := 0; j < k; j++ {
			stride *= t.shape[j]
		}

		pos += stride * idx[k]
	}

	return pos
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

// IsZero reports whether t represents a zero length tensor.
func (t *Tensor) IsZero() bool {
	return len(t.data) == 0
}

// Clone creates an exact copy of called tensor. When called on views, they will
// be converted to new Tensor instances.
func (t *Tensor) Clone() *Tensor {
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
