package tensor

import (
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Exp computes e**x of all tensor elements.
func (t *Tensor) Exp() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Exp)

	return t
}

// Pow rises each tensor element to the coresponding value of provided argument.
func (t *Tensor) Pow(v *Tensor) *Tensor {
	t.init()
	v.init()

	// Check if v is assignable to t.
	_ = math.ElementWiseDstShape(t.idx, v.idx, false)

	math.Unary(t.idx, v.idx, t.buf, v.buf, false, routine.Pow)

	return t
}

// Sqrt computes square root of all tensor elements.
func (t *Tensor) Sqrt() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Sqrt)

	return t
}

// Log computes natural logarithm of all tensor elements.
func (t *Tensor) Log() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Log)

	return t
}

// Log10 computes decimal logarithm of all tensor elements.
func (t *Tensor) Log10() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Log10)

	return t
}
