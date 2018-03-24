package core

import (
	"unsafe"
)

func asBoolean(typ DType, p unsafe.Pointer) bool {
	return false
}

func setReflected(typ DType, dst unsafe.Pointer, v interface{}, f func(dst, vp unsafe.Pointer)) {

}

func setConverted(dtyp, styp DType, dst, src unsafe.Pointer, f func(dst, vp unsafe.Pointer)) {

}
