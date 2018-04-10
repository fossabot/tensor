package index

// IsSameShape returns true when shapes of provided indexes are identical.
func IsSameShape(ai, bi *Index) bool {
	if len(ai.shape) != len(bi.shape) {
		return false
	}

	for i := range ai.shape {
		if ai.shape[i] != bi.shape[i] {
			return false
		}
	}

	return true
}
