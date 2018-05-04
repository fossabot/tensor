package tensor

import (
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
	"github.com/ppknap/tensor/internal/math"
)

// Delegate TODO.
type Delegate struct {
	dst *Tensor
}

// NewDelegate TODO.
func NewDelegate(dst *Tensor) *Delegate {
	return &Delegate{dst: dst}
}

// Add TODO.
func (d *Delegate) Add(a, b *Tensor) (dst *Tensor) {
	if a == nil || b == nil {
		panic(errorc.New("nil argument provided"))
	}

	var shape = math.EWArgShape(a.idx, b.idx, true)

	if dst = d.dst; dst == nil {
		//dst = New(shape...).AsType(core.Promote(a.DType(), b.DType()))
	} else if ds := dst.Shape(); !index.EqShape(ds, shape) {
		panic(errorc.New("invalid dst shape %v for %v", ds, shape))
	}

	math.Binary(dst.idx, a.idx, b.idx, dst.buf, a.buf, b.buf, false, math.Add)

	return dst
}

// Delegate TODO.
func (t *Tensor) Delegate() *Delegate {
	return NewDelegate(t)
}
