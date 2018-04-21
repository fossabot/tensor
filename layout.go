package tensor

import "github.com/ppknap/tensor/dtype"

// NDim TODO.
func (t *Tensor) NDim() int {
	return 0
}

// Shape TODO.
func (t *Tensor) Shape() []int {
	return t.idx.Shape()
}

// Strides TODO
func (t *Tensor) Strides() []int {
	return nil
}

// Size TODO.
func (t *Tensor) Size() int {
	return 0
}

// Owner TODO.
func (t *Tensor) Owner() bool {
	return false
}

// NBytes TODO.
func (t *Tensor) NBytes() int {
	return 0
}

// Base TODO.
func (t *Tensor) Base() *Tensor {
	return nil
}

// Data TODO.
func (t *Tensor) Data() []byte {
	// TODO: testgen when fill func is implemented.
	return nil
}

// FillBuf TODO.
func (t *Tensor) FillBuf(data interface{}) *Tensor {
	return nil
}

// DType TODO.
func (t *Tensor) DType() dtype.DType {
	return t.buf.DType()
}
