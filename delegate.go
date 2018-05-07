package tensor

import (
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Delegate TODO.
type Delegate struct {
	dst *Tensor
}

// NewDelegate creates a new Delegate with a given tensor set as destination
// object for all operations performed on created instance.
func NewDelegate(dst *Tensor) *Delegate {
	return &Delegate{dst: dst}
}

// Add adds elements from tensors 'a' and 'b' element-wise. The result will be
// saved to delegate's destination. If the destination is nil, a new tensor
// will be created. This method allows to use either tensors with scalars or
// tensors that have equal shapes.
func (d *Delegate) Add(a, b *Tensor) (res *Tensor) {
	if a == nil || b == nil {
		panic(errorc.New("nil argument provided"))
	}

	var shape = math.EWArgShape(a.idx, b.idx, true)

	if res = d.dst; res == nil {
		res = New(shape...).AsType(dtype.Promote(a.DType(), b.DType()))
	} else if ds := res.Shape(); !index.EqShape(ds, shape) {
		panic(errorc.New("invalid dst shape %v for %v", ds, shape))
	}

	math.Binary(res.idx, a.idx, b.idx, res.buf, a.buf, b.buf, false, routine.Add)

	return res
}

// Subtract substracts elements 'b' from tensor 'a' element-wise. The result
// will be saved to delegate's destination. If the destination is nil, a new
// tensor will be created. This method allows to use either tensors with scalars
// or tensors that have equal shapes.
func (d *Delegate) Subtract(a, b *Tensor) (res *Tensor) {
	return nil
}

// Multiply multiplies elements from tensors 'a' and 'b' element-wise. The
// result will be saved to delegate's destination. If the destination is nil,
// a new tensor will be created. This method allows to use either tensors with
// scalars or tensors that have equal shapes.
func (d *Delegate) Multiply(a, b *Tensor) (res *Tensor) {
	return nil
}

// Divide divides elements 'a' by elements from tensor 'b' element-wise. The
// result will be saved to delegate's destination. If the destination is nil,
// a new tensor will be created. This method allows to use either tensors with
// scalars or tensors that have equal shapes.
func (d *Delegate) Divide(a, b *Tensor) (res *Tensor) {
	return nil
}
