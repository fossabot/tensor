package stale

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

// Mean returns a mean of all tensor elements.
func (t *Tensor) Mean() (mean complex128) {
	return
}

// Median returns 50th percentile of all tensor elements.
func (t *Tensor) Median() (median complex128) {
	return
}

// Std returns standard deviation of all tensor elements.
func (t *Tensor) Std() (std complex128) {
	return
}
