package tensor

// Add adds elements from a given argument to the called tensor element-wise.
// This method allows to use either a scalar or a tensor that has the same shape
// as a called one.
func (t *Tensor) Add(u *Tensor) *Tensor {
	return t
}

// Subtract substracts 'u' elements from 't' element-wise. This method allows
// to use either a scalar or a tensor that has the same shape as a called one.
func (t *Tensor) Subtract(u *Tensor) *Tensor {
	return t
}

// Multiply multiplies elements from 't' and 'u' element-wise. This method
// allows to use either a scalar or a tensor that has the same shape as
// a called one.
func (t *Tensor) Multiply(u *Tensor) *Tensor {
	return t
}

// Divide divides elements from 't' using values from 'u' element-wise. This
// method allows to use either a scalar or a tensor that has the same shape as
// a called one.
func (t *Tensor) Divide(u *Tensor) *Tensor {
	return t
}

// Mod computes reminders of 't' divided by 'u'. This method allows to use
// either a scalar or a tensor that has the same shape as a called one.
func (t *Tensor) Mod(u *Tensor) *Tensor {
	return t
}

// Negative computes a numerical negative of all tensor elements.
func (t *Tensor) Negative() *Tensor {
	return t
}
