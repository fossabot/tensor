package tensor

import "github.com/ppknap/tensor/dtype"

// NDim returns the number of dimensions of the tensor.
func (t *Tensor) NDim() int {
	if t.idx != nil {
		return t.idx.NDim()
	}

	return 0
}

// Shape returns the shape of a tensor. For scalars this method returns nil.
func (t *Tensor) Shape() []int {
	if t.idx != nil {
		return t.idx.Shape()
	}

	return nil
}

// Strides returns the number of bytes to step in each dimmension to get the
// next element on traversed axis.
func (t *Tensor) Strides() []int {
	if t.idx == nil || t.buf == nil {
		return nil
	}

	strides, typ := t.idx.Strides(), t.buf.DType()

	for i := range strides {
		strides[i] *= int(typ.Size())
	}

	return strides
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
