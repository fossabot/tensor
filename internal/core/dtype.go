package core

import (
	"fmt"
	"unsafe"
)

type DType uint64

const (
	Bool  DType = DType(unsafe.Sizeof(bool(false)))<<32 | 1
	Int64 DType = DType(unsafe.Sizeof(int64(0)))<<32 | 2
)

var dTypeNames = map[DType]string{
	Bool:  "bool",
	Int64: "int64",
}

func (dt DType) AsStringFunc() func(unsafe.Pointer) string {
	switch dt {
	case Bool:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*bool)(p)) }
	case Int64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int64)(p)) }
	default:
		panic("core: unsupported type: " + dt.String())
	}
}

func (dt DType) Size() uintptr {
	return uintptr(dt >> 32)
}

func (dt DType) String() string {
	if name, ok := dTypeNames[dt]; ok {
		return name
	}

	return fmt.Sprintf("unknown(%x)", uint64(dt))
}
