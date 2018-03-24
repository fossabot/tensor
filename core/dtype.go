package core

import (
	"unsafe"
)

type DType uint64

const (
	Int64 DType = DType(unsafe.Sizeof(int64(0)))<<32 | 1
)

func (dt DType) Size() uintptr {
	return uintptr(dt >> 32)
}
