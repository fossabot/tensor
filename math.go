package tacvs

// Sum returns a sum of all tensor elements.
func (t *Tensor) Sum() (sum complex128) {
	for i := range t.data {
		sum += t.data[i]
	}

	return sum
}
