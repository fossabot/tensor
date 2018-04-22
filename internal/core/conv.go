package core

import (
	"fmt"
	"unsafe"
)

// AsBool converts value under provided pointer to bool type. Conversion depends
// on called data type.
func (dt DType) AsBool(p unsafe.Pointer) bool {
	switch dt {
	case Bool:
		return *(*bool)(p)
	case Int:
		return *(*int)(p) != 0
	case Int16:
		return *(*int16)(p) != int16(0)
	case Int64:
		return *(*int64)(p) != int64(0)
	case Uint:
		return *(*uint)(p) != uint(0)
	case Uint8:
		return *(*uint8)(p) != uint8(0)
	case Uint16:
		return *(*uint16)(p) != uint16(0)
	case Uint32:
		return *(*uint32)(p) != uint32(0)
	case Uint64:
		return *(*uint64)(p) != uint64(0)
	case Float32:
		return *(*float32)(p) != float32(0)
	case Float64:
		return *(*float64)(p) != float64(0)
	case Complex64:
		return *(*complex64)(p) != complex64(0)
	case Complex128:
		return *(*complex128)(p) != complex128(0)
	case String:
		return strAsBool(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsBoolPtr converts value under provided pointer to bool type and returns
// a pointer to its data.
func (dt DType) AsBoolPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Bool {
		return p
	}

	v := dt.AsBool(p)
	return unsafe.Pointer(&v)
}

// AsInt converts value under provided pointer to int type and saves the result
// to dst. Conversion depends on called data type.
func (dt DType) AsInt(p unsafe.Pointer) int {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return 1
		}
		return 0
	case Int:
		return *(*int)(p)
	case Int16:
		return (int)(*(*int16)(p))
	case Int64:
		return (int)(*(*int64)(p))
	case Uint:
		return (int)(*(*uint)(p))
	case Uint8:
		return (int)(*(*uint8)(p))
	case Uint16:
		return (int)(*(*uint16)(p))
	case Uint32:
		return (int)(*(*uint32)(p))
	case Uint64:
		return (int)(*(*uint64)(p))
	case Float32:
		return (int)(*(*float32)(p))
	case Float64:
		return (int)(*(*float64)(p))
	case Complex64:
		return (int)(real(*(*complex64)(p)))
	case Complex128:
		return (int)(real(*(*complex128)(p)))
	case String:
		return (int)(strAsInt(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsIntPtr converts value under provided pointer to int type and returns
// a pointer to its data.
func (dt DType) AsIntPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Int {
		return p
	}

	v := dt.AsInt(p)
	return unsafe.Pointer(&v)
}

// AsInt64 converts value under provided pointer to int64 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsInt64(p unsafe.Pointer) int64 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return int64(1)
		}
		return int64(0)
	case Int:
		return (int64)(*(*int)(p))
	case Int16:
		return (int64)(*(*int16)(p))
	case Int64:
		return *(*int64)(p)
	case Uint:
		return (int64)(*(*uint)(p))
	case Uint8:
		return (int64)(*(*uint8)(p))
	case Uint16:
		return (int64)(*(*uint16)(p))
	case Uint32:
		return (int64)(*(*uint32)(p))
	case Uint64:
		return (int64)(*(*uint64)(p))
	case Float32:
		return (int64)(*(*float32)(p))
	case Float64:
		return (int64)(*(*float64)(p))
	case Complex64:
		return (int64)(real(*(*complex64)(p)))
	case Complex128:
		return (int64)(real(*(*complex128)(p)))
	case String:
		return strAsInt(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsInt64Ptr converts value under provided pointer to int64 type and returns
// a pointer to its data.
func (dt DType) AsInt64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Int64 {
		return p
	}

	v := dt.AsInt64(p)
	return unsafe.Pointer(&v)
}

// AsInt16 converts value under provided pointer to int16 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsInt16(p unsafe.Pointer) int16 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return int16(1)
		}
		return int16(0)
	case Int:
		return (int16)(*(*int)(p))
	case Int16:
		return *(*int16)(p)
	case Int64:
		return (int16)(*(*int64)(p))
	case Uint:
		return (int16)(*(*uint)(p))
	case Uint8:
		return (int16)(*(*uint8)(p))
	case Uint16:
		return (int16)(*(*uint16)(p))
	case Uint32:
		return (int16)(*(*uint32)(p))
	case Uint64:
		return (int16)(*(*uint64)(p))
	case Float32:
		return (int16)(*(*float32)(p))
	case Float64:
		return (int16)(*(*float64)(p))
	case Complex64:
		return (int16)(real(*(*complex64)(p)))
	case Complex128:
		return (int16)(real(*(*complex128)(p)))
	case String:
		return (int16)(strAsInt(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsInt16Ptr converts value under provided pointer to int16 type and returns
// a pointer to its data.
func (dt DType) AsInt16Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Int16 {
		return p
	}

	v := dt.AsInt16(p)
	return unsafe.Pointer(&v)
}

// AsUint converts value under provided pointer to uint type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint(p unsafe.Pointer) uint {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return uint(1)
		}
		return uint(0)
	case Int:
		return (uint)(*(*int)(p))
	case Int16:
		return (uint)(*(*int16)(p))
	case Int64:
		return (uint)(*(*int64)(p))
	case Uint:
		return *(*uint)(p)
	case Uint8:
		return (uint)(*(*uint8)(p))
	case Uint16:
		return (uint)(*(*uint16)(p))
	case Uint32:
		return (uint)(*(*uint32)(p))
	case Uint64:
		return (uint)(*(*uint64)(p))
	case Float32:
		return (uint)(*(*float32)(p))
	case Float64:
		return (uint)(*(*float64)(p))
	case Complex64:
		return (uint)(real(*(*complex64)(p)))
	case Complex128:
		return (uint)(real(*(*complex128)(p)))
	case String:
		return (uint)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUintPtr converts value under provided pointer to uint type and returns
// a pointer to its data.
func (dt DType) AsUintPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint {
		return p
	}

	v := dt.AsUint(p)
	return unsafe.Pointer(&v)
}

// AsUint8 converts value under provided pointer to uint8 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint8(p unsafe.Pointer) uint8 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return uint8(1)
		}
		return uint8(0)
	case Int:
		return (uint8)(*(*int)(p))
	case Int16:
		return (uint8)(*(*int16)(p))
	case Int64:
		return (uint8)(*(*int64)(p))
	case Uint:
		return (uint8)(*(*uint)(p))
	case Uint8:
		return *(*uint8)(p)
	case Uint16:
		return (uint8)(*(*uint16)(p))
	case Uint32:
		return (uint8)(*(*uint32)(p))
	case Uint64:
		return (uint8)(*(*uint64)(p))
	case Float32:
		return (uint8)(*(*float32)(p))
	case Float64:
		return (uint8)(*(*float64)(p))
	case Complex64:
		return (uint8)(real(*(*complex64)(p)))
	case Complex128:
		return (uint8)(real(*(*complex128)(p)))
	case String:
		return (uint8)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint8Ptr converts value under provided pointer to uint8 type and returns
// a pointer to its data.
func (dt DType) AsUint8Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint8 {
		return p
	}

	v := dt.AsUint8(p)
	return unsafe.Pointer(&v)
}

// AsUint16 converts value under provided pointer to uint16 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint16(p unsafe.Pointer) uint16 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return uint16(1)
		}
		return uint16(0)
	case Int:
		return (uint16)(*(*int)(p))
	case Int16:
		return (uint16)(*(*int16)(p))
	case Int64:
		return (uint16)(*(*int64)(p))
	case Uint:
		return (uint16)(*(*uint)(p))
	case Uint8:
		return (uint16)(*(*uint8)(p))
	case Uint16:
		return *(*uint16)(p)
	case Uint32:
		return (uint16)(*(*uint32)(p))
	case Uint64:
		return (uint16)(*(*uint64)(p))
	case Float32:
		return (uint16)(*(*float32)(p))
	case Float64:
		return (uint16)(*(*float64)(p))
	case Complex64:
		return (uint16)(real(*(*complex64)(p)))
	case Complex128:
		return (uint16)(real(*(*complex128)(p)))
	case String:
		return (uint16)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint16Ptr converts value under provided pointer to uint16 type and returns
// a pointer to its data.
func (dt DType) AsUint16Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint16 {
		return p
	}

	v := dt.AsUint16(p)
	return unsafe.Pointer(&v)
}

// AsUint32 converts value under provided pointer to uint32 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint32(p unsafe.Pointer) uint32 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return uint32(1)
		}
		return uint32(0)
	case Int:
		return (uint32)(*(*int)(p))
	case Int16:
		return (uint32)(*(*int16)(p))
	case Int64:
		return (uint32)(*(*int64)(p))
	case Uint:
		return (uint32)(*(*uint)(p))
	case Uint8:
		return (uint32)(*(*uint8)(p))
	case Uint16:
		return (uint32)(*(*uint16)(p))
	case Uint32:
		return *(*uint32)(p)
	case Uint64:
		return (uint32)(*(*uint64)(p))
	case Float32:
		return (uint32)(*(*float32)(p))
	case Float64:
		return (uint32)(*(*float64)(p))
	case Complex64:
		return (uint32)(real(*(*complex64)(p)))
	case Complex128:
		return (uint32)(real(*(*complex128)(p)))
	case String:
		return (uint32)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint32Ptr converts value under provided pointer to uint32 type and returns
// a pointer to its data.
func (dt DType) AsUint32Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint32 {
		return p
	}

	v := dt.AsUint32(p)
	return unsafe.Pointer(&v)
}

// AsUint64 converts value under provided pointer to uint64 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint64(p unsafe.Pointer) uint64 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return uint64(1)
		}
		return uint64(0)
	case Int:
		return (uint64)(*(*int)(p))
	case Int16:
		return (uint64)(*(*int16)(p))
	case Int64:
		return (uint64)(*(*int64)(p))
	case Uint:
		return (uint64)(*(*uint)(p))
	case Uint8:
		return (uint64)(*(*uint8)(p))
	case Uint16:
		return (uint64)(*(*uint16)(p))
	case Uint32:
		return (uint64)(*(*uint32)(p))
	case Uint64:
		return *(*uint64)(p)
	case Float32:
		return (uint64)(*(*float32)(p))
	case Float64:
		return (uint64)(*(*float64)(p))
	case Complex64:
		return (uint64)(real(*(*complex64)(p)))
	case Complex128:
		return (uint64)(real(*(*complex128)(p)))
	case String:
		return (uint64)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint64Ptr converts value under provided pointer to uint64 type and returns
// a pointer to its data.
func (dt DType) AsUint64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint64 {
		return p
	}

	v := dt.AsUint64(p)
	return unsafe.Pointer(&v)
}

// AsFloat32 converts value under provided pointer to float32 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsFloat32(p unsafe.Pointer) float32 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return float32(1)
		}
		return float32(0)
	case Int:
		return (float32)(*(*int)(p))
	case Int16:
		return (float32)(*(*int16)(p))
	case Int64:
		return (float32)(*(*int64)(p))
	case Uint:
		return (float32)(*(*uint)(p))
	case Uint8:
		return (float32)(*(*uint8)(p))
	case Uint16:
		return (float32)(*(*uint16)(p))
	case Uint32:
		return (float32)(*(*uint32)(p))
	case Uint64:
		return (float32)(*(*uint64)(p))
	case Float32:
		return *(*float32)(p)
	case Float64:
		return (float32)(*(*float64)(p))
	case Complex64:
		return real(*(*complex64)(p))
	case Complex128:
		return (float32)(real(*(*complex128)(p)))
	case String:
		return (float32)(strAsFloat(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsFloat32Ptr converts value under provided pointer to float32 type and
// returns a pointer to its data.
func (dt DType) AsFloat32Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Float32 {
		return p
	}

	v := dt.AsFloat32(p)
	return unsafe.Pointer(&v)
}

// AsFloat64 converts value under provided pointer to float64 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsFloat64(p unsafe.Pointer) float64 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return float64(1)
		}
		return float64(0)
	case Int:
		return (float64)(*(*int)(p))
	case Int16:
		return (float64)(*(*int16)(p))
	case Int64:
		return (float64)(*(*int64)(p))
	case Uint:
		return (float64)(*(*uint)(p))
	case Uint8:
		return (float64)(*(*uint8)(p))
	case Uint16:
		return (float64)(*(*uint16)(p))
	case Uint32:
		return (float64)(*(*uint32)(p))
	case Uint64:
		return (float64)(*(*uint64)(p))
	case Float32:
		return (float64)(*(*float32)(p))
	case Float64:
		return *(*float64)(p)
	case Complex64:
		return (float64)(real(*(*complex64)(p)))
	case Complex128:
		return real(*(*complex128)(p))
	case String:
		return strAsFloat(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsFloat64Ptr converts value under provided pointer to float64 type and returns
// a pointer to its data.
func (dt DType) AsFloat64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Float64 {
		return p
	}

	v := dt.AsFloat64(p)
	return unsafe.Pointer(&v)
}

// AsComplex64 converts value under provided pointer to complex64 type and
// saves the result to dst. Conversion depends on called data type.
func (dt DType) AsComplex64(p unsafe.Pointer) complex64 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return complex(1, 0)
		}
		return complex(0, 0)
	case Int:
		return complex((float32)(*(*int)(p)), 0)
	case Int16:
		return complex((float32)(*(*int16)(p)), 0)
	case Int64:
		return complex((float32)(*(*int64)(p)), 0)
	case Uint:
		return complex((float32)(*(*uint)(p)), 0)
	case Uint8:
		return complex((float32)(*(*uint8)(p)), 0)
	case Uint16:
		return complex((float32)(*(*uint16)(p)), 0)
	case Uint32:
		return complex((float32)(*(*uint32)(p)), 0)
	case Uint64:
		return complex((float32)(*(*uint64)(p)), 0)
	case Float32:
		return complex(*(*float32)(p), 0)
	case Float64:
		return complex((float32)(*(*float64)(p)), 0)
	case Complex64:
		return *(*complex64)(p)
	case Complex128:
		return (complex64)(*(*complex128)(p))
	case String:
		return (complex64)(strAsComplex(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsComplex64Ptr converts value under provided pointer to complex64 type and
// returns a pointer to its data.
func (dt DType) AsComplex64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Complex64 {
		return p
	}

	v := dt.AsComplex64(p)
	return unsafe.Pointer(&v)
}

// AsComplex128 converts value under provided pointer to complex128 type and
// saves the result to dst. Conversion depends on called data type.
func (dt DType) AsComplex128(p unsafe.Pointer) complex128 {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			return complex(1, 0)
		}
		return complex(0, 0)
	case Int:
		return complex((float64)(*(*int)(p)), 0)
	case Int16:
		return complex((float64)(*(*int16)(p)), 0)
	case Int64:
		return complex((float64)(*(*int64)(p)), 0)
	case Uint:
		return complex((float64)(*(*uint)(p)), 0)
	case Uint8:
		return complex((float64)(*(*uint8)(p)), 0)
	case Uint16:
		return complex((float64)(*(*uint16)(p)), 0)
	case Uint32:
		return complex((float64)(*(*uint32)(p)), 0)
	case Uint64:
		return complex((float64)(*(*uint64)(p)), 0)
	case Float32:
		return complex((float64)(*(*float32)(p)), 0)
	case Float64:
		return complex(*(*float64)(p), 0)
	case Complex64:
		return (complex128)(*(*complex64)(p))
	case Complex128:
		return *(*complex128)(p)
	case String:
		return strAsComplex(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsComplex128Ptr converts value under provided pointer to complex128 type and
// returns a pointer to its data.
func (dt DType) AsComplex128Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Complex128 {
		return p
	}

	v := dt.AsComplex128(p)
	return unsafe.Pointer(&v)
}

// AsString converts value under provided pointer to string type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsString(p unsafe.Pointer) string {
	switch dt {
	case Bool:
		return fmt.Sprint(*(*bool)(p))
	case Int:
		return fmt.Sprint(*(*int)(p))
	case Int16:
		return fmt.Sprint(*(*int16)(p))
	case Int64:
		return fmt.Sprint(*(*int64)(p))
	case Uint:
		return fmt.Sprint(*(*uint)(p))
	case Uint8:
		return fmt.Sprint(*(*uint8)(p))
	case Uint16:
		return fmt.Sprint(*(*uint16)(p))
	case Uint32:
		return fmt.Sprint(*(*uint32)(p))
	case Uint64:
		return fmt.Sprint(*(*uint64)(p))
	case Float32:
		return fmt.Sprint(*(*float32)(p))
	case Float64:
		return fmt.Sprint(*(*float64)(p))
	case Complex64:
		return fmt.Sprint(*(*complex64)(p))
	case Complex128:
		return fmt.Sprint(*(*complex128)(p))
	case String:
		return *(*string)(p)
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsStringPtr converts value under provided pointer to string type and returns
// a pointer to its data.
func (dt DType) AsStringPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == String {
		return p
	}

	v := dt.AsString(p)
	return unsafe.Pointer(&v)
}

// AsStringFunc produces a converting function that returns string
// representation of a given value.
func (dt DType) AsStringFunc() func(unsafe.Pointer) string {
	switch dt {
	case Bool:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*bool)(p)) }
	case Int:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int)(p)) }
	case Int16:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int16)(p)) }
	case Int64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int64)(p)) }
	case Uint:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint)(p)) }
	case Uint8:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint8)(p)) }
	case Uint16:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint16)(p)) }
	case Uint32:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint32)(p)) }
	case Uint64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint64)(p)) }
	case Float32:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*float32)(p)) }
	case Float64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*float64)(p)) }
	case Complex64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*complex64)(p)) }
	case Complex128:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*complex128)(p)) }
	case String:
		return func(p unsafe.Pointer) string { return *(*string)(p) }
	}

	panic(NewError("unsupported type: %q", dt))
}

// Convert takes provided pointer and its data type and converts pointer's value
// to data representation given by called object. There are not any write
// operations on passed values.
func (dt DType) Convert(st DType, p unsafe.Pointer) unsafe.Pointer {
	switch dt {
	case Bool:
		return st.AsBoolPtr(p)
	case Int:
		return st.AsIntPtr(p)
	case Int16:
		return st.AsInt16Ptr(p)
	case Int64:
		return st.AsInt64Ptr(p)
	case Uint:
		return st.AsUintPtr(p)
	case Uint8:
		return st.AsUint8Ptr(p)
	case Uint16:
		return st.AsUint16Ptr(p)
	case Uint32:
		return st.AsUint32Ptr(p)
	case Uint64:
		return st.AsUint64Ptr(p)
	case Float32:
		return st.AsFloat32Ptr(p)
	case Float64:
		return st.AsFloat64Ptr(p)
	case Complex64:
		return st.AsComplex64Ptr(p)
	case Complex128:
		return st.AsComplex128Ptr(p)
	case String:
		return st.AsStringPtr(p)
	}

	panic(NewError("unsupported convert destination type: %q", dt))
}
