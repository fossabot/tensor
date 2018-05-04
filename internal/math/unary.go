package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/buffer"
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
)

// UnaryFunc represents a mathematical operation with only one operand. The
// result should be stored in destination pointer. The input one must not be
// modified.
type UnaryFunc func(pos []int, d, s unsafe.Pointer)

// Unary choses and executes the best strategy to call unary operator on
// provided buffers with respect to their indexes.
func Unary(di, si *index.Index, db, sb *buffer.Buffer, needsPos bool, op func(dtype.DType) UnaryFunc) {
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
func unaryRawEach(db, sb *buffer.Buffer, fn UnaryFunc) {
	srcAt := sb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(nil, dst, srcAt(i))
	})
}

// unaryIdxEach walks over elements in destination buffer pointed by all of its
// index's indices. It calls given unary function with source elements. Each
// element is found by their indexes using destination index indices.
func unaryIdxEach(di, si *index.Index, db, sb *buffer.Buffer, fn UnaryFunc) {
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
func unaryConvert(ddt, sdt dtype.DType, op func(dtype.DType) UnaryFunc) UnaryFunc {
	var fn = op(ddt)
	if ddt == sdt {
		return fn
	}

	switch ddt {
	case dtype.Bool:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsBoolPtr(s)) }
	case dtype.Int:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsIntPtr(s)) }
	case dtype.Int8:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt8Ptr(s)) }
	case dtype.Int16:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt16Ptr(s)) }
	case dtype.Int32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt32Ptr(s)) }
	case dtype.Int64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsInt64Ptr(s)) }
	case dtype.Uint:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUintPtr(s)) }
	case dtype.Uint8:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint8Ptr(s)) }
	case dtype.Uint16:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint16Ptr(s)) }
	case dtype.Uint32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint32Ptr(s)) }
	case dtype.Uint64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUint64Ptr(s)) }
	case dtype.Uintptr:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsUintptrPtr(s)) }
	case dtype.Float32:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsFloat32Ptr(s)) }
	case dtype.Float64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsFloat64Ptr(s)) }
	case dtype.Complex64:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsComplex64Ptr(s)) }
	case dtype.Complex128:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsComplex128Ptr(s)) }
	case dtype.String:
		return func(pos []int, d, s unsafe.Pointer) { fn(pos, d, sdt.AsStringPtr(s)) }
	}

	panic(errorc.New("unsupported destination type: %q", ddt))
}
