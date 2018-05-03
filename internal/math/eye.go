package math

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/core"
)

// Eye is a nullary function responsible for creating a tensor which has ones
// on its main diagonal and zeroes elsewere.
func Eye(dt core.DType) NullaryFunc {
	isDiag := func(pos []int) bool {
		if len(pos) == 0 {
			return true
		}

		if len(pos) == 1 {
			return pos[0] == 0
		}

		for i := 1; i < len(pos); i++ {
			if pos[i-1] != pos[i] {
				return false
			}
		}

		return true
	}

	switch dt {
	case core.Bool:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*bool)(d) = true
				return
			}
			*(*bool)(d) = false
		}
	case core.Int:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*int)(d) = 1
				return
			}
			*(*int)(d) = 0
		}
	case core.Int8:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*int8)(d) = 1
				return
			}
			*(*int8)(d) = 0
		}
	case core.Int16:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*int16)(d) = 1
				return
			}
			*(*int16)(d) = 0
		}
	case core.Int32:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*int32)(d) = 1
				return
			}
			*(*int32)(d) = 0
		}
	case core.Int64:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*int64)(d) = 1
				return
			}
			*(*int64)(d) = 0
		}
	case core.Uint:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uint)(d) = 1
				return
			}
			*(*uint)(d) = 0
		}
	case core.Uint8:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uint8)(d) = 1
				return
			}
			*(*uint8)(d) = 0
		}
	case core.Uint16:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uint16)(d) = 1
				return
			}
			*(*uint16)(d) = 0
		}
	case core.Uint32:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uint32)(d) = 1
				return
			}
			*(*uint32)(d) = 0
		}
	case core.Uint64:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uint64)(d) = 1
				return
			}
			*(*uint64)(d) = 0
		}
	case core.Uintptr:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*uintptr)(d) = 1
				return
			}
			*(*uintptr)(d) = 0
		}
	case core.Float32:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*float32)(d) = 1.
				return
			}
			*(*float32)(d) = 0.
		}
	case core.Float64:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*float64)(d) = 1.
				return
			}
			*(*float64)(d) = 0.
		}
	case core.Complex64:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*complex64)(d) = 1.
				return
			}
			*(*complex64)(d) = 0.
		}
	case core.Complex128:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*complex128)(d) = 1.
				return
			}
			*(*complex128)(d) = 0.
		}
	case core.String:
		return func(pos []int, d unsafe.Pointer) {
			if isDiag(pos) {
				*(*string)(d) = "1"
				return
			}
			*(*string)(d) = ""
		}
	}

	panic(core.NewError("unsupported type: %q", dt))
}
