package tacvs

// Sum returns a sum of all tensor elements.
func (t *Tensor) Sum() (sum complex128) {
	if t.parent != nil {
		t.Apply(func(t *Tensor, idx []int) {
			sum += t.At(idx...)
		})

		return sum
	}

	for i := range t.data {
		sum += t.data[i]
	}

	return sum
}

// Min returns minimum value of all tensor elements. It returns 0 when the
// tensor is empty.
func (t *Tensor) Min() (min complex128) {
	return
}

// Max returns maximum value of all tensor elements. It returns 0 when the
// tensor is empty.
func (t *Tensor) Max() (max complex128) {
	return
}

// CumSum returns a cumulative sum of all tensor elements.
func (t *Tensor) CumSum() (sum complex128) {
	// TODO: tests.
	return
}

// Mean returns a mean of all tensor elements.
func (t *Tensor) Mean() (mean complex128) {
	// TODO: tests.
	return
}

// Median returns 50th percentile of all tensor elements.
func (t *Tensor) Median() (median complex128) {
	// TODO: tests.
	return
}

// CorrCoef returns correletion coefficient of all tensor elements.
func (t *Tensor) CorrCoef() (cc complex128) {]
	// TODO: tests.
	return
}

// Std returns standard deviation of all tensor elements.
func (t *Tensor) Std() (std complex128) {
	// TODO: tests.
	return
}
