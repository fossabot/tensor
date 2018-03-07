package tacvs

// Data returns the internal root buffer of the tensor. The returned slice may
// not point to the data when called on views.
func (t *Tensor) Data() []complex128 {
	return t.data
}

// NDim returns the number of tensor dimensions.
func (t *Tensor) NDim() int {
	return len(t.shape)
}

// Size returns the total number of elements stored in tensor. It is equal to
// the product of shape elements.
func (t *Tensor) Size() int {
	if len(t.shape) == 0 {
		return 0
	}

	size := 1
	for i := range t.shape {
		size *= t.shape[i]
	}

	return size
}

// Shape returns the size of the tensor in each of its dimensions.
func (t *Tensor) Shape() []int {
	cp := make([]int, len(t.shape))
	copy(cp, t.shape)

	return cp
}
