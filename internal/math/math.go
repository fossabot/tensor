package math

import (
	"github.com/ppknap/tacvs/internal/core"
	"github.com/ppknap/tacvs/internal/index"
)

// EWArgShape returns the required destination shape for element-wise operation
// to succeed. It panics when such operation is not possible.
func EWArgShape(li, ri *index.Index) (shape []int) {
	switch lsz, rsz := li.Size(), ri.Size(); {
	case lsz > 1 && rsz > 1 && li.EqShape(ri):
		// Higher rank tensors operation. Only equal shapes are allowed.
		return li.Shape()
	case lsz == 1 && rsz == 1:
		// Scalar operation.
		return nil
	case lsz == 1:
		// Scalar to higher rank tensor operation. Scalar as left argument.
		return ri.Shape()
	case rsz == 1:
		// Scalar to higher rank tensor operation. Scalar as right argument.
		return li.Shape()
	}

	panic(core.NewError("invalid element-wise op. on %v and %v", li.Shape(), ri.Shape()))
}
