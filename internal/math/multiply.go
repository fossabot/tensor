package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
)

// Multiply is a binary function responsible for multiplication.
func Multiply(dt dtype.DType) BinaryFunc {
	switch dt {
	case dtype.Bool:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(l) && *(*bool)(r)
		}
	case dtype.Int:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int)(d) = *(*int)(l) * *(*int)(r)
		}
	case dtype.Int8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(l) * *(*int8)(r)
		}
	case dtype.Int16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(l) * *(*int16)(r)
		}
	case dtype.Int32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(l) * *(*int32)(r)
		}
	case dtype.Int64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(l) * *(*int64)(r)
		}
	case dtype.Uint:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(l) * *(*uint)(r)
		}
	case dtype.Uint8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(l) * *(*uint8)(r)
		}
	case dtype.Uint16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(l) * *(*uint16)(r)
		}
	case dtype.Uint32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(l) * *(*uint32)(r)
		}
	case dtype.Uint64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(l) * *(*uint64)(r)
		}
	case dtype.Uintptr:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(l) * *(*uintptr)(r)
		}
	case dtype.Float32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float32)(d) = *(*float32)(l) * *(*float32)(r)
		}
	case dtype.Float64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(l) * *(*float64)(r)
		}
	case dtype.Complex64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*complex64)(d) = *(*complex64)(l) * *(*complex64)(r)
		}
	case dtype.Complex128:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*complex128)(d) = *(*complex128)(l) * *(*complex128)(r)
		}
	case dtype.String:
		panic(errorc.New("invalid multiplication of strings"))
	}

	panic(errorc.New("unsupported type: %q", dt))
}
