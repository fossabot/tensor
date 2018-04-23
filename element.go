package tensor

import "github.com/ppknap/tensor/internal/core"

// Each TODO.
func (t *Tensor) Each(f func(pos []int, t *Tensor)) *Tensor {
	return t
}

// ItemAt TODO.
func (t *Tensor) ItemAt(pos ...int) *Tensor {
	if !t.idx.Validate(pos) {
		panic(core.NewError("invalid position %v for %v", pos, t.idx))
	}

	return &Tensor{
		idx: t.idx.Scalar(pos),
		buf: t.buf,
	}
}

// ItemSet TODO.
func (t *Tensor) ItemSet(v *Tensor, idx ...int) *Tensor {
	return t
}
