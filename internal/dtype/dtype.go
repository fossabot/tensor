package dtype

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ppknap/tensor/internal/errorc"
)

// DType represents supported data types. Its value is stored as a bitmask with
// information about data size and memory properties. One should not store nor
// depend on any values of this type since their format is unstable.
type DType uint64

// Group of all supported data types.
const (
	Bool       DType = DType(unsafe.Sizeof(bool(false)))<<32 | 1
	Int        DType = DType(unsafe.Sizeof(int(0)))<<32 | 2
	Int8       DType = DType(unsafe.Sizeof(int8(0)))<<32 | 3
	Int16      DType = DType(unsafe.Sizeof(int16(0)))<<32 | 4
	Int32      DType = DType(unsafe.Sizeof(int32(0)))<<32 | 5
	Int64      DType = DType(unsafe.Sizeof(int64(0)))<<32 | 6
	Uint       DType = DType(unsafe.Sizeof(uint(0)))<<32 | 7
	Uint8      DType = DType(unsafe.Sizeof(uint8(0)))<<32 | 8
	Uint16     DType = DType(unsafe.Sizeof(uint16(0)))<<32 | 9
	Uint32     DType = DType(unsafe.Sizeof(uint32(0)))<<32 | 10
	Uint64     DType = DType(unsafe.Sizeof(uint64(0)))<<32 | 11
	Uintptr    DType = DType(unsafe.Sizeof(uintptr(0)))<<32 | 12
	Float32    DType = DType(unsafe.Sizeof(float32(0)))<<32 | 13
	Float64    DType = DType(unsafe.Sizeof(float64(0)))<<32 | 14
	Complex64  DType = DType(unsafe.Sizeof(complex64(0)))<<32 | 15
	Complex128 DType = DType(unsafe.Sizeof(complex128(0)))<<32 | 16
	String     DType = DType(unsafe.Sizeof(unsafe.Pointer(nil)))<<32 | 17 | flagDynamic
)

const (
	// flagDynamic indicates that the type is not an owner of the data it
	// represents.
	flagDynamic DType = 1<<8 + iota
)

// FromKind transforms reflect Kind value to coresponding DType object.
func FromKind(k reflect.Kind) DType {
	if dt, ok := kindDType[k]; ok {
		return dt
	}

	panic(errorc.New("unsupported reflect.Kind (%s)", k))
}

var kindDType = map[reflect.Kind]DType{
	reflect.Bool:       Bool,
	reflect.Int:        Int,
	reflect.Int8:       Int8,
	reflect.Int16:      Int16,
	reflect.Int32:      Int32,
	reflect.Int64:      Int64,
	reflect.Uint:       Uint,
	reflect.Uint8:      Uint8,
	reflect.Uint16:     Uint16,
	reflect.Uint32:     Uint32,
	reflect.Uint64:     Uint64,
	reflect.Uintptr:    Uintptr,
	reflect.Float32:    Float32,
	reflect.Float64:    Float64,
	reflect.Complex64:  Complex64,
	reflect.Complex128: Complex128,
	reflect.String:     String,
}

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

// Num returns a unique number bound to each data type.
func (dt DType) Num() int {
	return int(dt & 0xFF)
}

// String satisfies fmt.Stringer interface. It behaves like Name method.
func (dt DType) String() string {
	return dt.Name()
}

var dTypeNames = map[DType]string{
	Bool:       "bool",
	Int:        "int",
	Int8:       "int8",
	Int16:      "int16",
	Int32:      "int32",
	Int64:      "int64",
	Uint:       "uint",
	Uint8:      "uint8",
	Uint16:     "uint16",
	Uint32:     "uint32",
	Uint64:     "uint64",
	Uintptr:    "uintptr",
	Float32:    "float32",
	Float64:    "float64",
	Complex64:  "complex64",
	Complex128: "complex128",
	String:     "string",
}
