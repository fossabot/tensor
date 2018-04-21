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
	Uint64     DType = DType(unsafe.Sizeof(uint64(0)))<<32 | 6
	Float32    DType = DType(unsafe.Sizeof(float32(0)))<<32 | 7
	Float64    DType = DType(unsafe.Sizeof(float64(0)))<<32 | 8
	Complex64  DType = DType(unsafe.Sizeof(complex64(0)))<<32 | 9
	Complex128 DType = DType(unsafe.Sizeof(complex128(0)))<<32 | 10
	String     DType = DType(unsafe.Sizeof(unsafe.Pointer(nil)))<<32 | 11 | flagDynamic
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
	Uint64:     "uint64",
	Float32:    "float32",
	Float64:    "float64",
	Complex64:  "complex64",
	Complex128: "complex128",
	String:     "string",
}
