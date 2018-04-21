package core

import (
	"unsafe"
)

// Zero returns an unsafe pointer to a zero value of the given type.
func (dt DType) Zero() unsafe.Pointer {
	switch dt {
	case Bool:
		v := false
		return unsafe.Pointer(&v)
	case Int:
		v := 0
		return unsafe.Pointer(&v)
	case Int64:
		v := int64(0)
		return unsafe.Pointer(&v)
	case Float64:
		v := float64(0)
		return unsafe.Pointer(&v)
	case Complex128:
		v := complex128(0)
		return unsafe.Pointer(&v)
	case String:
		v := ""
		return unsafe.Pointer(&v)
	}

	panic(NewError("unsupported type: %q", dt))
}
