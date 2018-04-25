package tensor

import (
	"github.com/ppknap/tensor/dtype"
	"github.com/ppknap/tensor/internal/core"
)

// Copy creates a copy of called tensor. If this function is called on views, a
// newly created object will own copied data. Thus, the view property will be
// removed.
func (t *Tensor) Copy() *Tensor {
	if t.idx == nil || t.buf == nil {
		return &Tensor{}
	}

	cp := &Tensor{idx: t.idx.CopyNoView()}
	if t.IsOwner() {
		cp.buf = t.buf.Copy()
		return cp
	}
	typ := t.buf.DType()
	cp.buf = core.NewBuffer(cp.idx.Size()).AsType(typ)

	cpSetptr, cpAt := cp.buf.Setptr(), cp.idx.At()
	tBufAt, tIdxAt := t.buf.At(), t.idx.At()

	cp.idx.Iterate(func(pos []int) {
		cpSetptr(cpAt(pos), typ, tBufAt(tIdxAt(pos)))
	})

	return cp
}

// View creates a view over the tensor. Views share the same data as their
// owners but may differ in shape and element order.
func (t *Tensor) View() *Tensor {
	if t.idx != nil && t.idx.Flags().IsView() {
		return t
	}

	return &Tensor{
		idx: t.idx.View(),
		buf: t.buf,
	}
}

// AsType TODO.
func (t *Tensor) AsType(dt dtype.DType) *Tensor {
	return nil
}

// Bool TODO.
func (t *Tensor) Bool() bool {
	return false
}

// Byte TODO.
func (t *Tensor) Byte() byte {
	return 0
}

// Int TODO.
func (t *Tensor) Int() int {
	return 0
}

// Float TODO.
func (t *Tensor) Float() float64 {
	return 0
}

// Cmplx TODO.
func (t *Tensor) Cmplx() complex128 {
	return 0
}

// Object TODO.
func (t *Tensor) Object() interface{} {
	return nil
}
