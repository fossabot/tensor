package math

import (
	"unsafe"

	"github.com/ppknap/tacvs/internal/core"
	"github.com/ppknap/tacvs/internal/index"
)

// BinaryFunc represents a mathematical operation that combines two operands and
// produce another element of the field that should be saved in d pointer. Left
// and right operands should not be modified.
type BinaryFunc func(d, l, r unsafe.Pointer)

// BinaryRawEach is the simplest binary iterator. It walks over all elements in
// destination buffer and calls binary function giving corresponding elements
// from left and right buffers.
func BinaryRawEach(db, lb, rb *core.Buffer, op func(core.DType) BinaryFunc) {
	var fn = binaryConvert(db.DType(), lb.DType(), rb.DType(), op)

	leftAt, rightAt := lb.At(), rb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(dst, leftAt(i), rightAt(i))
	})
}

// BinaryIdxEach walks over elements in destination buffer pointed by all of its
// index's indices. It calls produced binary function with elements from left
// and right buffers. Each element is found by their indexes using destination
// index indices.
func BinaryIdxEach(di, li, ri *index.Index, db, lb, rb *core.Buffer, op func(core.DType) BinaryFunc) {
	var fn = binaryConvert(db.DType(), lb.DType(), rb.DType(), op)

	dstAt, leftAt, rightAt := db.At(), lb.At(), rb.At()
	di.Iterate(func(pos []int) {
		fn(dstAt(di.At(pos)), leftAt(li.At(pos)), rightAt(ri.At(pos)))
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

	panic("core: unsupported destination type: " + ddt.String())
}
