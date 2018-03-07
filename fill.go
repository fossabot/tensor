package tacvs

// Fill copies provided buffer to tensor in column-wise order. The size of
// provided slice must be identical as the size of the tensor.
func (t *Tensor) Fill(vs []complex128) *Tensor {
	if len(vs) != len(t.data) {
		panic("fill with a buffer of invalid size")
	}

	for i := range t.data {
		t.data[i] = vs[i]
	}

	return t
}
