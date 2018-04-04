package core

import (
	"unsafe"
)

func Binary(ddt, ldt, rdt DType, op func(d, l, r unsafe.Pointer)) {

}

func operator(dt DType, ptrDst, ptrLhv, ptrRhv unsafe.Pointer) {
}

func selectConversion(ddt, ldt, rdt DType) func(dt DType, d, l, r unsafe.Pointer) {
	switch ddt {
	case Bool:
		switch {
		case ldt == Bool && rdt == Bool:

		case ldt == Bool && rdt != Bool:

		case ldt != Bool && rdt == Bool:

		case ldt != Bool && rdt != Bool:

		}
		*dst = *(*bool)(p)
	case Int:
		*dst = *(*int)(p) != 0
	default:
		panic("core: unsupported destination type: " + ddt.String())
	}
	return func(dt DType, d, l, r unsafe.Pointer) {

	}
}

func add(dt DType) func(d, l, r unsafe.Pointer) {
	switch dt {
	case Bool:
		return func(d, l, r unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(l) || *(*bool)(r)
		}
	case Int:
		return func(d, l, r unsafe.Pointer) {
			*(*int)(d) = *(*int)(l) + *(*int)(r)
		}
	default:
		panic("core: unsupported type: " + dt.String())
	}
}
