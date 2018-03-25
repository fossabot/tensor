package core

import (
	"unsafe"
)

func asBool(typ DType, p unsafe.Pointer) bool {
	switch typ {
	case Bool:
		return *(*bool)(p)
	case Int64:
		return *(*int64)(p) != 0
	}

	panic("core: unsupported type: " + typ.String())
}

func asInt64(typ DType, p unsafe.Pointer) int {
	switch typ {
	case Bool:
		if *(*bool)(p) {
			return 1
		}
		return 0
	case Int64:
		return int(*(*int64)(p))
	}

	panic("core: unsupported type: " + typ.String())
}

func asDTypeBool(st DType, sv unsafe.Pointer) unsafe.Pointer {
	if st == Bool {
		return sv
	}

	v := asBool(st, sv)
	return unsafe.Pointer(&v)
}

func asDTypeInt64(st DType, sv unsafe.Pointer) unsafe.Pointer {
	if st == Int64 {
		return sv
	}

	v := asInt64(st, sv)
	return unsafe.Pointer(&v)
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

func setReflected(typ DType, dst unsafe.Pointer, v interface{}, f func(dst, vp unsafe.Pointer)) {

}

func setConverted(dtyp, styp DType, dst, src unsafe.Pointer, f func(dst, vp unsafe.Pointer)) {

}
