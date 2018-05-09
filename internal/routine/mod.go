package routine

import (
	stdmath "math"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Mod is a binary function responsible for computing a reminder of arithmetic
// division.
func Mod(dt dtype.DType) math.BinaryFunc {
	switch dt {
	case dtype.Bool, dtype.String:
		panic(errorc.New("invalid reminder on %s type", dt))
	case dtype.Int:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int)(d) = *(*int)(l) % *(*int)(r)
		}
	case dtype.Int8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int8)(d) = *(*int8)(l) % *(*int8)(r)
		}
	case dtype.Int16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int16)(d) = *(*int16)(l) % *(*int16)(r)
		}
	case dtype.Int32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int32)(d) = *(*int32)(l) % *(*int32)(r)
		}
	case dtype.Int64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*int64)(d) = *(*int64)(l) % *(*int64)(r)
		}
	case dtype.Uint:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint)(d) = *(*uint)(l) % *(*uint)(r)
		}
	case dtype.Uint8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint8)(d) = *(*uint8)(l) % *(*uint8)(r)
		}
	case dtype.Uint16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint16)(d) = *(*uint16)(l) % *(*uint16)(r)
		}
	case dtype.Uint32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint32)(d) = *(*uint32)(l) % *(*uint32)(r)
		}
	case dtype.Uint64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uint64)(d) = *(*uint64)(l) % *(*uint64)(r)
		}
	case dtype.Uintptr:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*uintptr)(d) = *(*uintptr)(l) % *(*uintptr)(r)
		}
	case dtype.Float32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float32)(d) = float32(stdmath.Mod(float64(*(*float32)(l)), float64(*(*float32)(r))))
		}
	case dtype.Float64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			*(*float64)(d) = stdmath.Mod(*(*float64)(l), *(*float64)(r))
		}
	case dtype.Complex64, dtype.Complex128:
		panic(errorc.New("reminder of %s type not supported", dt))
	}

	panic(errorc.New("unsupported type: %q", dt))
}
