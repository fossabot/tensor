package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
)

// Fill is a unary function responsible for value assignment.
func Fill(dt core.DType) UnaryFunc {
	switch dt {
	case core.Bool:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*bool)(d) = *(*bool)(s)
		}
	case core.Int:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int)(d) = *(*int)(s)
		}
	case core.Int8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(s)
		}
	case core.Int16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(s)
		}
	case core.Int32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(s)
		}
	case core.Int64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(s)
		}
	case core.Uint:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(s)
		}
	case core.Uint8:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(s)
		}
	case core.Uint16:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(s)
		}
	case core.Uint32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(s)
		}
	case core.Uint64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(s)
		}
	case core.Uintptr:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(s)
		}
	case core.Float32:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float32)(d) = *(*float32)(s)
		}
	case core.Float64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(s)
		}
	case core.Complex64:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex64)(d) = *(*complex64)(s)
		}
	case core.Complex128:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*complex128)(d) = *(*complex128)(s)
		}
	case core.String:
		return func(_ []int, d, s unsafe.Pointer) {
			*(*string)(d) = *(*string)(s)
		}
	}

	panic(core.NewError("unsupported type: %q", dt))
}
