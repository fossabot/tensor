package index

// Flags stores index properties.
type Flags uint64

// IdxScheme gets indexing scheme stored in flags value.
func (f Flags) IdxScheme() IdxScheme {
	return IdxScheme(f & maskIdxScheme)
}

// View indicates whether the index represents a view over its data.
func (f Flags) View() bool {
	return f&flagIdxView == 1
}

// WithView sets or removes the view flag.
func (f Flags) WithView(view bool) Flags {
	return f.set(view, flagIdxView)
}

func (f Flags) set(v bool, flag Flags) Flags {
	if v {
		return f | flag
	}

	return f &^ flag
}

const (
	maskIdxScheme Flags = 0xFF
	flagIdxView         = 1<<9 + iota
)
