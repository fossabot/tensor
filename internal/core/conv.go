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
	case Float64:
		*dst = *(*float64)(p) != float64(0)
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
	case Float64:
		*dst = (int)(*(*float64)(p))
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
	case Float64:
		*dst = (int64)(*(*float64)(p))
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
		*dst = (uint)(*(*uint)(p))
	case Float64:
		*dst = (uint)(*(*float64)(p))
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
	case Float64:
		*dst = *(*float64)(p)
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
	case Float64:
		*dst = complex(*(*float64)(p), 0)
	case Complex128:
		*dst = *(*complex128)(p)
	case String:
		*dst = strAsComplex(*(*string)(p))
	}

	panic(NewError("unsupported type: %q", dt))
}

// AsComplex128Ptr converts value under provided pointer to float64 type and returns
// a pointer to its data.
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
	case Float64:
		*dst = fmt.Sprint(*(*float64)(p))
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
	case Float64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*float64)(p)) }
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
	case Float64:
		return st.AsFloat64Ptr(p)
	case Complex128:
		return st.AsComplex128Ptr(p)
	case String:
		return st.AsStringPtr(p)
	}

	panic(NewError("unsupported convert destination type: %q", dt))
}
