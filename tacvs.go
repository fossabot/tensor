package tacvs

import "fmt"

type Tensor struct {
	// FmtMaxElems defines the number of elements returned when pretty printed.
	// It defaults to DefaultMaxFmtElements when less than two.
	FmtMaxElems int `json:"-"`

	data []complex128 // tensor buffer, nil for views.

	parent *Tensor // true for tensor views.
	offset int     // offset along dim.
	dim    int

	shape []int // tensor shape.
}

// NewTensor creates a new Tensor object.
func NewTensor(shape ...int) *Tensor {
	return newTensor(nil, 0, 0, shape...)
}

func newTensor(parent *Tensor, offset, dim int, shape ...int) *Tensor {
	// Flat dimmensions.
	shape = mustGe(0, fitIndex(shape, 1, 1))

	// Return a view if parent is present.
	if parent != nil {
		return &Tensor{
			parent: parent,
			offset: offset,
			dim:    dim,
			shape:  shape,
		}
	}

	if len(shape) == 0 {
		return &Tensor{}
	}

	var size = 1
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
	if t.parent != nil {
		idx := t.getResizedIdx(idx)
		for len(idx) <= t.dim {
			idx = append(idx, 0)
		}
		idx[t.dim] += t.offset
		return t.parent.At(idx...)
	}

	return t.data[t.position(t.getResizedIdx(idx))]
}

// Set inserts a value on a given position. It panics when called on empty
// tensor or when the index is out of tensor shape range.
func (t *Tensor) Set(val complex128, idx ...int) *Tensor {
	t.data[t.position(t.getResizedIdx(idx))] = val
	return t
}

// checkIdxConst checks if index is valid in terms of its shape.
func (t *Tensor) getResizedIdx(idx []int) []int {
	if len(idx) == 0 {
		panic("tensor: cannot index empty tensor")
	}

	return checkRange(t.shape, fitIndex(idx, 0, len(t.shape)))
}

// position computes the index of value described by column-wise coordinates.
func (t *Tensor) position(idx []int) (pos int) {
	fmt.Println(t.shape, idx)
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
	if dim >= t.NDim() {
		panic("tensor: invalid dimension")
	}

	if len(to) > 1 {
		panic("tensor: too many slice arguments")
	}

	dimsize := t.shape[dim]
	if from < 0 || from >= dimsize {
		panic(fmt.Sprintf("tensor: invalid from range %d for [0, %d)", from, dimsize))
	}

	limit := dimsize
	if len(to) > 0 {
		limit = to[0]
	}

	if limit < 0 || limit > dimsize {
		panic(fmt.Sprintf("tensor: invalid to range %d for [0, %d]", limit, dimsize))
	}

	if limit-from < 0 {
		panic("tensor: invalid slice range")
	}

	shape := make([]int, len(t.shape))
	copy(shape, t.shape)
	shape[dim] = limit - from

	return newTensor(t, from, dim, shape...)
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

func fitIndex(slice []int, val, till int) []int {
	for len(slice) < till {
		slice = append(slice, val)
	}

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
			panic(fmt.Sprintf("tensor: invalid value in: %v (min:%d)", slice, min))
		}
	}

	return slice
}

func checkRange(shape, idx []int) []int {
	if len(idx) != len(shape) {
		panic("tensor: invalid index")
	}

	for i := range shape {
		if idx[i] < 0 || idx[i] >= shape[i] {
			panic(fmt.Sprintf("tensor: invalid index %v on shape %v", idx, shape))
		}
	}

	return idx
}
