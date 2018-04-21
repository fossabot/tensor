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
	Bool       DType = DType(unsafe.Sizeof(bool(false)))<<32 | 1
	Int        DType = DType(unsafe.Sizeof(int(0)))<<32 | 2
	Int64      DType = DType(unsafe.Sizeof(int64(0)))<<32 | 3
	Uint       DType = DType(unsafe.Sizeof(uint(0)))<<32 | 4
	Uint8      DType = DType(unsafe.Sizeof(uint8(0)))<<32 | 5
	Float64    DType = DType(unsafe.Sizeof(float64(0)))<<32 | 6
	Complex128 DType = DType(unsafe.Sizeof(complex128(0)))<<32 | 7
	String     DType = DType(unsafe.Sizeof(unsafe.Pointer(nil)))<<32 | 8 | flagDynamic
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

// Name returns the name of called data type.
func (dt DType) Name() string {
	if name, ok := dTypeNames[dt]; ok {
		return name
	}

	return fmt.Sprintf("unknown(%x)", uint64(dt))
}

// String satisfies fmt.Stringer interface. It behaves like Name method.
func (dt DType) String() string {
	return dt.Name()
}

var dTypeNames = map[DType]string{
	Bool:       "bool",
	Int:        "int",
	Int64:      "int64",
	Uint:       "uint",
	Uint8:      "uint8",
	Float64:    "float64",
	Complex128: "complex128",
	String:     "string",
}
