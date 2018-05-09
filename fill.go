package tensor

import (
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Zeros fills all elements of called tensor with zeroes.
func (t *Tensor) Zeros() *Tensor {
	return t.Fill(NewScalar(0.))
}

// Ones fills all elements of called tensor with ones.
func (t *Tensor) Ones() *Tensor {
	return t.Fill(NewScalar(1.))
}

// Fill sets all elements of called tensor with provided value. This function
// panics if a given argument is not a scalar or the shapes of two tensors do
// not equal.
func (t *Tensor) Fill(v *Tensor) *Tensor {
	t.init()
	v.init()

	// Check if v is assignable to t.
	_ = math.ElementWiseDstShape(t.idx, v.idx, false)

	math.Unary(t.idx, v.idx, t.buf, v.buf, false, routine.Fill)

	return t
}

// Arange fills the tensor starting from a 'start' scalar at zero position and
// increasing it every iteration by a given 'step' scalar.
func (t *Tensor) Arange(start, step *Tensor) *Tensor {
	start.mustScalar("start")
	step.mustScalar("step")

	t.init()
	start.init()
	step.init()

	math.Binary(t.idx, start.idx, step.idx, t.buf, start.buf, step.buf, true, routine.Arange)

	return t
}

// Linspace fills the tensor with evenly spaced samples within a given
// ['start', 'end'] interval. Both, 'start' and 'end' tensors must be scalars.
func (t *Tensor) Linspace(start, end *Tensor) *Tensor {
	start.mustScalar("start")
	end.mustScalar("end")

	t.init()
	start.init()
	end.init()

	fn := routine.Linspace(t.Size())
	math.Binary(t.idx, start.idx, end.idx, t.buf, start.buf, end.buf, true, fn)

	return t
}

// Eye sets all main diagonal elements to one. All other elements will be set
// to zero. This mathod works on tensors of any dimensionality.
func (t *Tensor) Eye() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, true, routine.Eye)

	return t
}
