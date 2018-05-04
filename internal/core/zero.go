package core

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/errorc"
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
	case Int8:
		v := int8(0)
		return unsafe.Pointer(&v)
	case Int16:
		v := int16(0)
		return unsafe.Pointer(&v)
	case Int32:
		v := int32(0)
		return unsafe.Pointer(&v)
	case Int64:
		v := int64(0)
		return unsafe.Pointer(&v)
	case Uint:
		v := uint(0)
		return unsafe.Pointer(&v)
	case Uint8:
		v := uint8(0)
		return unsafe.Pointer(&v)
	case Uint16:
		v := uint16(0)
		return unsafe.Pointer(&v)
	case Uint32:
		v := uint32(0)
		return unsafe.Pointer(&v)
	case Uint64:
		v := uint64(0)
		return unsafe.Pointer(&v)
	case Uintptr:
		v := uintptr(0)
		return unsafe.Pointer(&v)
	case Float32:
		v := float32(0)
		return unsafe.Pointer(&v)
	case Float64:
		v := float64(0)
		return unsafe.Pointer(&v)
	case Complex64:
		v := complex64(0)
		return unsafe.Pointer(&v)
	case Complex128:
		v := complex128(0)
		return unsafe.Pointer(&v)
	case String:
		v := ""
		return unsafe.Pointer(&v)
	}

	panic(errorc.New("unsupported type: %q", dt))
}
