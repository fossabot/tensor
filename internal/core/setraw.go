package core

import (
	"unsafe"
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
	case 2:
		*(*uint16)(dst) = *(*uint16)(src)
	case 4:
		*(*uint32)(dst) = *(*uint32)(src)
	case 8:
		*(*uint64)(dst) = *(*uint64)(src)
	case 16:
		*(*uint64)(dst) = *(*uint64)(src)
		*(*uint64)(unsafe.Pointer(uintptr(dst) + 8)) = *(*uint64)(unsafe.Pointer(uintptr(src) + 8))
	}

	panic(NewError("core: unsupported data size: %q", dt))
}
