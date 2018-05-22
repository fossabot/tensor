package tensor

// T TODO.
func (t *Tensor) T() *Tensor {
	return t
}

// Transpose TODO.
func (t *Tensor) Transpose() *Tensor {
	return t
}

// H TODO.
func (t *Tensor) H() *Tensor {
	return t
}

// Destruct TODO.
func (t *Tensor) Destruct() []*Tensor {
	return []*Tensor{t}
}

// Split TODO.
func (t *Tensor) Split(dim int) []*Tensor {
	return []*Tensor{t}
}

// Slice TODO.
func (t *Tensor) Slice(dim, from int, to ...int) *Tensor {
	return t
}

// Reshape TODO.
func (t *Tensor) Reshape(shape ...int) *Tensor {
	return t
}

// Resize TODO.
func (t *Tensor) Resize(shape ...int) *Tensor {
	return t
}

// Ravel TODO.
func (t *Tensor) Ravel() *Tensor {
	return t
}
