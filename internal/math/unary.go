package math

import (
	"unsafe"

	"github.com/ppknap/tacvs/internal/core"
)

// UnaryFunc represents a mathematical operation with only one operand. The
// result should be stored in destination pointer. The input one should not be
// modified.
type UnaryFunc func(d, m unsafe.Pointer)

// Unary ensures that unary operation function will have all its arguments in
// the exact same type.
func Unary(ddt, mdt core.DType, op func(core.DType) UnaryFunc) UnaryFunc {
	var fn = op(ddt)
	if ddt == mdt {
		return fn
	}

	switch ddt {
	case core.Bool:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsBoolPtr(m)) }
	case core.Int:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsIntPtr(m)) }
	case core.Int64:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsInt64Ptr(m)) }
	case core.String:
		return func(d, m unsafe.Pointer) { fn(d, mdt.AsStringPtr(m)) }
	}

	panic("core: unsupported destination type: " + ddt.String())
}
