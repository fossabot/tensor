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
