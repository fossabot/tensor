package math

import (
	"unsafe"

	"github.com/ppknap/tacvs/internal/core"
)

func Binary(db, lb, rb *core.Buffer, op func(core.DType) core.BinaryFunc) {
	var fn = core.Binary(db.DType(), lb.DType(), rb.DType(), op)

	leftAt, rightAt := lb.At(), rb.At()
	db.Iterate(func(i int, dst unsafe.Pointer) {
		fn(dst, leftAt(i), rightAt(i))
	})
}
