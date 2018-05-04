package dtype

import (
	"unsafe"

	"github.com/ppknap/tensor/internal/errorc"
)

// Setraw sets underlying pointer value from src to dst based on provided data
// type. The values pointed by both pointers must be of identical type and the
// size of provided data type must fit the size of underlying pointers type. If
// any of these requirements is not met, the behavior of this function is
// undefined.
func (dt DType) Setraw(dst, src unsafe.Pointer) {
	switch dt.Size() {
	case 1:
		*(*uint8)(dst) = *(*uint8)(src)
		return
	case 2:
		*(*uint16)(dst) = *(*uint16)(src)
		return
	case 4:
		*(*uint32)(dst) = *(*uint32)(src)
		return
	case 8:
		*(*uint64)(dst) = *(*uint64)(src)
		return
	case 16:
		*(*uint64)(dst) = *(*uint64)(src)
		*(*uint64)(unsafe.Pointer(uintptr(dst) + 8)) = *(*uint64)(unsafe.Pointer(uintptr(src) + 8))
		return
	}

	panic(errorc.New("core: unsupported data size: %q(%d)", dt, dt.Size()))
}
