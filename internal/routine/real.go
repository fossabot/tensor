package routine

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/math"
)

// Real is a nullary function that removes imaginary part from a called object.
func Real(dt dtype.DType) math.NullaryFunc {
	switch dt {
	case dtype.Bool, dtype.Int, dtype.Int8, dtype.Int16, dtype.Int32, dtype.Int64,
		dtype.Uint, dtype.Uint8, dtype.Uint16, dtype.Uint32, dtype.Uint64,
		dtype.Uintptr, dtype.Float32, dtype.Float64:
		return func(_ []int, d unsafe.Pointer) { /* no-op */ }
	case dtype.Complex64:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex64)(d) = complex(real(*(*complex64)(d)), 0)
		}
	case dtype.Complex128:
		return func(_ []int, d unsafe.Pointer) {
			*(*complex128)(d) = complex(real(*(*complex128)(d)), 0)
		}
	case dtype.String:
		panic(errorc.New("invalid real part of string type"))
	}

	panic(errorc.New("unsupported type: %q", dt))
}
