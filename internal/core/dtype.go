package core

import (
	"fmt"
	"unsafe"
)

// DType represents supported data types. Its value is stored as a bitmask with
// information about data size and memory properties. One should not store nor
// depend on any values of this type since their format is unstable.
type DType uint64

const (
	Bool   DType = DType(unsafe.Sizeof(bool(false)))<<32 | 1
	Int    DType = DType(unsafe.Sizeof(int(0)))<<32 | 2
	Int64  DType = DType(unsafe.Sizeof(int64(0)))<<32 | 3
	String DType = DType(unsafe.Sizeof(unsafe.Pointer(nil)))<<32 | 4 | flagDynamic
)

const (
	// flagDynamic indicates that the type is not an owner of the data it
	// represents.
	flagDynamic DType = 1<<8 + iota
)

// Size returns the size in bytes of provided type.
func (dt DType) Size() uintptr { return uintptr(dt >> 32) }

// IsDynamic returns true when data type does not own all its data.
func (dt DType) IsDynamic() bool { return dt&flagDynamic != 0 }

// String returns the type name.
func (dt DType) String() string {
	if name, ok := dTypeNames[dt]; ok {
		return name
	}

	return fmt.Sprintf("unknown(%x)", uint64(dt))
}

var dTypeNames = map[DType]string{
	Bool:   "bool",
	Int:    "int",
	Int64:  "int64",
	String: "string",
}
