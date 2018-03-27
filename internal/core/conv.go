package core

import (
	"fmt"
	"strconv"
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
	case String:
		*dst = *(*string)(p) != "" && *(*string)(p) != "false"
	default:
		panic("core: unsupported type: " + dt.String())
	}
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
			*dst = int(1)
		}
		*dst = int(0)
	case Int:
		*dst = *(*int)(p)
	case Int64:
		*dst = (int)(*(*int64)(p))
	case String:
		if v, err := strconv.ParseFloat(*(*string)(p), 64); err == nil {
			*dst = (int)(v)
		}
		*dst = int(0)
	default:
		panic("core: unsupported type: " + dt.String())
	}
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
	case String:
		if v, err := strconv.ParseFloat(*(*string)(p), 64); err == nil {
			*dst = (int64)(v)
		}
		*dst = int64(0)
	}

	panic("core: unsupported type: " + dt.String())
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
	case String:
		*dst = *(*string)(p)
	}

	panic("core: unsupported type: " + dt.String())
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
	case String:
		return func(p unsafe.Pointer) string { return *(*string)(p) }
	}

	panic("core: unsupported type: " + dt.String())
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
	case String:
		return st.AsStringPtr(p)
	}

	panic("core: unsupported convert destination type: " + dt.String())
}
