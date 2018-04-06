package core

import (
	"unsafe"
)

// BinaryFunc represents a mathematical operation that combines two operands and
// produce another element of the field that should be saved in d pointer. Left
// and right operands should not be modified.
type BinaryFunc func(d, l, r unsafe.Pointer)

// Binary ensures that binary operation function will have all its arguments in
// the exact same type.
func Binary(ddt, ldt, rdt DType, op func(DType) BinaryFunc) BinaryFunc {
	var fn = op(ddt)

	switch ddt {
	case Bool:
		switch {
		case ldt == Bool && rdt == Bool:
			return fn
		case ldt == Bool && rdt != Bool:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsBoolPtr(r)) }
		case ldt != Bool && rdt == Bool:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsBoolPtr(l), r) }
		case ldt != Bool && rdt != Bool:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsBoolPtr(l), rdt.AsBoolPtr(r))
			}
		}
	case Int:
		switch {
		case ldt == Int && rdt == Int:
			return fn
		case ldt == Int && rdt != Int:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsIntPtr(r)) }
		case ldt != Int && rdt == Int:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsIntPtr(l), r) }
		case ldt != Int && rdt != Int:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsIntPtr(l), rdt.AsIntPtr(r))
			}
		}
	case Int64:
		switch {
		case ldt == Int64 && rdt == Int64:
			return fn
		case ldt == Int64 && rdt != Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsInt64Ptr(r)) }
		case ldt != Int64 && rdt == Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsInt64Ptr(l), r) }
		case ldt != Int64 && rdt != Int64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsInt64Ptr(l), rdt.AsInt64Ptr(r))
			}
		}
	}

	panic("core: unsupported destination type: " + ddt.String())
}

func add(dt DType) BinaryFunc {
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
