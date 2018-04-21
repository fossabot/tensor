package core

import (
	"unsafe"
)

// Destruct finds data type of provided value. The returned pointer should be
// used only for read operations.
func Destruct(v interface{}) (DType, unsafe.Pointer) {
	switch v := v.(type) {
	case bool:
		return Bool, unsafe.Pointer(&v)
	case int:
		return Int, unsafe.Pointer(&v)
	case int64:
		return Int64, unsafe.Pointer(&v)
	case uint:
		return Uint, unsafe.Pointer(&v)
	case uint8:
		return Uint8, unsafe.Pointer(&v)
	case uint64:
		return Uint64, unsafe.Pointer(&v)
	case float32:
		return Float32, unsafe.Pointer(&v)
	case float64:
		return Float64, unsafe.Pointer(&v)
	case complex128:
		return Complex128, unsafe.Pointer(&v)
	case string:
		return String, unsafe.Pointer(&v)
	}

	panic(NewError("core: unsupported type: %T", v))
}
