package routine

import (
	"strings"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Arange is a binary function responsible for spacing elements starting with a
// given value and incrising it by provided step.
func Arange(dt dtype.DType) math.BinaryFunc {
	var n int

	switch dt {
	case dtype.Bool:
		panic(errorc.New("invalid arange on boolean type"))
	case dtype.Int:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*int)(d) = *(*int)(start) + *(*int)(step)*n
			n++
		}
	case dtype.Int8:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(start) + *(*int8)(step)*int8(n)
			n++
		}
	case dtype.Int16:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(start) + *(*int16)(step)*int16(n)
			n++
		}
	case dtype.Int32:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(start) + *(*int32)(step)*int32(n)
			n++
		}
	case dtype.Int64:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(start) + *(*int64)(step)*int64(n)
			n++
		}
	case dtype.Uint:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(start) + *(*uint)(step)*uint(n)
			n++
		}
	case dtype.Uint8:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(start) + *(*uint8)(step)*uint8(n)
			n++
		}
	case dtype.Uint16:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(start) + *(*uint16)(step)*uint16(n)
			n++
		}
	case dtype.Uint32:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(start) + *(*uint32)(step)*uint32(n)
			n++
		}
	case dtype.Uint64:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(start) + *(*uint64)(step)*uint64(n)
			n++
		}
	case dtype.Uintptr:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(start) + *(*uintptr)(step)*uintptr(n)
			n++
		}
	case dtype.Float32:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*float32)(d) = *(*float32)(start) + *(*float32)(step)*float32(n)
			n++
		}
	case dtype.Float64:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*float64)(d) = *(*float64)(start) + *(*float64)(step)*float64(n)
			n++
		}
	case dtype.Complex64:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*complex64)(d) = *(*complex64)(start) + *(*complex64)(step)*complex(float32(n), 0)
			n++
		}
	case dtype.Complex128:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*complex128)(d) = *(*complex128)(start) + *(*complex128)(step)*complex(float64(n), 0)
			n++
		}
	case dtype.String:
		return func(_ []int, d, start, step unsafe.Pointer) {
			*(*string)(d) = *(*string)(start) + strings.Repeat(*(*string)(step), n)
			n++
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
