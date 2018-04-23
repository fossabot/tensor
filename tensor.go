package tensor

import (
	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
)

// Tensor represents an organized multidimensional array of fixed-size elements.
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

func (t *Tensor) init() {
	if t.idx == nil && t.buf == nil {
		*t = *New()
	}
}
