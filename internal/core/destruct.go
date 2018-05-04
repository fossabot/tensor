package core

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/errorc"
)

// Destruct finds data type of provided value. The returned pointer should be
// used only for read operations.
func Destruct(v interface{}) (DType, unsafe.Pointer) {
	switch v := v.(type) {
	case bool:
		return Bool, unsafe.Pointer(&v)
	case int:
		return Int, unsafe.Pointer(&v)
	case int8:
		return Int8, unsafe.Pointer(&v)
	case int16:
		return Int16, unsafe.Pointer(&v)
	case int32:
		return Int32, unsafe.Pointer(&v)
	case int64:
		return Int64, unsafe.Pointer(&v)
	case uint:
		return Uint, unsafe.Pointer(&v)
	case uint8:
		return Uint8, unsafe.Pointer(&v)
	case uint16:
		return Uint16, unsafe.Pointer(&v)
	case uint32:
		return Uint32, unsafe.Pointer(&v)
	case uint64:
		return Uint64, unsafe.Pointer(&v)
	case uintptr:
		return Uintptr, unsafe.Pointer(&v)
	case float32:
		return Float32, unsafe.Pointer(&v)
	case float64:
		return Float64, unsafe.Pointer(&v)
	case complex64:
		return Complex64, unsafe.Pointer(&v)
	case complex128:
		return Complex128, unsafe.Pointer(&v)
	case string:
		return String, unsafe.Pointer(&v)
	}

	panic(errorc.New("core: unsupported type: %T", v))
}
