package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
)

// NullaryFunc represents a mathematical operation with no arguments.
type NullaryFunc func(pos []int, d unsafe.Pointer)

// Nullary choses and executes the best strategy to call nullary operation on
// provided buffer with respect to its indexes.
func Nullary(di *index.Index, db *core.Buffer, needsPos bool, op func(core.DType) NullaryFunc) {
	var fn = op(db.DType())

	if needsPos {
		nullaryIdxEach(di, db, fn)
	}

	if !di.Flags().IsView() {
		// Iterate directly on destination buffer.
		nullaryRawEach(db, fn)
		return
	}

	nullaryIdxEach(di, db, fn)
}

// nullaryRawEach is the simplest nullary iterator. It walks over all elements
// in destination buffer and calls nullary function on all its elements.
func nullaryRawEach(db *core.Buffer, fn NullaryFunc) {
	db.Iterate(func(_ int, dst unsafe.Pointer) {
		fn(nil, dst)
	})
}

// nullaryIdxEach walks over elements in destination buffer pointed by all of
// its index's indices. It calls given nullary function on each visited element.
func nullaryIdxEach(di *index.Index, db *core.Buffer, fn NullaryFunc) {
	diAt, dbAt := di.At(), db.At()

	di.Iterate(func(pos []int) {
		fn(pos, dbAt(diAt(pos)))
	})
}
