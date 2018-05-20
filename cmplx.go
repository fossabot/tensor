package tensor

import (
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Real removes imaginary part from all elements. This method is a no-op for all
// non-complex numerical elements.
func (t *Tensor) Real() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Real)

	return t
}

// Imag removes real part from all elements. Non complex numerical elements will
// be set to their zero values.
func (t *Tensor) Imag() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Imag)

	return t
}
