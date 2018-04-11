package index

// EqShape returns true when provided shapes are identical.
func EqShape(ai, bi []int) bool {
	if len(ai) != len(bi) {
		return false
	}

	for i := range ai {
		if ai[i] != bi[i] {
			return false
		}
	}

	return true
}
