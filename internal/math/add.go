package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
)

// Add is a binary function responsible for addition operation.
func Add(dt core.DType) BinaryFunc {
	switch dt {
	case core.Bool:
		return func(d, l, r unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(l) || *(*bool)(r)
		}
	case core.Int:
		return func(d, l, r unsafe.Pointer) {
			*(*int)(d) = *(*int)(l) + *(*int)(r)
		}
	case core.Int64:
		return func(d, l, r unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(l) + *(*int64)(r)
		}
	case core.Float64:
		return func(d, l, r unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(l) + *(*float64)(r)
		}
	case core.String:
		return func(d, l, r unsafe.Pointer) {
			*(*string)(d) = *(*string)(l) + *(*string)(r)
		}
	}

	panic(core.NewError("unsupported type: %q", dt))
}
