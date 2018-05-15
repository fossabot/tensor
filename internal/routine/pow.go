package routine

import (
	stdmath "math"
	"math/cmplx"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Pow is a unary function that rises destination object to power value taken
// from source argument.
func Pow(dt dtype.DType) math.UnaryFunc {
	switch dt {
	case dtype.Bool, dtype.String:
		panic(errorc.New("invalid power on %s type", dt))
	case dtype.Int:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int)(d) = int(stdmath.Round(stdmath.Pow(float64(*(*int)(d)), float64(*(*int)(s)))))
		}
	case dtype.Int8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int8)(d) = int8(stdmath.Round(stdmath.Pow(float64(*(*int8)(d)), float64(*(*int8)(s)))))
		}
	case dtype.Int16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int16)(d) = int16(stdmath.Round(stdmath.Pow(float64(*(*int16)(d)), float64(*(*int16)(s)))))
		}
	case dtype.Int32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int32)(d) = int32(stdmath.Round(stdmath.Pow(float64(*(*int32)(d)), float64(*(*int32)(s)))))
		}
	case dtype.Int64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int64)(d) = int64(stdmath.Round(stdmath.Pow(float64(*(*int64)(d)), float64(*(*int64)(s)))))
		}
	case dtype.Uint:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint)(d) = uint(stdmath.Round(stdmath.Pow(float64(*(*uint)(d)), float64(*(*uint)(s)))))
		}
	case dtype.Uint8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint8)(d) = uint8(stdmath.Round(stdmath.Pow(float64(*(*uint8)(d)), float64(*(*uint8)(s)))))
		}
	case dtype.Uint16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint16)(d) = uint16(stdmath.Round(stdmath.Pow(float64(*(*uint16)(d)), float64(*(*uint16)(s)))))
		}
	case dtype.Uint32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint32)(d) = uint32(stdmath.Round(stdmath.Pow(float64(*(*uint32)(d)), float64(*(*uint32)(s)))))
		}
	case dtype.Uint64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint64)(d) = uint64(stdmath.Round(stdmath.Pow(float64(*(*uint64)(d)), float64(*(*uint64)(s)))))
		}
	case dtype.Uintptr:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uintptr)(d) = uintptr(stdmath.Round(stdmath.Pow(float64(*(*uintptr)(d)), float64(*(*uintptr)(s)))))
		}
	case dtype.Float32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float32)(d) = float32(stdmath.Pow(float64(*(*float32)(d)), float64(*(*float32)(s))))
		}
	case dtype.Float64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float64)(d) = stdmath.Pow(*(*float64)(d), *(*float64)(s))
		}
	case dtype.Complex64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex64)(d) = complex64(cmplx.Pow(complex128(*(*complex64)(d)), complex128(*(*complex64)(s))))
		}
	case dtype.Complex128:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex128)(d) = cmplx.Pow(*(*complex128)(d), *(*complex128)(s))
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
