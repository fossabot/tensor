package tensor

import (
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Sin computes trigonometric sine on all radian elements.
func (t *Tensor) Sin() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Sin)

	return t
}

// Cos computes trigonometric cosine on all radian elements.
func (t *Tensor) Cos() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Cos)

	return t
}

// Tan computes trigonometric tangent on all radian elements.
func (t *Tensor) Tan() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, false, routine.Tan)

	return t
}
