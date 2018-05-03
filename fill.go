package tensor

import (
	"github.com/ppknap/tensor/internal/math"
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

	// Check if v is assignable to t.
	_ = math.EWArgShape(t.idx, v.idx, false)

	math.Unary(t.idx, v.idx, t.buf, v.buf, false, math.Fill)

	return t
}

// Arange TODO.
func (*Tensor) Arange(start, step *Tensor) *Tensor {
	return nil
}

// Linspace TODO.
func (*Tensor) Linspace(start, end *Tensor) *Tensor {
	return nil
}

// Eye sets all main diagonal elements to one. All other elements will be set
// to zero. This mathod works on tensors of any dimensionality.
func (t *Tensor) Eye() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, true, math.Eye)

	return t
}
