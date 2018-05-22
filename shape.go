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

// Slice creates a view over Tensor elements taking them from a specified range
// along given dimmension. When the 'to' argument is omitted, the upper slicing
// limit will be set to dimmension size. Scalars cannot be sliced.
func (t *Tensor) Slice(dim, from int, to ...int) *Tensor {
	return &Tensor{
		idx: t.idx.Slice(dim, from, to...),
		buf: t.buf,
	}
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
