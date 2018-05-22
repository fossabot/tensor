package tensor

import "github.com/ppknap/tensor/internal/errorc"

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

// Split splits the tensor into multiple sub-tensor views along the given
// dimmension.
func (t *Tensor) Split(dim int) []*Tensor {
	if nd := t.NDim(); nd == 0 {
		return []*Tensor{t}
	} else if dim >= nd {
		panic(errorc.New("invalid dimension %d (max: %d)", dim, nd-1))
	}

	var (
		indexes = t.idx.Split(dim)
		tensors = make([]*Tensor, len(indexes))
	)

	for i := range indexes {
		tensors[i] = &Tensor{
			idx: indexes[i],
			buf: t.buf,
		}
	}

	return tensors
}

// Slice creates a view over Tensor elements taking them from a specified range
// along given dimmension. When the 'to' argument is omitted, the upper slicing
// limit will be set to dimmension size. Scalars cannot be sliced.
func (t *Tensor) Slice(dim, from int, to ...int) *Tensor {
	if nd := t.NDim(); nd == 0 {
		panic(errorc.New("slice on scalar value"))
	} else if dim >= nd {
		panic(errorc.New("invalid dimension %d (max: %d)", dim, nd-1))
	}

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
