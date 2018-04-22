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

// Size returns the size of a tensor which is the same as the product of its
// shape elements. Scalars size is equal to one.
func (t *Tensor) Size() int {
	if t.idx != nil {
		return t.idx.Size()
	}

	return 1
}

// Owner checks if called tensor owns data buffer it uses. This method returns
// false for views.
func (t *Tensor) Owner() bool {
	if t.idx != nil {
		return !t.idx.Flags().IsView()
	}

	return true
}

// NBytes returns the number of bytes consumed by tensor elements. For dynamic
// types which do not own their data, this function returns only the size
// consumed by the pointers to the actual objects.
func (t *Tensor) NBytes() int { return t.buf.NBytes() }

// Base TODO.
func (t *Tensor) Base() *Tensor {
	return nil
}

// Data returns the internal byte slice object with tensor data. For dynamic
// types, this function returns nil.
func (t *Tensor) Data() []byte {
	// TODO: testgen when fill func is implemented.
	t.init()

	return t.buf.Data()
}

// FillBuf TODO.
func (t *Tensor) FillBuf(data interface{}) *Tensor {
	return nil
}

// DType TODO.
func (t *Tensor) DType() dtype.DType {
	return t.buf.DType()
}
