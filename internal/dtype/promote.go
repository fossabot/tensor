package dtype

import "github.com/ppknap/tensor/internal/errorc"

// Promote selects the best type that can safely store values of both arguments.
func Promote(at, bt DType) DType {
	// No type promotion for equal types.
	if at == bt {
		return at
	}

	// Reorder types if needed.
	if at.Num() > bt.Num() {
		at, bt = bt, at
	}

	switch at {
	case Bool:
		return bt
	case Int:
		switch bt {
		case Int8, Int16, Int32, Uint8, Uint16:
			return at
		case Uint32:
			return Int64
		case Uint, Uint64, Uintptr, Float32:
			return Float64
		case Complex64:
			return Complex128
		case Int64, Float64, Complex128, String:
			return bt
		}
	case Int8:
		switch bt {
		case Uint8:
			return Int16
		case Uint16:
			return Int32
		case Uint32:
			return Int64
		case Uint, Uint64, Uintptr:
			return Float64
		case Int16, Int32, Int64, Float32, Float64, Complex64, Complex128, String:
			return bt
		}
	case Int16:
		switch bt {
		case Uint8:
			return at
		case Uint16:
			return Int32
		case Uint32:
			return Int64
		case Uint, Uint64, Uintptr:
			return Float64
		case Int32, Int64, Float32, Float64, Complex64, Complex128, String:
			return bt
		}
	case Int32:
		switch bt {
		case Uint8, Uint16:
			return at
		case Uint32:
			return Int64
		case Uint, Uint64, Uintptr, Float32:
			return Float64
		case Complex64:
			return Complex128
		case Int64, Float64, Complex128, String:
			return bt
		}
	case Int64:
		switch bt {
		case Uint8, Uint16, Uint32:
			return at
		case Uint, Uint64, Uintptr, Float32:
			return Float64
		case Complex64:
			return Complex128
		case Float64, Complex128, String:
			return bt
		}
	case Uint:
		switch bt {
		case Uint8, Uint16, Uint32:
			return at
		case Float32:
			return Float64
		case Complex64:
			return Complex128
		case Uint64, Uintptr, Float64, Complex128, String:
			return bt
		}
	case Uint8:
		switch bt {
		case Uint16, Uint32, Uint64, Uintptr, Float32, Float64, Complex64, Complex128, String:
			return bt
		}
	case Uint16:
		switch bt {
		case Uint32, Uint64, Uintptr, Float32, Float64, Complex64, Complex128, String:
			return bt
		}
	case Uint32:
		switch bt {
		case Float32:
			return Float64
		case Complex64:
			return Complex128
		case Uint64, Uintptr, Float64, Complex128, String:
			return bt
		}
	case Uint64:
		switch bt {
		case Uintptr:
			return at
		case Float32:
			return Float64
		case Complex64:
			return Complex128
		case Float64, Complex128, String:
			return bt
		}
	case Uintptr:
		switch bt {
		case Float32:
			return Float64
		case Complex64:
			return Complex128
		case Float64, Complex128, String:
			return bt
		}
	case Float32:
		switch bt {
		case Float64, Complex64, Complex128, String:
			return bt
		}
	case Float64:
		switch bt {
		case Complex64:
			return Complex128
		case Complex128, String:
			return bt
		}
	case Complex64:
		switch bt {
		case Complex128, String:
			return bt
		}
	case Complex128:
		if bt == String {
			return bt
		}
	}

	panic(errorc.New("unsupported type %q and %q", at, bt))
}
