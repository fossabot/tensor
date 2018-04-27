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

// AsType sets or changes underlying tensor data type. This method panics if
// called on views.
func (t *Tensor) AsType(dt dtype.DType) *Tensor {
	t.init()
	t.buf.AsType(dt)

	return t
}

// Bool converts tensor scalar to boolean value. This function panics if the
// called tensor stores more than one element.
func (t *Tensor) Bool() bool {
	if t.idx == nil {
		return false
	}

	if t.idx.Size() != 1 {
		panic(core.NewError("cannot convert shape %v to boolean value", t.Shape()))
	}

	return *(*bool)(core.Bool.Convert(
		t.buf.DType(), t.buf.At()(t.idx.At()(nil)),
	))
}

// Byte converts tensor scalar to byte value. This function panics if the called
// tensor stores more than one element.
func (t *Tensor) Byte() byte {
	if t.idx == nil {
		return 0
	}

	if t.idx.Size() != 1 {
		panic(core.NewError("cannot convert shape %v to byte value", t.Shape()))
	}

	return *(*uint8)(core.Uint8.Convert(
		t.buf.DType(), t.buf.At()(t.idx.At()(nil)),
	))
}

// Int converts tensor scalar to integer value. This function panics if the
// called tensor stores more than one element.
func (t *Tensor) Int() int {
	if t.idx == nil {
		return 0
	}

	if t.idx.Size() != 1 {
		panic(core.NewError("cannot convert shape %v to integer value", t.Shape()))
	}

	return *(*int)(core.Int.Convert(
		t.buf.DType(), t.buf.At()(t.idx.At()(nil)),
	))
}

// Float converts tensor scalar to floating point value. This function panics
// if the called tensor stores more than one element.
func (t *Tensor) Float() float64 {
	if t.idx == nil {
		return 0
	}

	if t.idx.Size() != 1 {
		panic(core.NewError("cannot convert shape %v to floating point value", t.Shape()))
	}

	return *(*float64)(core.Float64.Convert(
		t.buf.DType(), t.buf.At()(t.idx.At()(nil)),
	))
}

// Cmplx TODO.
func (t *Tensor) Cmplx() complex128 {
	return 0
}

// Object TODO.
func (t *Tensor) Object() interface{} {
	return nil
}
