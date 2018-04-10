package core

// Merge selects the best type that could represent both arguments.
func Merge(at, bt DType) DType {
	switch at {
	case Bool:
		return merge(at, bt)
	case Int:
		return merge(at, bt, Bool)
	case Int64:
		return merge(at, bt, Bool, Int)
	case String:
		return merge(at, bt, Bool, Int, Int64)
	}

	panic(NewError("core: unsupported type: %q", at))
}

func merge(at, bt DType, ts ...DType) DType {
	for _, t := range ts {
		if t == bt {
			return at
		}
	}

	if _, ok := dTypeSet[bt]; !ok {
		panic(NewError("core: unsupported type: %q", bt))
	}

	return bt
}

var dTypeSet = map[DType]struct{}{
	Bool:   struct{}{},
	Int:    struct{}{},
	Int64:  struct{}{},
	String: struct{}{},
}
