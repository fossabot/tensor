package core

import (
	"fmt"
	"unsafe"
)

func asBool(dst *bool, typ DType, p unsafe.Pointer) {
	switch typ {
	case Bool:
		*dst = *(*bool)(p)
	case Int64:
		*dst = *(*int64)(p) != 0
	}

	panic("core: unsupported type: " + typ.String())
}

func asInt64(dst *int, typ DType, p unsafe.Pointer) {
	switch typ {
	case Bool:
		if *(*bool)(p) {
			*dst = 1
		}
		*dst = 0
	case Int64:
		*dst = int(*(*int64)(p))
	}

	panic("core: unsupported type: " + typ.String())
}

func asDTypeBool(st DType, sv unsafe.Pointer) unsafe.Pointer {
	if st == Bool {
		return sv
	}

	var v bool
	asBool(&v, st, sv)
	return unsafe.Pointer(&v)
}

func asDTypeInt64(st DType, sv unsafe.Pointer) unsafe.Pointer {
	if st == Int64 {
		return sv
	}

	var v int
	asInt64(&v, st, sv)
	return unsafe.Pointer(&v)
}

func (dt DType) AsStringFunc() func(unsafe.Pointer) string {
	switch dt {
	case Bool:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*bool)(p)) }
	case Int64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int64)(p)) }
	default:
		panic("core: unsupported type: " + dt.String())
	}
}

func convert(dt, st DType, sv unsafe.Pointer) unsafe.Pointer {
	switch dt {
	case Bool:
		return asDTypeBool(st, sv)
	case Int64:
		return asDTypeInt64(st, sv)
	}

	panic("core: unsupported convert destination type: " + dt.String())
}

// destruct finds data type of provided value. The returned pointer should be
// used only for read operations.
func destruct(v interface{}) (DType, unsafe.Pointer) {
	switch v := v.(type) {
	case bool:
		return Bool, unsafe.Pointer(&v)
	case int64:
		return Int64, unsafe.Pointer(&v)
	}

	panic(fmt.Sprintf("core: unsupported type: %T", v))
}

// setraw sets underlying pointer value from src to dst based on provided data
// type. The values pointed by both pointers must be of identical type and the
// size of provided data type must fit the size of underlying pointers type. If
// any of these requirements is not met, the behavior of this function is
// undefined.
func setraw(t DType, dst, src unsafe.Pointer) {
	switch t.Size() {
	case 1:
		*(*uint8)(dst) = *(*uint8)(src)
	case 2:
		*(*uint16)(dst) = *(*uint16)(src)
	case 4:
		*(*uint32)(dst) = *(*uint32)(src)
	case 8:
		*(*uint64)(dst) = *(*uint64)(src)
	case 16:
		*(*uint64)(dst) = *(*uint64)(src)
		*(*uint64)(unsafe.Pointer(uintptr(dst) + 8)) = *(*uint64)(unsafe.Pointer(uintptr(src) + 8))
	default:
		panic("core: unsupported data size: " + t.String())
	}
}
