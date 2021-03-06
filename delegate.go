package tensor

import (
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Delegate allows to optimize memory usage. A delegated tensor can store
// results of operations performed on its delegate. Thus, the system does not
// need to allocate a new tensor for each mathematical operation.
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
	return d.call(a, b, routine.Add)
}

// Subtract subtracts elements 'b' from tensor 'a' element-wise. The result
// will be saved to delegate's destination. If the destination is nil, a new
// tensor will be created. This method allows to use either tensors with scalars
// or tensors that have equal shapes.
func (d *Delegate) Subtract(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Subtract)
}

// Multiply multiplies elements from tensors 'a' and 'b' element-wise. The
// result will be saved to delegate's destination. If the destination is nil,
// a new tensor will be created. This method allows to use either tensors with
// scalars or tensors that have equal shapes.
func (d *Delegate) Multiply(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Multiply)
}

// Divide divides elements 'a' by elements from tensor 'b' element-wise. The
// result will be saved to delegate's destination. If the destination is nil,
// a new tensor will be created. This method allows to use either tensors with
// scalars or tensors that have equal shapes.
func (d *Delegate) Divide(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Divide)
}

// Mod computes remainders of elements 'a' divided by elements from tensor 'b'
// element-wise. The result will be saved to delegate's destination. If the
// destination is nil, a new tensor will be created. This method allows to use
// either tensors with scalars or tensors that have equal shapes.
func (d *Delegate) Mod(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Mod)
}

// Maximum is a element-wise maximum of tensor elements. It propagates NaN
// values. The result will be saved to delegate's destination. If the
// destination is nil, a new tensor will be created. This method allows to use
// either tensors with scalars or tensors that have equal shapes.
func (d *Delegate) Maximum(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Maximum)
}

// Minimum is a element-wise minimum of tensor elements. It propagates NaN
// values. The result will be saved to delegate's destination. If the
// destination is nil, a new tensor will be created. This method allows to use
// either tensors with scalars or tensors that have equal shapes.
func (d *Delegate) Minimum(a, b *Tensor) (res *Tensor) {
	return d.call(a, b, routine.Minimum)
}

func (d *Delegate) call(a, b *Tensor, fn func(dtype.DType) math.BinaryFunc) *Tensor {
	if a == nil || b == nil {
		panic(errorc.New("nil argument provided"))
	}

	var shape = math.ElementWiseDstShape(a.idx, b.idx, true)

	var dst = d.dst
	if dst == nil {
		dst = New(shape...).AsType(dtype.Promote(a.DType(), b.DType()))
	} else if ds := dst.Shape(); !index.EqShape(ds, shape) {
		panic(errorc.New("invalid dst shape %v for %v", ds, shape))
	}

	math.Binary(dst.idx, a.idx, b.idx, dst.buf, a.buf, b.buf, false, fn)

	return dst
}
