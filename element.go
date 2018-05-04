package tensor

import (
	"github.com/ppknap/tensor/internal/errorc"
)

// Each calls f on each element of the tensor. The current iteration position
// and a scalar view of underlying element will be passed to f. These arguments
// may be reused by internal implementation after each iteration.
func (t *Tensor) Each(f func(pos []int, t *Tensor)) *Tensor {
	t.idx.Iterate(func(pos []int) {
		f(pos, &Tensor{
			idx: t.idx.Scalar(pos),
			buf: t.buf,
		})
	})

	return t
}

// ItemAt returns the element at a given position. The returned tensor is a
// mutable scalar view over called object.
func (t *Tensor) ItemAt(pos ...int) *Tensor {
	if !t.idx.Validate(pos) {
		panic(errorc.New("invalid position %v for %v", pos, t.idx))
	}

	return &Tensor{
		idx: t.idx.Scalar(pos),
		buf: t.buf,
	}
}

// ItemSet sets the value from provided scalar at a given position in called
// tensor. This function panics if provided tensor size is not equal to one.
func (t *Tensor) ItemSet(v *Tensor, pos ...int) *Tensor {
	if !t.idx.Validate(pos) {
		panic(errorc.New("invalid position %v for %v", pos, t.idx))
	}

	if v.Size() != 1 {
		panic(errorc.New("invalid non scalar argument (shape:%v)", v.Shape()))
	}

	t.buf.Setptr()(t.idx.At()(pos), v.buf.DType(), v.buf.At()(0))

	return t
}
