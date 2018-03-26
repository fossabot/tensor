package core

import (
	"fmt"
	"unsafe"
)

// Destruct finds data type of provided value. The returned pointer should be
// used only for read operations.
func Destruct(v interface{}) (DType, unsafe.Pointer) {
	switch v := v.(type) {
	case bool:
		return Bool, unsafe.Pointer(&v)
	case int64:
		return Int64, unsafe.Pointer(&v)
	}

	panic(fmt.Sprintf("core: unsupported type: %T", v))
}