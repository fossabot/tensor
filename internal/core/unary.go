package core

import (
	"unsafe"
)

// UnaryFunc represents a mathematical operation with only one operand. The
// result should be stored in destination pointer. The input one should not be
// modified.
type UnaryFunc func(d, m unsafe.Pointer)

// Unary ensures that unary operation function will have all its arguments in
// the exact same type.
func Unary(ddt, mdt DType, op func(DType) UnaryFunc) UnaryFunc {
	var fn = op(ddt)
	if ddt == mdt {
		return fn
	}

	switch ddt {
	case Bool:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsBoolPtr(m)) }
	case Int:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsIntPtr(m)) }
	case Int64:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsInt64Ptr(m)) }
	case String:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsStringPtr(m)) }
	}

	panic("core: unsupported destination type: " + ddt.String())
}
