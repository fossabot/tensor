package tensor

import (
	"github.com/ppknap/tensor/internal/math"
	"github.com/ppknap/tensor/internal/routine"
)

// Add adds elements from a given argument to the called tensor element-wise.
// This method allows to use either a scalar or a tensor that has the same shape
// as a called one.
func (t *Tensor) Add(u *Tensor) *Tensor {
	t.init()
	u.init()

	return t.Delegate().Add(t, u)
}

// Subtract substracts 'u' elements from 't' element-wise. This method allows
// to use either a scalar or a tensor that has the same shape as a called one.
func (t *Tensor) Subtract(u *Tensor) *Tensor {
	t.init()
	u.init()

	return t.Delegate().Subtract(t, u)
}

// Multiply multiplies elements from 't' and 'u' element-wise. This method
// allows to use either a scalar or a tensor that has the same shape as
// a called one.
func (t *Tensor) Multiply(u *Tensor) *Tensor {
	t.init()
	u.init()

	return t.Delegate().Multiply(t, u)
}

// Divide divides elements from 't' using values from 'u' element-wise. This
// method allows to use either a scalar or a tensor that has the same shape as
// a called one.
func (t *Tensor) Divide(u *Tensor) *Tensor {
	t.init()
	u.init()

	return t.Delegate().Divide(t, u)
}

// Mod computes reminders of 't' divided by 'u'. This method allows to use
// either a scalar or a tensor that has the same shape as a called one.
func (t *Tensor) Mod(u *Tensor) *Tensor {
	t.init()
	u.init()

	return t.Delegate().Mod(t, u)
}

// Negative computes a numerical negative of all tensor elements.
func (t *Tensor) Negative() *Tensor {
	t.init()

	math.Nullary(t.idx, t.buf, true, routine.Negative)

	return t
}
