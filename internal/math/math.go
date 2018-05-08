package math

import (
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
)

// EWArgShape returns the required destination shape for element-wise operation
// to succeed. It panics when such operation is not possible.
//
// If commutativity argument is true, the indexes are indicated as commutative
// that is: li (+) ri <==> ri (+) li.
func EWArgShape(li, ri *index.Index, commutativity bool) (shape []int) {
	switch {
	case li == nil && ri == nil:
		return nil
	case li == nil:
		return ri.Shape()
	case ri == nil:
		return li.Shape()
	}

	switch lsz, rsz := li.Size(), ri.Size(); {
	case lsz > 1 && rsz > 1 && li.EqShape(ri):
		// Higher rank tensors operation. Only equal shapes are allowed.
		return li.Shape()
	case lsz == 1 && rsz == 1:
		// Scalar operation.
		return li.MergeShape(ri)
	case lsz == 1 && commutativity:
		// Scalar to higher rank tensor operation. Scalar as left argument.
		return li.MergeShape(ri)
	case rsz == 1:
		// Scalar to higher rank tensor operation. Scalar as right argument.
		return li.MergeShape(ri)
	}

	panic(errorc.New("invalid element-wise op. on %v and %v", li.Shape(), ri.Shape()))
}
