package index

// Flags stores index properties.
type Flags uint64

// IdxScheme gets indexing scheme stored in flags value.
func (f Flags) IdxScheme() IdxScheme {
	if scheme := IdxScheme(f & maskIdxScheme); scheme != 0 {
		return scheme
	}

	return IdxSchemeColMajor
}

// IsView indicates whether the index represents a view over the data.
func (f Flags) IsView() bool {
	return f&flagIdxView != 0
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
