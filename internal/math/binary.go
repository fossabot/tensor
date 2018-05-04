package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/buffer"
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
)

// BinaryFunc represents a mathematical operation that combines two operands and
// produce another element of the field that should be saved in d pointer. Left
// and right operands must not be modified.
type BinaryFunc func(pos []int, d, l, r unsafe.Pointer)

// Binary choses and executes the best strategy to call binary operator on
// provided buffers with respect to their indexes.
func Binary(di, li, ri *index.Index, db, lb, rb *buffer.Buffer, needsPos bool, op func(dtype.DType) BinaryFunc) {
	var fn = binaryConvert(db.DType(), lb.DType(), rb.DType(), op)

	if needsPos {
		binaryIdxEach(di, li, ri, db, lb, rb, fn)
	}

	var (
		dScheme, dIsView = di.Flags().IdxScheme(), di.Flags().IsView()
		lScheme, lIsView = li.Flags().IdxScheme(), li.Flags().IsView()
		rScheme, rIsView = ri.Flags().IdxScheme(), ri.Flags().IsView()
	)

	if (dScheme == lScheme) && (lScheme == rScheme) && !dIsView && !lIsView && !rIsView {
		// Iterate directly on buffers since they have the same memory layout.
		binaryRawEach(db, lb, rb, fn)
		return
	}

	binaryIdxEach(di, li, ri, db, lb, rb, fn)
}

// binaryRawEach is the simplest binary iterator. It walks over all elements in
// destination buffer and calls binary function giving corresponding elements
// from left and right buffers.
func binaryRawEach(db, lb, rb *buffer.Buffer, fn BinaryFunc) {
	leftAt, rightAt := lb.At(), rb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(nil, dst, leftAt(i), rightAt(i))
	})
}

// binaryIdxEach walks over elements in destination buffer pointed by all of its
// index's indices. It calls given binary function with elements from left and
// right buffers. Each element is found by their indexes using destination index
// indices.
func binaryIdxEach(di, li, ri *index.Index, db, lb, rb *buffer.Buffer, fn BinaryFunc) {
	var (
		diAt, liAt, riAt = di.At(), li.At(), ri.At()
		dbAt, lbAt, rbAt = db.At(), lb.At(), rb.At()
	)

	di.Iterate(func(pos []int) {
		fn(pos, dbAt(diAt(pos)), lbAt(liAt(pos)), rbAt(riAt(pos)))
	})
}

// binaryConvert ensures that binary operation function will have all its
// arguments in the exact same type.
func binaryConvert(ddt, ldt, rdt dtype.DType, op func(dtype.DType) BinaryFunc) BinaryFunc {
	var fn = op(ddt)
	if (ddt == ldt) && (ddt == rdt) {
		return fn
	}

	switch ddt {
	case dtype.Bool:
		switch {
		case ldt == dtype.Bool && rdt != dtype.Bool:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsBoolPtr(r)) }
		case ldt != dtype.Bool && rdt == dtype.Bool:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsBoolPtr(l), r) }
		case ldt != dtype.Bool && rdt != dtype.Bool:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsBoolPtr(l), rdt.AsBoolPtr(r))
			}
		}
	case dtype.Int:
		switch {
		case ldt == dtype.Int && rdt != dtype.Int:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsIntPtr(r)) }
		case ldt != dtype.Int && rdt == dtype.Int:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsIntPtr(l), r) }
		case ldt != dtype.Int && rdt != dtype.Int:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsIntPtr(l), rdt.AsIntPtr(r))
			}
		}
	case dtype.Int8:
		switch {
		case ldt == dtype.Int8 && rdt != dtype.Int8:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsInt8Ptr(r)) }
		case ldt != dtype.Int8 && rdt == dtype.Int8:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsInt8Ptr(l), r) }
		case ldt != dtype.Int8 && rdt != dtype.Int8:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsInt8Ptr(l), rdt.AsInt8Ptr(r))
			}
		}
	case dtype.Int16:
		switch {
		case ldt == dtype.Int16 && rdt != dtype.Int16:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsInt16Ptr(r)) }
		case ldt != dtype.Int16 && rdt == dtype.Int16:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsInt16Ptr(l), r) }
		case ldt != dtype.Int16 && rdt != dtype.Int16:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsInt16Ptr(l), rdt.AsInt16Ptr(r))
			}
		}
	case dtype.Int32:
		switch {
		case ldt == dtype.Int32 && rdt != dtype.Int32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsInt32Ptr(r)) }
		case ldt != dtype.Int32 && rdt == dtype.Int32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsInt32Ptr(l), r) }
		case ldt != dtype.Int32 && rdt != dtype.Int32:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsInt32Ptr(l), rdt.AsInt32Ptr(r))
			}
		}
	case dtype.Int64:
		switch {
		case ldt == dtype.Int64 && rdt != dtype.Int64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsInt64Ptr(r)) }
		case ldt != dtype.Int64 && rdt == dtype.Int64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsInt64Ptr(l), r) }
		case ldt != dtype.Int64 && rdt != dtype.Int64:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsInt64Ptr(l), rdt.AsInt64Ptr(r))
			}
		}
	case dtype.Uint:
		switch {
		case ldt == dtype.Uint && rdt != dtype.Uint:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUintPtr(r)) }
		case ldt != dtype.Uint && rdt == dtype.Uint:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUintPtr(l), r) }
		case ldt != dtype.Uint && rdt != dtype.Uint:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUintPtr(l), rdt.AsUintPtr(r))
			}
		}
	case dtype.Uint8:
		switch {
		case ldt == dtype.Uint8 && rdt != dtype.Uint8:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUint8Ptr(r)) }
		case ldt != dtype.Uint8 && rdt == dtype.Uint8:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUint8Ptr(l), r) }
		case ldt != dtype.Uint8 && rdt != dtype.Uint8:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUint8Ptr(l), rdt.AsUint8Ptr(r))
			}
		}
	case dtype.Uint16:
		switch {
		case ldt == dtype.Uint16 && rdt != dtype.Uint16:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUint16Ptr(r)) }
		case ldt != dtype.Uint16 && rdt == dtype.Uint16:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUint16Ptr(l), r) }
		case ldt != dtype.Uint16 && rdt != dtype.Uint16:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUint16Ptr(l), rdt.AsUint16Ptr(r))
			}
		}
	case dtype.Uint32:
		switch {
		case ldt == dtype.Uint32 && rdt != dtype.Uint32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUint32Ptr(r)) }
		case ldt != dtype.Uint32 && rdt == dtype.Uint32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUint32Ptr(l), r) }
		case ldt != dtype.Uint32 && rdt != dtype.Uint32:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUint32Ptr(l), rdt.AsUint32Ptr(r))
			}
		}
	case dtype.Uint64:
		switch {
		case ldt == dtype.Uint64 && rdt != dtype.Uint64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUint64Ptr(r)) }
		case ldt != dtype.Uint64 && rdt == dtype.Uint64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUint64Ptr(l), r) }
		case ldt != dtype.Uint64 && rdt != dtype.Uint64:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUint64Ptr(l), rdt.AsUint64Ptr(r))
			}
		}
	case dtype.Uintptr:
		switch {
		case ldt == dtype.Uintptr && rdt != dtype.Uintptr:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsUintptrPtr(r)) }
		case ldt != dtype.Uintptr && rdt == dtype.Uintptr:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsUintptrPtr(l), r) }
		case ldt != dtype.Uintptr && rdt != dtype.Uintptr:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsUintptrPtr(l), rdt.AsUintptrPtr(r))
			}
		}
	case dtype.Float32:
		switch {
		case ldt == dtype.Float32 && rdt != dtype.Float32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsFloat32Ptr(r)) }
		case ldt != dtype.Float32 && rdt == dtype.Float32:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsFloat32Ptr(l), r) }
		case ldt != dtype.Float32 && rdt != dtype.Float32:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsFloat32Ptr(l), rdt.AsFloat32Ptr(r))
			}
		}
	case dtype.Float64:
		switch {
		case ldt == dtype.Float64 && rdt != dtype.Float64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsFloat64Ptr(r)) }
		case ldt != dtype.Float64 && rdt == dtype.Float64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsFloat64Ptr(l), r) }
		case ldt != dtype.Float64 && rdt != dtype.Float64:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsFloat64Ptr(l), rdt.AsFloat64Ptr(r))
			}
		}
	case dtype.Complex64:
		switch {
		case ldt == dtype.Complex64 && rdt != dtype.Complex64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsComplex64Ptr(r)) }
		case ldt != dtype.Complex64 && rdt == dtype.Complex64:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsComplex64Ptr(l), r) }
		case ldt != dtype.Complex64 && rdt != dtype.Complex64:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsComplex64Ptr(l), rdt.AsComplex64Ptr(r))
			}
		}
	case dtype.Complex128:
		switch {
		case ldt == dtype.Complex128 && rdt != dtype.Complex128:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsComplex128Ptr(r)) }
		case ldt != dtype.Complex128 && rdt == dtype.Complex128:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsComplex128Ptr(l), r) }
		case ldt != dtype.Complex128 && rdt != dtype.Complex128:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsComplex128Ptr(l), rdt.AsComplex128Ptr(r))
			}
		}
	case dtype.String:
		switch {
		case ldt == dtype.String && rdt != dtype.String:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, l, rdt.AsStringPtr(r)) }
		case ldt != dtype.String && rdt == dtype.String:
			return func(pos []int, d, l, r unsafe.Pointer) { fn(pos, d, ldt.AsStringPtr(l), r) }
		case ldt != dtype.String && rdt != dtype.String:
			return func(pos []int, d, l, r unsafe.Pointer) {
				fn(pos, d, ldt.AsStringPtr(l), rdt.AsStringPtr(r))
			}
		}
	}

	panic(errorc.New("unsupported destination type: %q", ddt))
}
