package tensor

// View TODO.
func (t *Tensor) View() *Tensor {
	if t.idx != nil && t.idx.Flags().IsView() {
		return t
	}

	return &Tensor{
		idx: t.idx.View(),
		buf: t.buf,
	}
}
