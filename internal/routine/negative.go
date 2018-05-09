package routine

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Negative is a nullary function that computes numerical negative of called
// object.
func Negative(dt dtype.DType) math.NullaryFunc {
	switch dt {
	case dtype.Bool, dtype.String:
		panic(errorc.New("invalid numerical negation on %s type", dt))
	case dtype.Int:
		return func(pos []int, d unsafe.Pointer) {
			*(*int)(d) = -(*(*int)(d))
		}
	case dtype.Int8:
		return func(pos []int, d unsafe.Pointer) {
			*(*int8)(d) = -(*(*int8)(d))
		}
	case dtype.Int16:
		return func(pos []int, d unsafe.Pointer) {
			*(*int16)(d) = -(*(*int16)(d))
		}
	case dtype.Int32:
		return func(pos []int, d unsafe.Pointer) {
			*(*int32)(d) = -(*(*int32)(d))
		}
	case dtype.Int64:
		return func(pos []int, d unsafe.Pointer) {
			*(*int64)(d) = -(*(*int64)(d))
		}
	case dtype.Uint:
		return func(pos []int, d unsafe.Pointer) {
			*(*uint)(d) = -(*(*uint)(d))
		}
	case dtype.Uint8:
		return func(pos []int, d unsafe.Pointer) {
			*(*uint8)(d) = -(*(*uint8)(d))
		}
	case dtype.Uint16:
		return func(pos []int, d unsafe.Pointer) {
			*(*uint16)(d) = -(*(*uint16)(d))
		}
	case dtype.Uint32:
		return func(pos []int, d unsafe.Pointer) {
			*(*uint32)(d) = -(*(*uint32)(d))
		}
	case dtype.Uint64:
		return func(pos []int, d unsafe.Pointer) {
			*(*uint64)(d) = -(*(*uint64)(d))
		}
	case dtype.Uintptr:
		return func(pos []int, d unsafe.Pointer) {
			*(*uintptr)(d) = -(*(*uintptr)(d))
		}
	case dtype.Float32:
		return func(pos []int, d unsafe.Pointer) {
			*(*float32)(d) = -(*(*float32)(d))
		}
	case dtype.Float64:
		return func(pos []int, d unsafe.Pointer) {
			*(*float64)(d) = -(*(*float64)(d))
		}
	case dtype.Complex64:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex64)(d) = -(*(*complex64)(d))
		}
	case dtype.Complex128:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex128)(d) = -(*(*complex128)(d))
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
