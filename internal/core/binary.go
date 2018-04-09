package core

import (
	"unsafe"
)

// BinaryFunc represents a mathematical operation that combines two operands and
// produce another element of the field that should be saved in d pointer. Left
// and right operands should not be modified.
type BinaryFunc func(d, l, r unsafe.Pointer)

// BinaryEach is the simplest binary iterator. It walks over all elements in
// destination buffer and calls binary function giving corresponding elements
// from left and right buffers.
func BinaryEach(db, lb, rb *Buffer, op func(DType) BinaryFunc) {
	var fn = binary(db.typ, lb.typ, rb.typ, op)

	leftAt, rightAt := lb.At(), rb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(dst, leftAt(i), rightAt(i))
	})
}

// binary ensures that binary operation function will have all its arguments in
// the exact same type.
func binary(ddt, ldt, rdt DType, op func(DType) BinaryFunc) BinaryFunc {
	var fn = op(ddt)
	if (ddt == ldt) && (ddt == rdt) {
		return fn
	}

	switch ddt {
	case Bool:
		switch {
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
		case ldt == Int64 && rdt != Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsInt64Ptr(r)) }
		case ldt != Int64 && rdt == Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsInt64Ptr(l), r) }
		case ldt != Int64 && rdt != Int64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsInt64Ptr(l), rdt.AsInt64Ptr(r))
			}
		}
	case String:
		switch {
		case ldt == String && rdt != String:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsStringPtr(r)) }
		case ldt != String && rdt == String:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsStringPtr(l), r) }
		case ldt != String && rdt != String:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsStringPtr(l), rdt.AsStringPtr(r))
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
	case Int64:
		return func(d, l, r unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(l) + *(*int64)(r)
		}
	case String:
		return func(d, l, r unsafe.Pointer) {
			*(*string)(d) = *(*string)(l) + *(*string)(r)
		}
	default:
		panic("core: unsupported type: " + dt.String())
	}
}
