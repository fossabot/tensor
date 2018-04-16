package tensor

import "github.com/ppknap/tacvs/dtype"

// NDim TODO.
func (t *Tensor) NDim() int {
	return 0
}

// Owner TODO.
func (t *Tensor) Owner() bool {
	return false
}

// Base TODO.
func (t *Tensor) Base() *Tensor {
	return nil
}

// Shape TODO.
func (t *Tensor) Shape() []int {
	return t.idx.Shape()
}

// Strides TODO
func (t *Tensor) Strides() []int {
	return nil
}

// Reshape TODO.
func (t *Tensor) Reshape(shape ...int) *Tensor {
	return nil
}

// Size TODO.
func (t *Tensor) Size() int {
	return 0
}

// NBytes TODO.
func (t *Tensor) NBytes() int {
	return 0
}

// Data TODO.
func (t *Tensor) Data() []byte {
	return nil
}

// FillBuf TODO.
func (t *Tensor) FillBuf(data []interface{}) *Tensor {
	return nil
}

// DType TODO.
func (t *Tensor) DType() dtype.DType {
	return t.buf.DType()
}

// AsType TODO.
func (t *Tensor) AsType(typ dtype.DType) *Tensor {
	return t
}
