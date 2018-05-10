package routine

import (
	stdmath "math"
	"math/cmplx"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Log is a nullary function that computes natural logarithm of a called object.
func Log(dt dtype.DType) math.NullaryFunc {
	switch dt {
	case dtype.Bool, dtype.String:
		panic(errorc.New("invalid natural logarithm on %s type", dt))
	case dtype.Int:
		return func(_ []int, d unsafe.Pointer) {
			*(*int)(d) = int(stdmath.Round(stdmath.Log(float64(*(*int)(d)))))
		}
	case dtype.Int8:
		return func(_ []int, d unsafe.Pointer) {
			*(*int8)(d) = int8(stdmath.Round(stdmath.Log(float64(*(*int8)(d)))))
		}
	case dtype.Int16:
		return func(_ []int, d unsafe.Pointer) {
			*(*int16)(d) = int16(stdmath.Round(stdmath.Log(float64(*(*int16)(d)))))
		}
	case dtype.Int32:
		return func(_ []int, d unsafe.Pointer) {
			*(*int32)(d) = int32(stdmath.Round(stdmath.Log(float64(*(*int32)(d)))))
		}
	case dtype.Int64:
		return func(_ []int, d unsafe.Pointer) {
			*(*int64)(d) = int64(stdmath.Round(stdmath.Log(float64(*(*int64)(d)))))
		}
	case dtype.Uint:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint)(d) = uint(stdmath.Round(stdmath.Log(float64(*(*uint)(d)))))
		}
	case dtype.Uint8:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint8)(d) = uint8(stdmath.Round(stdmath.Log(float64(*(*uint8)(d)))))
		}
	case dtype.Uint16:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint16)(d) = uint16(stdmath.Round(stdmath.Log(float64(*(*uint16)(d)))))
		}
	case dtype.Uint32:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint32)(d) = uint32(stdmath.Round(stdmath.Log(float64(*(*uint32)(d)))))
		}
	case dtype.Uint64:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint64)(d) = uint64(stdmath.Round(stdmath.Log(float64(*(*uint64)(d)))))
		}
	case dtype.Uintptr:
		return func(_ []int, d unsafe.Pointer) {
			*(*uintptr)(d) = uintptr(stdmath.Round(stdmath.Log(float64(*(*uintptr)(d)))))
		}
	case dtype.Float32:
		return func(_ []int, d unsafe.Pointer) {
			*(*float32)(d) = float32(stdmath.Log(float64(*(*float32)(d))))
		}
	case dtype.Float64:
		return func(_ []int, d unsafe.Pointer) {
			*(*float64)(d) = stdmath.Log(*(*float64)(d))
		}
	case dtype.Complex64:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex64)(d) = complex64(cmplx.Log(complex128(*(*complex64)(d))))
		}
	case dtype.Complex128:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex128)(d) = cmplx.Log(*(*complex128)(d))
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
