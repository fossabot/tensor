package routine

import (
	stdmath "math"
	"math/cmplx"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Exp is a nullary function that computes exponential of a called object.
func Exp(dt dtype.DType) math.NullaryFunc {
	switch dt {
	case dtype.Bool, dtype.Int, dtype.Int8, dtype.Int16, dtype.Int32, dtype.Int64,
		dtype.Uint, dtype.Uint8, dtype.Uint16, dtype.Uint32, dtype.Uint64,
		dtype.Uintptr:
		return func(_ []int, d unsafe.Pointer) {
			panic(errorc.New("invalid exponential on %s type", dt))
		}
	case dtype.Float32:
		return func(pos []int, d unsafe.Pointer) {
			*(*float32)(d) = float32(stdmath.Exp(float64(*(*float32)(d))))
		}
	case dtype.Float64:
		return func(pos []int, d unsafe.Pointer) {
			*(*float64)(d) = stdmath.Exp(*(*float64)(d))
		}
	case dtype.Complex64:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex64)(d) = complex64(cmplx.Exp(complex128(*(*complex64)(d))))
		}
	case dtype.Complex128:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex128)(d) = cmplx.Exp(*(*complex128)(d))
		}
	case dtype.String:
		panic(errorc.New("invalid exponential on %s type", dt))
	}

	panic(errorc.New("unsupported type: %q", dt))
}
