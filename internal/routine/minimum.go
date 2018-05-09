package routine

import (
	stdmath "math"
	"math/cmplx"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Minimum is a binary function that selects min elements. It propagates NaNs.
func Minimum(dt dtype.DType) math.BinaryFunc {
	switch dt {
	case dtype.Bool:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(l) && *(*bool)(r)
		}
	case dtype.Int:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*int)(l), *(*int)(r); lv < rv {
				*(*int)(d) = lv
			} else {
				*(*int)(d) = rv
			}
		}
	case dtype.Int8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*int8)(l), *(*int8)(r); lv < rv {
				*(*int8)(d) = lv
			} else {
				*(*int8)(d) = rv
			}
		}
	case dtype.Int16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*int16)(l), *(*int16)(r); lv < rv {
				*(*int16)(d) = lv
			} else {
				*(*int16)(d) = rv
			}
		}
	case dtype.Int32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*int32)(l), *(*int32)(r); lv < rv {
				*(*int32)(d) = lv
			} else {
				*(*int32)(d) = rv
			}
		}
	case dtype.Int64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*int64)(l), *(*int64)(r); lv < rv {
				*(*int64)(d) = lv
			} else {
				*(*int64)(d) = rv
			}
		}
	case dtype.Uint:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uint)(l), *(*uint)(r); lv < rv {
				*(*uint)(d) = lv
			} else {
				*(*uint)(d) = rv
			}
		}
	case dtype.Uint8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uint8)(l), *(*uint8)(r); lv < rv {
				*(*uint8)(d) = lv
			} else {
				*(*uint8)(d) = rv
			}
		}
	case dtype.Uint16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uint16)(l), *(*uint16)(r); lv < rv {
				*(*uint16)(d) = lv
			} else {
				*(*uint16)(d) = rv
			}
		}
	case dtype.Uint32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uint32)(l), *(*uint32)(r); lv < rv {
				*(*uint32)(d) = lv
			} else {
				*(*uint32)(d) = rv
			}
		}
	case dtype.Uint64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uint64)(l), *(*uint64)(r); lv < rv {
				*(*uint64)(d) = lv
			} else {
				*(*uint64)(d) = rv
			}
		}
	case dtype.Uintptr:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*uintptr)(l), *(*uintptr)(r); lv < rv {
				*(*uintptr)(d) = lv
			} else {
				*(*uintptr)(d) = rv
			}
		}
	case dtype.Float32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float32)(d) = float32(stdmath.Min(float64(*(*float32)(l)), float64(*(*float32)(r))))
		}
	case dtype.Float64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float64)(d) = stdmath.Min(*(*float64)(l), *(*float64)(r))
		}
	case dtype.Complex64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*complex64)(l), *(*complex64)(r)
			if cmplx.IsNaN(complex128(lv)) {
				*(*complex64)(d) = lv
			} else if cmplx.IsNaN(complex128(rv)) {
				*(*complex64)(d) = rv
			} else if (real(lv) < real(rv)) || (real(lv) == real(rv) && imag(lv) < imag(rv)) {
				*(*complex64)(d) = lv
			} else {
				*(*complex64)(d) = rv
			}
		}
	case dtype.Complex128:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*complex128)(l), *(*complex128)(r)
			if cmplx.IsNaN(lv) {
				*(*complex128)(d) = lv
			} else if cmplx.IsNaN(rv) {
				*(*complex128)(d) = rv
			} else if (real(lv) < real(rv)) || (real(lv) == real(rv) && imag(lv) < imag(rv)) {
				*(*complex128)(d) = lv
			} else {
				*(*complex128)(d) = rv
			}
		}
	case dtype.String:
		return func(_ []int, d, l, r unsafe.Pointer) {
			if lv, rv := *(*string)(l), *(*string)(r); lv < rv {
				*(*string)(d) = lv
			} else {
				*(*string)(d) = rv
			}
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
