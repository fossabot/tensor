package routine

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Linspace is a binary function responsible for filling a tensor with evenly
// spaced values from start to end.
func Linspace(size int) func(dt dtype.DType) math.BinaryFunc {
	return func(dt dtype.DType) math.BinaryFunc {
		if size == 1 {
			ufn := Fill(dt)
			return func(pos []int, d, start, _ unsafe.Pointer) {
				ufn(pos, d, start)
			}
		}

		var n float64

		switch dt {
		case dtype.Bool, dtype.String:
			panic(errorc.New("invalid linspace on %s type", dt))
		case dtype.Int:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*int)(d) = *(*int)(start) + int((float64(*(*int)(end))-float64(*(*int)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Int8:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*int8)(d) = *(*int8)(start) + int8((float64(*(*int8)(end))-float64(*(*int8)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Int16:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*int16)(d) = *(*int16)(start) + int16((float64(*(*int16)(end))-float64(*(*int16)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Int32:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*int32)(d) = *(*int32)(start) + int32((float64(*(*int32)(end))-float64(*(*int32)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Int64:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*int64)(d) = *(*int64)(start) + int64((float64(*(*int64)(end))-float64(*(*int64)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uint:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uint)(d) = *(*uint)(start) + uint((float64(*(*uint)(end))-float64(*(*uint)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uint8:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uint8)(d) = *(*uint8)(start) + uint8((float64(*(*uint8)(end))-float64(*(*uint8)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uint16:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uint16)(d) = *(*uint16)(start) + uint16((float64(*(*uint16)(end))-float64(*(*uint16)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uint32:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uint32)(d) = *(*uint32)(start) + uint32((float64(*(*uint32)(end))-float64(*(*uint32)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uint64:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uint64)(d) = *(*uint64)(start) + uint64((float64(*(*uint64)(end))-float64(*(*uint64)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Uintptr:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*uintptr)(d) = *(*uintptr)(start) + uintptr((float64(*(*uintptr)(end))-float64(*(*uintptr)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Float32:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*float32)(d) = *(*float32)(start) + float32((float64(*(*float32)(end))-float64(*(*float32)(start)))*n/float64(size-1))
				n++
			}
		case dtype.Float64:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*float64)(d) = *(*float64)(start) + (*(*float64)(end)-*(*float64)(start))*n/float64(size-1)
				n++
			}
		case dtype.Complex64:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*complex64)(d) = *(*complex64)(start) + complex64((complex128(*(*complex64)(end))-complex128(*(*complex64)(start)))*complex(float64(n), 0)/complex(float64(size-1), 0))
				n++
			}
		case dtype.Complex128:
			return func(_ []int, d, start, end unsafe.Pointer) {
				*(*complex128)(d) = *(*complex128)(start) + (*(*complex128)(end)-*(*complex128)(start))*complex(float64(n), 0)/complex(float64(size-1), 0)
				n++
			}
		}

		panic(errorc.New("unsupported type: %q", dt))
	}
}
