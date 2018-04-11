package tensor

import (
	"fmt"

	"github.com/ppknap/tacvs/internal/core"
	"github.com/ppknap/tacvs/internal/index"
	"github.com/ppknap/tacvs/internal/math"
)

type Tensor struct {
	idx *index.Index
	buf *core.Buffer
}

// New creates a new tensor with a given shape. Empty shape creates a scalar.
func New(shape ...int) *Tensor {
	var idx = index.NewIndex(shape, 0)

	return &Tensor{
		idx: idx,
		buf: core.NewBuffer(idx.Size()),
	}
}

func (t *Tensor) At(pos ...int) *Tensor {
	if !t.idx.Validate(pos) {
		panic(fmt.Sprintf("tensor: invalid position %v for %v", pos, t.idx))
	}

	return &Tensor{
		idx: t.idx.Scalar(pos),
		buf: t.buf,
	}
}

// Each TODO.
func Each(f func(t *Tensor)) *Tensor {
	return nil
}

func (t *Tensor) View() *Tensor {
	if t.idx != nil && t.idx.Flags().IsView() {
		return t
	}

	return &Tensor{
		idx: t.idx.View(),
		buf: t.buf,
	}
}

func (t *Tensor) Delegate() *Delegate {
	return NewDelegate(t)
}

type Delegate struct {
	dst *Tensor
}

func NewDelegate(dst *Tensor) *Delegate {
	return &Delegate{dst: dst}
}

func (d *Delegate) Add(a, b *Tensor) (dst *Tensor) {
	if a == nil || b == nil {
		panic(core.NewError("nil argument provided"))
	}

	var shape = math.EWArgShape(a.idx, b.idx)

	if dst = d.dst; dst == nil {
		dst = New(shape...).AsType(core.Merge(a.DType(), b.DType()))
	} else if ds := dst.Shape(); !index.EqShape(ds, shape) {
		panic(core.NewError("invalid dst shape %v for %v", ds, shape))
	}

	math.Binary(dst.idx, a.idx, b.idx, dst.buf, a.buf, b.buf, math.Add)

	return dst
}

// View
// Scheme
// Size(l) && Size(r) || Size(l) && Size(r) == 1
