package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
)

// Multiply is a binary function responsible for multiplication.
func Multiply(dt core.DType) BinaryFunc {
	switch dt {
	case core.Bool:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(l) && *(*bool)(r)
		}
	case core.Int:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int)(d) = *(*int)(l) * *(*int)(r)
		}
	case core.Int8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(l) * *(*int8)(r)
		}
	case core.Int16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(l) * *(*int16)(r)
		}
	case core.Int32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(l) * *(*int32)(r)
		}
	case core.Int64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(l) * *(*int64)(r)
		}
	case core.Uint:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(l) * *(*uint)(r)
		}
	case core.Uint8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(l) * *(*uint8)(r)
		}
	case core.Uint16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(l) * *(*uint16)(r)
		}
	case core.Uint32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(l) * *(*uint32)(r)
		}
	case core.Uint64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(l) * *(*uint64)(r)
		}
	case core.Uintptr:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(l) * *(*uintptr)(r)
		}
	case core.Float32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float32)(d) = *(*float32)(l) * *(*float32)(r)
		}
	case core.Float64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(l) * *(*float64)(r)
		}
	case core.Complex64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*complex64)(d) = *(*complex64)(l) * *(*complex64)(r)
		}
	case core.Complex128:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*complex128)(d) = *(*complex128)(l) * *(*complex128)(r)
		}
	case core.String:
		panic(core.NewError("invalid multiplication of strings"))
	}

	panic(core.NewError("unsupported type: %q", dt))
}
