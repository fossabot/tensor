package routine

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Imag is a nullary function that removes real part from a called object.
func Imag(dt dtype.DType) math.NullaryFunc {
	switch dt {
	case dtype.Bool:
		return func(_ []int, d unsafe.Pointer) {
			*(*bool)(d) = false
		}
	case dtype.Int:
		return func(_ []int, d unsafe.Pointer) {
			*(*int)(d) = 0
		}
	case dtype.Int8:
		return func(_ []int, d unsafe.Pointer) {
			*(*int8)(d) = 0
		}
	case dtype.Int16:
		return func(_ []int, d unsafe.Pointer) {
			*(*int16)(d) = 0
		}
	case dtype.Int32:
		return func(_ []int, d unsafe.Pointer) {
			*(*int32)(d) = 0
		}
	case dtype.Int64:
		return func(_ []int, d unsafe.Pointer) {
			*(*int64)(d) = 0
		}
	case dtype.Uint:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint)(d) = 0
		}
	case dtype.Uint8:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint8)(d) = 0
		}
	case dtype.Uint16:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint16)(d) = 0
		}
	case dtype.Uint32:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint32)(d) = 0
		}
	case dtype.Uint64:
		return func(_ []int, d unsafe.Pointer) {
			*(*uint64)(d) = 0
		}
	case dtype.Uintptr:
		return func(_ []int, d unsafe.Pointer) {
			*(*uintptr)(d) = 0
		}
	case dtype.Float32:
		return func(_ []int, d unsafe.Pointer) {
			*(*float32)(d) = 0.
		}
	case dtype.Float64:
		return func(_ []int, d unsafe.Pointer) {
			*(*float64)(d) = 0.
		}
	case dtype.Complex64:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex64)(d) = complex(0, imag(*(*complex64)(d)))
		}
	case dtype.Complex128:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex128)(d) = complex(0, imag(*(*complex128)(d)))
		}
	case dtype.String:
		panic(errorc.New("invalid imaginary part of string type"))
	}

	panic(errorc.New("unsupported type: %q", dt))
}
