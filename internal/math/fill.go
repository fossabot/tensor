package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
)

// Fill is a unary function responsible for value assignment.
func Fill(dt dtype.DType) UnaryFunc {
	switch dt {
	case dtype.Bool:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(s)
		}
	case dtype.Int:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int)(d) = *(*int)(s)
		}
	case dtype.Int8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(s)
		}
	case dtype.Int16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(s)
		}
	case dtype.Int32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(s)
		}
	case dtype.Int64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(s)
		}
	case dtype.Uint:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(s)
		}
	case dtype.Uint8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(s)
		}
	case dtype.Uint16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(s)
		}
	case dtype.Uint32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(s)
		}
	case dtype.Uint64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(s)
		}
	case dtype.Uintptr:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(s)
		}
	case dtype.Float32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float32)(d) = *(*float32)(s)
		}
	case dtype.Float64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(s)
		}
	case dtype.Complex64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex64)(d) = *(*complex64)(s)
		}
	case dtype.Complex128:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex128)(d) = *(*complex128)(s)
		}
	case dtype.String:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*string)(d) = *(*string)(s)
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
