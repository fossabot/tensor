package routine

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Any is a unary function that checks if given argument and destination object
// evaluates to true after logical OR.
func Any(dt dtype.DType) math.UnaryFunc {
	switch dt {
	case dtype.Bool:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(d) || *(*bool)(s)
		}
	case dtype.Int:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*int)(d) == 0 && *(*int)(s) != 0 {
				*(*int)(d) = 1
			}
		}
	case dtype.Int8:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*int8)(d) == 0 && *(*int8)(s) != 0 {
				*(*int8)(d) = 1
			}
		}
	case dtype.Int16:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*int16)(d) == 0 && *(*int16)(s) != 0 {
				*(*int16)(d) = 1
			}
		}
	case dtype.Int32:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*int32)(d) == 0 && *(*int32)(s) != 0 {
				*(*int32)(d) = 1
			}
		}
	case dtype.Int64:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*int64)(d) == 0 && *(*int64)(s) != 0 {
				*(*int64)(d) = 1
			}
		}
	case dtype.Uint:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uint)(d) == 0 && *(*uint)(s) != 0 {
				*(*uint)(d) = 1
			}
		}
	case dtype.Uint8:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uint8)(d) == 0 && *(*uint8)(s) != 0 {
				*(*uint8)(d) = 1
			}
		}
	case dtype.Uint16:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uint16)(d) == 0 && *(*uint16)(s) != 0 {
				*(*uint16)(d) = 1
			}
		}
	case dtype.Uint32:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uint32)(d) == 0 && *(*uint32)(s) != 0 {
				*(*uint32)(d) = 1
			}
		}
	case dtype.Uint64:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uint64)(d) == 0 && *(*uint64)(s) != 0 {
				*(*uint64)(d) = 1
			}
		}
	case dtype.Uintptr:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*uintptr)(d) == 0 && *(*uintptr)(s) != 0 {
				*(*uintptr)(d) = 1
			}
		}
	case dtype.Float32:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*float32)(d) == 0. && *(*float32)(s) != 0. {
				*(*float32)(d) = 1.
			}
		}
	case dtype.Float64:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*float64)(d) == 0. && *(*float64)(s) != 0. {
				*(*float64)(d) = 1.
			}
		}
	case dtype.Complex64:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*complex64)(d) == 0. && *(*complex64)(s) != 0. {
				*(*complex64)(d) = 1.
			}
		}
	case dtype.Complex128:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*complex128)(d) == 0. && *(*complex128)(s) != 0. {
				*(*complex128)(d) = 1.
			}
		}
	case dtype.String:
		return func(_ []int, d, s unsafe.Pointer) {
			if *(*string)(d) == "" && *(*string)(s) != "" {
				*(*string)(d) = "1"
			}
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
