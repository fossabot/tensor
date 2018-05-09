package routine

import (
	stdmath "math"
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Divide is a binary function responsible for division.
func Divide(dt dtype.DType) math.BinaryFunc {
	switch dt {
	case dtype.Bool, dtype.String:
		panic(errorc.New("invalid division on %s type", dt))
	case dtype.Int:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*int)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*int)(d) = *(*int)(l) / v
		}
	case dtype.Int8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*int8)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*int8)(d) = *(*int8)(l) / v
		}
	case dtype.Int16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*int16)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*int16)(d) = *(*int16)(l) / v
		}
	case dtype.Int32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*int32)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*int32)(d) = *(*int32)(l) / v
		}
	case dtype.Int64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*int64)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*int64)(d) = *(*int64)(l) / v
		}
	case dtype.Uint:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uint)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uint)(d) = *(*uint)(l) / v
		}
	case dtype.Uint8:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uint8)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uint8)(d) = *(*uint8)(l) / v
		}
	case dtype.Uint16:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uint16)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uint16)(d) = *(*uint16)(l) / v
		}
	case dtype.Uint32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uint32)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uint32)(d) = *(*uint32)(l) / v
		}
	case dtype.Uint64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uint64)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uint64)(d) = *(*uint64)(l) / v
		}
	case dtype.Uintptr:
		return func(_ []int, d, l, r unsafe.Pointer) {
			v := *(*uintptr)(r)
			if v == 0 {
				panic(errorc.New("division by zero"))
			}
			*(*uintptr)(d) = *(*uintptr)(l) / v
		}
	case dtype.Float32:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*float32)(l), *(*float32)(r)
			if rv != 0 {
				*(*float32)(d) = lv / rv
				return
			}

			switch {
			case lv == 0:
				*(*float32)(d) = float32(stdmath.NaN())
			case lv > 0:
				*(*float32)(d) = float32(stdmath.Inf(1))
			default:
				*(*float32)(d) = float32(stdmath.Inf(-1))
			}
		}
	case dtype.Float64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*float64)(l), *(*float64)(r)
			if rv != 0 {
				*(*float64)(d) = lv / rv
				return
			}

			switch {
			case lv == 0:
				*(*float64)(d) = stdmath.NaN()
			case lv > 0:
				*(*float64)(d) = stdmath.Inf(1)
			default:
				*(*float64)(d) = stdmath.Inf(-1)
			}
		}
	case dtype.Complex64:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*complex64)(l), *(*complex64)(r)
			if real(rv) != 0 {
				*(*complex64)(d) = lv / rv
				return
			}

			var re, im float32
			switch rlv := real(lv); {
			case rlv == 0:
				re = float32(stdmath.NaN())
			case rlv > 0:
				re = float32(stdmath.Inf(1))
			default:
				re = float32(stdmath.Inf(-1))
			}

			switch ilv := imag(lv); {
			case ilv == 0:
				im = float32(stdmath.NaN())
			case ilv > 0:
				im = float32(stdmath.Inf(1))
			default:
				im = float32(stdmath.Inf(-1))
			}

			*(*complex64)(d) = complex(re, im)
		}
	case dtype.Complex128:
		return func(_ []int, d, l, r unsafe.Pointer) {
			lv, rv := *(*complex128)(l), *(*complex128)(r)
			if real(rv) != 0 {
				*(*complex128)(d) = lv / rv
				return
			}

			var re, im float64
			switch rlv := real(lv); {
			case rlv == 0:
				re = stdmath.NaN()
			case rlv > 0:
				re = stdmath.Inf(1)
			default:
				re = stdmath.Inf(-1)
			}

			switch ilv := imag(lv); {
			case ilv == 0:
				im = stdmath.NaN()
			case ilv > 0:
				im = stdmath.Inf(1)
			default:
				im = stdmath.Inf(-1)
			}

			*(*complex128)(d) = complex(re, im)
		}
	}

	panic(errorc.New("unsupported type: %q", dt))
}
