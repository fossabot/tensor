package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
)

// BinaryFunc represents a mathematical operation that combines two operands and
// produce another element of the field that should be saved in d pointer. Left
// and right operands should not be modified.
type BinaryFunc func(d, l, r unsafe.Pointer)

// Binary choses and executes the best strategy to call binary operator on
// provided buffers with respect to their indexes.
func Binary(di, li, ri *index.Index, db, lb, rb *core.Buffer, op func(core.DType) BinaryFunc) {
	var fn = binaryConvert(db.DType(), lb.DType(), rb.DType(), op)

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
func binaryRawEach(db, lb, rb *core.Buffer, fn BinaryFunc) {
	leftAt, rightAt := lb.At(), rb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(dst, leftAt(i), rightAt(i))
	})
}

// binaryIdxEach walks over elements in destination buffer pointed by all of its
// index's indices. It calls produced binary function with elements from left
// and right buffers. Each element is found by their indexes using destination
// index indices.
func binaryIdxEach(di, li, ri *index.Index, db, lb, rb *core.Buffer, fn BinaryFunc) {
	var (
		diAt, liAt, riAt = di.At(), li.At(), ri.At()
		dbAt, lbAt, rbAt = db.At(), lb.At(), rb.At()
	)

	di.Iterate(func(pos []int) {
		fn(dbAt(diAt(pos)), lbAt(liAt(pos)), rbAt(riAt(pos)))
	})
}

// binaryConvert ensures that binary operation function will have all its
// arguments in the exact same type.
func binaryConvert(ddt, ldt, rdt core.DType, op func(core.DType) BinaryFunc) BinaryFunc {
	var fn = op(ddt)
	if (ddt == ldt) && (ddt == rdt) {
		return fn
	}

	switch ddt {
	case core.Bool:
		switch {
		case ldt == core.Bool && rdt != core.Bool:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsBoolPtr(r)) }
		case ldt != core.Bool && rdt == core.Bool:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsBoolPtr(l), r) }
		case ldt != core.Bool && rdt != core.Bool:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsBoolPtr(l), rdt.AsBoolPtr(r))
			}
		}
	case core.Int:
		switch {
		case ldt == core.Int && rdt != core.Int:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsIntPtr(r)) }
		case ldt != core.Int && rdt == core.Int:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsIntPtr(l), r) }
		case ldt != core.Int && rdt != core.Int:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsIntPtr(l), rdt.AsIntPtr(r))
			}
		}
	case core.Int64:
		switch {
		case ldt == core.Int64 && rdt != core.Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsInt64Ptr(r)) }
		case ldt != core.Int64 && rdt == core.Int64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsInt64Ptr(l), r) }
		case ldt != core.Int64 && rdt != core.Int64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsInt64Ptr(l), rdt.AsInt64Ptr(r))
			}
		}
	case core.Uint:
		switch {
		case ldt == core.Uint && rdt != core.Uint:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsUintPtr(r)) }
		case ldt != core.Uint && rdt == core.Uint:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsUintPtr(l), r) }
		case ldt != core.Uint && rdt != core.Uint:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsUintPtr(l), rdt.AsUintPtr(r))
			}
		}
	case core.Uint8:
		switch {
		case ldt == core.Uint8 && rdt != core.Uint8:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsUint8Ptr(r)) }
		case ldt != core.Uint8 && rdt == core.Uint8:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsUint8Ptr(l), r) }
		case ldt != core.Uint8 && rdt != core.Uint8:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsUint8Ptr(l), rdt.AsUint8Ptr(r))
			}
		}
	case core.Uint16:
		switch {
		case ldt == core.Uint16 && rdt != core.Uint16:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsUint16Ptr(r)) }
		case ldt != core.Uint16 && rdt == core.Uint16:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsUint16Ptr(l), r) }
		case ldt != core.Uint16 && rdt != core.Uint16:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsUint16Ptr(l), rdt.AsUint16Ptr(r))
			}
		}
	case core.Uint32:
		switch {
		case ldt == core.Uint32 && rdt != core.Uint32:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsUint32Ptr(r)) }
		case ldt != core.Uint32 && rdt == core.Uint32:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsUint32Ptr(l), r) }
		case ldt != core.Uint32 && rdt != core.Uint32:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsUint32Ptr(l), rdt.AsUint32Ptr(r))
			}
		}
	case core.Uint64:
		switch {
		case ldt == core.Uint64 && rdt != core.Uint64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsUint64Ptr(r)) }
		case ldt != core.Uint64 && rdt == core.Uint64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsUint64Ptr(l), r) }
		case ldt != core.Uint64 && rdt != core.Uint64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsUint64Ptr(l), rdt.AsUint64Ptr(r))
			}
		}
	case core.Float32:
		switch {
		case ldt == core.Float32 && rdt != core.Float32:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsFloat32Ptr(r)) }
		case ldt != core.Float32 && rdt == core.Float32:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsFloat32Ptr(l), r) }
		case ldt != core.Float32 && rdt != core.Float32:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsFloat32Ptr(l), rdt.AsFloat32Ptr(r))
			}
		}
	case core.Float64:
		switch {
		case ldt == core.Float64 && rdt != core.Float64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsFloat64Ptr(r)) }
		case ldt != core.Float64 && rdt == core.Float64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsFloat64Ptr(l), r) }
		case ldt != core.Float64 && rdt != core.Float64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsFloat64Ptr(l), rdt.AsFloat64Ptr(r))
			}
		}
	case core.Complex64:
		switch {
		case ldt == core.Complex64 && rdt != core.Complex64:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsComplex64Ptr(r)) }
		case ldt != core.Complex64 && rdt == core.Complex64:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsComplex64Ptr(l), r) }
		case ldt != core.Complex64 && rdt != core.Complex64:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsComplex64Ptr(l), rdt.AsComplex64Ptr(r))
			}
		}
	case core.Complex128:
		switch {
		case ldt == core.Complex128 && rdt != core.Complex128:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsComplex128Ptr(r)) }
		case ldt != core.Complex128 && rdt == core.Complex128:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsComplex128Ptr(l), r) }
		case ldt != core.Complex128 && rdt != core.Complex128:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsComplex128Ptr(l), rdt.AsComplex128Ptr(r))
			}
		}
	case core.String:
		switch {
		case ldt == core.String && rdt != core.String:
			return func(d, l, r unsafe.Pointer) { fn(d, l, rdt.AsStringPtr(r)) }
		case ldt != core.String && rdt == core.String:
			return func(d, l, r unsafe.Pointer) { fn(d, ldt.AsStringPtr(l), r) }
		case ldt != core.String && rdt != core.String:
			return func(d, l, r unsafe.Pointer) {
				fn(d, ldt.AsStringPtr(l), rdt.AsStringPtr(r))
			}
		}
	}

	panic(core.NewError("unsupported destination type: %q", ddt))
}
