package core

// Promote selects the best type that could represent both arguments.
func Promote(at, bt DType) DType {
	switch at {
	case Bool:
		return promote(at, bt)
	case Int:
		return promote(at, bt, Bool)
	case Int64:
		return promote(at, bt, Bool, Int)
	case String:
		return promote(at, bt, Bool, Int, Int64)
	}

	panic(NewError("core: unsupported type: %q", at))
}

func promote(at, bt DType, ts ...DType) DType {
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
	Bool:   {},
	Int:    {},
	Int64:  {},
	String: {},
}
