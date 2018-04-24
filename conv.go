package tensor

import "github.com/ppknap/tensor/dtype"

// Copy TODO.
func (t *Tensor) Copy() *Tensor {
	return nil
}

// View TODO.
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
