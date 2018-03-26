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
	case Int64:
		*dst = *(*int64)(p) != 0
	}

	panic("core: unsupported type: " + dt.String())
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

// AsInt64 converts value under provided pointer to int64 type and saves
// the result to dst. Conversion depends on called data type.
func (dt DType) AsInt64(dst *int64, p unsafe.Pointer) {
	switch dt {
	case Bool:
		if *(*bool)(p) {
			*dst = 1
		}
		*dst = 0
	case Int64:
		*dst = *(*int64)(p)
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

func (dt DType) AsStringFunc() func(unsafe.Pointer) string {
	switch dt {
	case Bool:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*bool)(p)) }
	case Int64:
		return func(p unsafe.Pointer) string { return fmt.Sprint(*(*int64)(p)) }
	}

	panic("core: unsupported type: " + dt.String())
}

func convert(dt, st DType, sv unsafe.Pointer) unsafe.Pointer {
	switch dt {
	case Bool:
		return asDTypeBool(st, sv)
	case Int64:
		return asDTypeInt64(st, sv)
	}

	panic("core: unsupported convert destination type: " + dt.String())
}
