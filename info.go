package tacvs

// Data returns the internal root buffer of the tensor. The returned slice may
// not point to the data when called on views.
func (t *Tensor) Data() []complex128 {
	return t.data
}

// NDim returns the number of tensor dimensions.
func (t *Tensor) NDim() int {
	l := len(t.shape)
	if l == 1 {
		return 2
	}

	return l
}

// Size returns the total number of elements stored in a tensor. It is equal to
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

	if len(cp) == 1 {
		cp = append(cp, 1)
	}

	return cp
}
