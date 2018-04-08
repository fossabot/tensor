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
	}

	panic("core: unsupported type: " + dt.String())
}
