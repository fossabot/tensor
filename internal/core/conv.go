package core

import (
	"fmt"
	"unsafe"
)

// AsBool converts value under provided pointer to bool type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsBool(dst *bool, p unsafe.Pointer) {
	switch dt {
	case Bool:
		*dst = *(*bool)(p)
	case Int:
		*dst = *(*int)(p) != 0
	case Int64:
		*dst = *(*int64)(p) != int64(0)
	case Uint:
		*dst = *(*uint)(p) != uint(0)
	case Uint8:
		*dst = *(*uint8)(p) != uint8(0)
	case Uint16:
		*dst = *(*uint16)(p) != uint16(0)
	case Uint64:
		*dst = *(*uint64)(p) != uint64(0)
	case Float32:
		*dst = *(*float32)(p) != float32(0)
	case Float64:
		*dst = *(*float64)(p) != float64(0)
	case Complex64:
		*dst = *(*complex64)(p) != complex64(0)
	case Complex128:
		*dst = *(*complex128)(p) != complex128(0)
	case String:
		*dst = strAsBool(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsBoolPtr converts value under provided pointer to bool type and returns
// a pointer to its data.
func (dt DType) AsBoolPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Bool {
		return p
	}

	var v bool
	dt.AsBool(&v, p)

	return unsafe.Pointer(&v)
}

// AsInt converts value under provided pointer to int type and saves the result
// to dst. Conversion depends on called data type.
func (dt DType) AsInt(dst *int, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = 1
		}
		*dst = 0
	case Int:
		*dst = *(*int)(p)
	case Int64:
		*dst = (int)(*(*int64)(p))
	case Uint:
		*dst = (int)(*(*uint)(p))
	case Uint8:
		*dst = (int)(*(*uint8)(p))
	case Uint16:
		*dst = (int)(*(*uint16)(p))
	case Uint64:
		*dst = (int)(*(*uint64)(p))
	case Float32:
		*dst = (int)(*(*float32)(p))
	case Float64:
		*dst = (int)(*(*float64)(p))
	case Complex64:
		*dst = (int)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (int)(real(*(*complex128)(p)))
	case String:
		*dst = (int)(strAsInt(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsIntPtr converts value under provided pointer to int type and returns
// a pointer to its data.
func (dt DType) AsIntPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Int {
		return p
	}

	var v int
	dt.AsInt(&v, p)

	return unsafe.Pointer(&v)
}

// AsInt64 converts value under provided pointer to int64 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsInt64(dst *int64, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = int64(1)
		}
		*dst = int64(0)
	case Int:
		*dst = (int64)(*(*int)(p))
	case Int64:
		*dst = *(*int64)(p)
	case Uint:
		*dst = (int64)(*(*uint)(p))
	case Uint8:
		*dst = (int64)(*(*uint8)(p))
	case Uint16:
		*dst = (int64)(*(*uint16)(p))
	case Uint64:
		*dst = (int64)(*(*uint64)(p))
	case Float32:
		*dst = (int64)(*(*float32)(p))
	case Float64:
		*dst = (int64)(*(*float64)(p))
	case Complex64:
		*dst = (int64)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (int64)(real(*(*complex128)(p)))
	case String:
		*dst = strAsInt(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsInt64Ptr converts value under provided pointer to int64 type and returns
// a pointer to its data.
func (dt DType) AsInt64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Int64 {
		return p
	}

	var v int64
	dt.AsInt64(&v, p)

	return unsafe.Pointer(&v)
}

// AsUint converts value under provided pointer to uint type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint(dst *uint, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = uint(1)
		}
		*dst = uint(0)
	case Int:
		*dst = (uint)(*(*int)(p))
	case Int64:
		*dst = (uint)(*(*int64)(p))
	case Uint:
		*dst = *(*uint)(p)
	case Uint8:
		*dst = (uint)(*(*uint8)(p))
	case Uint16:
		*dst = (uint)(*(*uint16)(p))
	case Uint64:
		*dst = (uint)(*(*uint64)(p))
	case Float32:
		*dst = (uint)(*(*float32)(p))
	case Float64:
		*dst = (uint)(*(*float64)(p))
	case Complex64:
		*dst = (uint)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (uint)(real(*(*complex128)(p)))
	case String:
		*dst = (uint)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUintPtr converts value under provided pointer to uint type and returns
// a pointer to its data.
func (dt DType) AsUintPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint {
		return p
	}

	var v uint
	dt.AsUint(&v, p)

	return unsafe.Pointer(&v)
}

// AsUint8 converts value under provided pointer to uint8 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint8(dst *uint8, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = uint8(1)
		}
		*dst = uint8(0)
	case Int:
		*dst = (uint8)(*(*int)(p))
	case Int64:
		*dst = (uint8)(*(*int64)(p))
	case Uint:
		*dst = (uint8)(*(*uint)(p))
	case Uint8:
		*dst = *(*uint8)(p)
	case Uint16:
		*dst = (uint8)(*(*uint16)(p))
	case Uint64:
		*dst = (uint8)(*(*uint64)(p))
	case Float32:
		*dst = (uint8)(*(*float32)(p))
	case Float64:
		*dst = (uint8)(*(*float64)(p))
	case Complex64:
		*dst = (uint8)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (uint8)(real(*(*complex128)(p)))
	case String:
		*dst = (uint8)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint8Ptr converts value under provided pointer to uint8 type and returns
// a pointer to its data.
func (dt DType) AsUint8Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint8 {
		return p
	}

	var v uint8
	dt.AsUint8(&v, p)

	return unsafe.Pointer(&v)
}

// AsUint16 converts value under provided pointer to uint16 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint16(dst *uint16, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = uint16(1)
		}
		*dst = uint16(0)
	case Int:
		*dst = (uint16)(*(*int)(p))
	case Int64:
		*dst = (uint16)(*(*int64)(p))
	case Uint:
		*dst = (uint16)(*(*uint)(p))
	case Uint8:
		*dst = (uint16)(*(*uint8)(p))
	case Uint16:
		*dst = *(*uint16)(p)
	case Uint64:
		*dst = (uint16)(*(*uint64)(p))
	case Float32:
		*dst = (uint16)(*(*float32)(p))
	case Float64:
		*dst = (uint16)(*(*float64)(p))
	case Complex64:
		*dst = (uint16)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (uint16)(real(*(*complex128)(p)))
	case String:
		*dst = (uint16)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint16Ptr converts value under provided pointer to uint16 type and returns
// a pointer to its data.
func (dt DType) AsUint16Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint16 {
		return p
	}

	var v uint16
	dt.AsUint16(&v, p)

	return unsafe.Pointer(&v)
}

// AsUint64 converts value under provided pointer to uint64 type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsUint64(dst *uint64, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = uint64(1)
		}
		*dst = uint64(0)
	case Int:
		*dst = (uint64)(*(*int)(p))
	case Int64:
		*dst = (uint64)(*(*int64)(p))
	case Uint:
		*dst = (uint64)(*(*uint)(p))
	case Uint8:
		*dst = (uint64)(*(*uint8)(p))
	case Uint16:
		*dst = (uint64)(*(*uint16)(p))
	case Uint64:
		*dst = *(*uint64)(p)
	case Float32:
		*dst = (uint64)(*(*float32)(p))
	case Float64:
		*dst = (uint64)(*(*float64)(p))
	case Complex64:
		*dst = (uint64)(real(*(*complex64)(p)))
	case Complex128:
		*dst = (uint64)(real(*(*complex128)(p)))
	case String:
		*dst = (uint64)(strAsUint(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsUint64Ptr converts value under provided pointer to uint64 type and returns
// a pointer to its data.
func (dt DType) AsUint64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Uint64 {
		return p
	}

	var v uint64
	dt.AsUint64(&v, p)

	return unsafe.Pointer(&v)
}

// AsFloat32 converts value under provided pointer to float32 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsFloat32(dst *float32, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = float32(1)
		}
		*dst = float32(0)
	case Int:
		*dst = (float32)(*(*int)(p))
	case Int64:
		*dst = (float32)(*(*int64)(p))
	case Uint:
		*dst = (float32)(*(*uint)(p))
	case Uint8:
		*dst = (float32)(*(*uint8)(p))
	case Uint16:
		*dst = (float32)(*(*uint16)(p))
	case Uint64:
		*dst = (float32)(*(*uint64)(p))
	case Float32:
		*dst = *(*float32)(p)
	case Float64:
		*dst = (float32)(*(*float64)(p))
	case Complex64:
		*dst = real(*(*complex64)(p))
	case Complex128:
		*dst = (float32)(real(*(*complex128)(p)))
	case String:
		*dst = (float32)(strAsFloat(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsFloat32Ptr converts value under provided pointer to float32 type and
// returns a pointer to its data.
func (dt DType) AsFloat32Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Float32 {
		return p
	}

	var v float32
	dt.AsFloat32(&v, p)

	return unsafe.Pointer(&v)
}

// AsFloat64 converts value under provided pointer to float64 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsFloat64(dst *float64, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = float64(1)
		}
		*dst = float64(0)
	case Int:
		*dst = (float64)(*(*int)(p))
	case Int64:
		*dst = (float64)(*(*int64)(p))
	case Uint:
		*dst = (float64)(*(*uint)(p))
	case Uint8:
		*dst = (float64)(*(*uint8)(p))
	case Uint16:
		*dst = (float64)(*(*uint16)(p))
	case Uint64:
		*dst = (float64)(*(*uint64)(p))
	case Float32:
		*dst = (float64)(*(*float32)(p))
	case Float64:
		*dst = *(*float64)(p)
	case Complex64:
		*dst = (float64)(real(*(*complex64)(p)))
	case Complex128:
		*dst = real(*(*complex128)(p))
	case String:
		*dst = strAsFloat(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsFloat64Ptr converts value under provided pointer to float64 type and returns
// a pointer to its data.
func (dt DType) AsFloat64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Float64 {
		return p
	}

	var v float64
	dt.AsFloat64(&v, p)

	return unsafe.Pointer(&v)
}

// AsComplex64 converts value under provided pointer to complex64 type and
// saves the result to dst. Conversion depends on called data type.
func (dt DType) AsComplex64(dst *complex64, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = complex(1, 0)
		}
		*dst = complex(0, 0)
	case Int:
		*dst = complex((float32)(*(*int)(p)), 0)
	case Int64:
		*dst = complex((float32)(*(*int64)(p)), 0)
	case Uint:
		*dst = complex((float32)(*(*uint)(p)), 0)
	case Uint8:
		*dst = complex((float32)(*(*uint8)(p)), 0)
	case Uint16:
		*dst = complex((float32)(*(*uint16)(p)), 0)
	case Uint64:
		*dst = complex((float32)(*(*uint64)(p)), 0)
	case Float32:
		*dst = complex(*(*float32)(p), 0)
	case Float64:
		*dst = complex((float32)(*(*float64)(p)), 0)
	case Complex64:
		*dst = *(*complex64)(p)
	case Complex128:
		*dst = (complex64)(*(*complex128)(p))
	case String:
		*dst = (complex64)(strAsComplex(*(*string)(p)))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsComplex64Ptr converts value under provided pointer to complex64 type and
// returns a pointer to its data.
func (dt DType) AsComplex64Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Complex64 {
		return p
	}

	var v complex64
	dt.AsComplex64(&v, p)

	return unsafe.Pointer(&v)
}

// AsComplex128 converts value under provided pointer to complex128 type and
// saves the result to dst. Conversion depends on called data type.
func (dt DType) AsComplex128(dst *complex128, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = complex(1, 0)
		}
		*dst = complex(0, 0)
	case Int:
		*dst = complex((float64)(*(*int)(p)), 0)
	case Int64:
		*dst = complex((float64)(*(*int64)(p)), 0)
	case Uint:
		*dst = complex((float64)(*(*uint)(p)), 0)
	case Uint8:
		*dst = complex((float64)(*(*uint8)(p)), 0)
	case Uint16:
		*dst = complex((float64)(*(*uint16)(p)), 0)
	case Uint64:
		*dst = complex((float64)(*(*uint64)(p)), 0)
	case Float32:
		*dst = complex((float64)(*(*float32)(p)), 0)
	case Float64:
		*dst = complex(*(*float64)(p), 0)
	case Complex64:
		*dst = (complex128)(*(*complex64)(p))
	case Complex128:
		*dst = *(*complex128)(p)
	case String:
		*dst = strAsComplex(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsComplex128Ptr converts value under provided pointer to complex128 type and
// returns a pointer to its data.
func (dt DType) AsComplex128Ptr(p unsafe.Pointer) unsafe.Pointer {
	if dt == Complex128 {
		return p
	}

	var v complex128
	dt.AsComplex128(&v, p)

	return unsafe.Pointer(&v)
}

// AsString converts value under provided pointer to string type and saves the
// result to dst. Conversion depends on called data type.
func (dt DType) AsString(dst *string, p unsafe.Pointer) {
	switch dt {
	case Bool:
		*dst = fmt.Sprint(*(*bool)(p))
	case Int:
		*dst = fmt.Sprint(*(*int)(p))
	case Int64:
		*dst = fmt.Sprint(*(*int64)(p))
	case Uint:
		*dst = fmt.Sprint(*(*uint)(p))
	case Uint8:
		*dst = fmt.Sprint(*(*uint8)(p))
	case Uint16:
		*dst = fmt.Sprint(*(*uint16)(p))
	case Uint64:
		*dst = fmt.Sprint(*(*uint64)(p))
	case Float32:
		*dst = fmt.Sprint(*(*float32)(p))
	case Float64:
		*dst = fmt.Sprint(*(*float64)(p))
	case Complex64:
		*dst = fmt.Sprint(*(*complex64)(p))
	case Complex128:
		*dst = fmt.Sprint(*(*complex128)(p))
	case String:
		*dst = *(*string)(p)
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsStringPtr converts value under provided pointer to string type and returns
// a pointer to its data.
func (dt DType) AsStringPtr(p unsafe.Pointer) unsafe.Pointer {
	if dt == String {
		return p
	}

	var v string
	dt.AsString(&v, p)

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
	case Int64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int64)(p)) }
	case Uint:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint)(p)) }
	case Uint8:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint8)(p)) }
	case Uint16:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*uint16)(p)) }
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
	case Int64:
		return st.AsInt64Ptr(p)
	case Uint:
		return st.AsUintPtr(p)
	case Uint8:
		return st.AsUint8Ptr(p)
	case Uint16:
		return st.AsUint16Ptr(p)
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
