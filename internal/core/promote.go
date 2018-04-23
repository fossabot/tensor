package core

// Promote selects the best type that could represent both arguments.
func Promote(at, bt DType) DType {
	switch at {
	case Bool:
		return promote(at, bt)
	case Int:
		return promote(at, bt, Bool)
	case Int8:
		return Float64 // TODO
	case Int16:
		return Float64 // TODO
	case Int64:
		return promote(at, bt, Bool, Int)
	case Uint:
		return Float64 // TODO
	case Uint8:
		return Float64 // TODO
	case Uint16:
		return Float64 // TODO
	case Uint32:
		return Float64 // TODO
	case Uint64:
		return Float64 // TODO
	case Float32:
		return Float64 // TODO
	case Float64:
		return promote(at, bt, Bool, Int, Int64)
	case Complex64:
		return Float64 // TODO
	case Complex128:
		return promote(at, bt, Bool, Int, Int64, Float64)
	case String:
		return promote(at, bt, Bool, Int, Int64, Float64, Complex128)
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
	Bool:       {},
	Int:        {},
	Int8:       {},
	Int16:      {},
	Int64:      {},
	Uint:       {},
	Uint8:      {},
	Uint16:     {},
	Uint32:     {},
	Uint64:     {},
	Float32:    {},
	Float64:    {},
	Complex64:  {},
	Complex128: {},
	String:     {},
}
