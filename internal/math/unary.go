package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
)

// UnaryFunc represents a mathematical operation with only one operand. The
// result should be stored in destination pointer. The input one must not be
// modified.
type UnaryFunc func(pos []int, d, s unsafe.Pointer)

// Unary choses and executes the best strategy to call unary operator on
// provided buffers with respect to their indexes.
func Unary(di, si *index.Index, db, sb *core.Buffer, needsPos bool, op func(core.DType) UnaryFunc) {
	var fn = unaryConvert(db.DType(), sb.DType(), op)

	if needsPos {
		unaryIdxEach(di, si, db, sb, fn)
	}

	var (
		dScheme, dIsView = di.Flags().IdxScheme(), di.Flags().IsView()
		sScheme, sIsView = si.Flags().IdxScheme(), si.Flags().IsView()
	)

	if (dScheme == sScheme) && !dIsView && !sIsView {
		// Iterate directly on buffers since they have the same memory layout.
		unaryRawEach(db, sb, fn)
		return
	}

	unaryIdxEach(di, si, db, sb, fn)
}

// unaryRawEach is the simplest unary iterator. It walks over all elements in
// destination buffer and calls unary function giving corresponding from source
// buffer.
func unaryRawEach(db, sb *core.Buffer, fn UnaryFunc) {
	srcAt := sb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(nil, dst, srcAt(i))
	})
}

// unaryIdxEach walks over elements in destination buffer pointed by all of its
// index's indices. It calls given unary function with source elements. Each
// element is found by their indexes using destination index indices.
func unaryIdxEach(di, si *index.Index, db, sb *core.Buffer, fn UnaryFunc) {
	var (
		diAt, siAt = di.At(), si.At()
		dbAt, sbAt = db.At(), sb.At()
	)

	di.Iterate(func(pos []int) {
		fn(pos, dbAt(diAt(pos)), sbAt(siAt(pos)))
	})
}

// unaryConvert ensures that unary operation function will have all its
// arguments in the exact same type.
func unaryConvert(ddt, sdt core.DType, op func(core.DType) UnaryFunc) UnaryFunc {
	var fn = op(ddt)
	if ddt == sdt {
		return fn
	}

	switch ddt {
	case core.Bool:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsBoolPtr(s)) }
	case core.Int:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsIntPtr(s)) }
	case core.Int8:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt8Ptr(s)) }
	case core.Int16:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt16Ptr(s)) }
	case core.Int32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt32Ptr(s)) }
	case core.Int64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt64Ptr(s)) }
	case core.Uint:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUintPtr(s)) }
	case core.Uint8:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint8Ptr(s)) }
	case core.Uint16:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint16Ptr(s)) }
	case core.Uint32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint32Ptr(s)) }
	case core.Uint64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint64Ptr(s)) }
	case core.Uintptr:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUintptrPtr(s)) }
	case core.Float32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsFloat32Ptr(s)) }
	case core.Float64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsFloat64Ptr(s)) }
	case core.Complex64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsComplex64Ptr(s)) }
	case core.Complex128:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsComplex128Ptr(s)) }
	case core.String:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsStringPtr(s)) }
	}

	panic(core.NewError("unsupported destination type: %q", ddt))
}
