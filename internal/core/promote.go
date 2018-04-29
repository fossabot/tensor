package core

// Promote selects the best type that could store both arguments.
func Promote(at, bt DType) DType {
	if at == bt {
		return at
	}

	switch at {
	case Bool:
		return bt
	case Int:
		switch bt {
		case Bool, Int8, Int16, Int32, Uint8, Uint16:
			return at
		case Uint32:
			return Int64
		case Uint, Uint64, Uintptr, Float32:
			return Float64
		case Complex64:
			return Complex128
		case Int, Int64, Float64, Complex128, String:
			return bt
		}
	case Int8:
		switch bt {
		case Bool, Int8:
			return at
		case Uint:
			return Float64
		case Uint8:
			return Int16
		case Uint16:
			return Int32
		case Uint32:
			return Int64
		case Uint64, Uintptr:
			return Float64
		case Int, Int16, Int32, Int64, Float32, Float64, Complex64, Complex128, String:
			return bt
		}
	case Int16:
		switch bt {
		case Bool:
			return at
		case String:
			return bt
		}
	case Int32:
		return Float64 // TODO
	case Int64:
		return Float64 // TODO
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
	case Uintptr:
		return Float64 // TODO
	case Float32:
		return Float64 // TODO
	case Float64:
		return Float64 // TODO
	case Complex64:
		switch bt {
		case Bool, Int8, Int16, Uint8, Uint16, Float32, Complex64:
			return at
		case Int, Int32, Int64, Uint, Uint32, Uint64, Float64, Uintptr:
			return Complex128
		case Complex128, String:
			return bt
		}
	case Complex128:
		switch bt {
		case Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32,
			Uint64, Uintptr, Float32, Float64, Complex64, Complex128:
			return at
		case String:
			return bt
		}
	case String:
		return at
	}

	panic(NewError("unsupported type promotion: %q and %q", at, bt))
}
