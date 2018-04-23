package tensor

import (
	"fmt"

	"github.com/ppknap/tensor/dtype"
	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
	"github.com/ppknap/tensor/internal/math"
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

// NewScalar creates a 0-dimensional tensor from a given value. The returned
// object's data type will be inherited from a given argument.
func NewScalar(scalar interface{}) *Tensor {
	t := &Tensor{
		idx: index.NewIndex(nil, 0),
		buf: core.NewBuffer(1),
	}

	typ, p := core.Destruct(scalar)
	t.buf.AsType(typ).Setptr()(0, typ, p)

	return t
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

// Reshape TODO.
func (t *Tensor) Reshape(shape ...int) *Tensor {
	return nil
}

// AsType TODO.
func (t *Tensor) AsType(typ dtype.DType) *Tensor {
	return t
}

func (t *Tensor) init() {
	if t.idx == nil && t.buf == nil {
		*t = *New()
	}
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
		dst = New(shape...).AsType(core.Promote(a.DType(), b.DType()))
	} else if ds := dst.Shape(); !index.EqShape(ds, shape) {
		panic(core.NewError("invalid dst shape %v for %v", ds, shape))
	}

	math.Binary(dst.idx, a.idx, b.idx, dst.buf, a.buf, b.buf, math.Add)

	return dst
}

// View
// Scheme
// Size(l) && Size(r) || Size(l) && Size(r) == 1
