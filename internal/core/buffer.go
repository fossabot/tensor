package core

import (
	"strings"
	"unsafe"
)

type Buffer struct {
	n    int
	data []byte
	pts  []unsafe.Pointer
	typ  DType
}

// NewBuffer creates a new Buffer instance with provided number of elements.
func NewBuffer(n int) *Buffer {
	return &Buffer{
		n: n,
	}
}

// Size returns the number of items the buffer can store.
func (b *Buffer) Size() int {
	return b.n
}

// NBytes returns the number of bytes used to store buffer's data. For dynamic
// types only the object pointer size is counted.
func (b *Buffer) NBytes() int {
	return b.Size() * int(b.typ.Size())
}

// DType returns the underlying buffer's data type.
func (b *Buffer) DType() DType {
	return b.typ
}

// AsType transforms buffer's underlying type to provided one. This function
// reallocates the internal data buffer when the size of provided type is
// greater than the currently stored by called object.
func (b *Buffer) AsType(typ DType) *Buffer {
	// Type is already set so this function is no-op.
	if b.typ == typ {
		return b
	}

	// No need to work with buffer when it's meant to be empty.
	if b.Size() == 0 {
		b.typ = typ
		return b
	}

	// Unallocated or empty buffer or type with no size.
	if len(b.data) == 0 || typ.Size() == 0 {
		b.typ = typ
		b.data = make([]byte, b.Size())
		return b
	}

	if b.typ.Size() >= typ.Size() {
		// Reuse existing buffer when replacing type size is smaller or equal
		// than the size of type which is being replaced.
		b.data = b.data[:b.transfer(typ, b.data)]
	} else {
		// Allocate a new buffer since the old one will not be able to store
		// existing data after conversion.
		data := make([]byte, uintptr(b.Size())*typ.Size())
		b.transfer(typ, data)
		b.data = data
	}

	b.typ = typ

	// ---
	switch {
	case typ.IsDynamic() && b.typ.IsDynamic():
		// Reuse existing pointer buffer.
	}

	// --

	return b
}

// transfer copies data from buffer to provided destination. It makes all
// necessary conversions between object type and provided one.
func (b *Buffer) transfer(typ DType, dst []byte) (pos uintptr) {
	var size = uintptr(b.Size())
	for oldpos := uintptr(0); oldpos < size; oldpos += b.typ.Size() {
		typ.Setraw(
			unsafe.Pointer(uintptr(unsafe.Pointer(&dst[0]))+pos),
			typ.Convert(b.typ, unsafe.Pointer(uintptr(unsafe.Pointer(&b.data[0]))+oldpos)),
		)

		pos += typ.Size()
	}

	return pos
}

// Setval sets interface value to a given position in the buffer. Conversion
// between types may occur when v and buffer types differ.
func (b *Buffer) Setval(i int, v interface{}) {
	typ, p := Destruct(v)
	b.Setptr(i, typ, p)
}

// Setraw directly sets value under p to a given element in the buffer. The
// type of value p points to must be identical as the array type. Method
// behavior is undefined otherwise.
func (b *Buffer) Setraw(i int, p unsafe.Pointer) {
	b.typ.Setraw(b.At(i), p)
}

// Setptr sets value under p to a given position in a buffer. Conversion might
// happen when types differ.
func (b *Buffer) Setptr(i int, typ DType, p unsafe.Pointer) {
	b.typ.Setraw(b.At(i), b.typ.Convert(typ, p))
}

// At gets element at a given position.
func (b *Buffer) At(i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(&b.data[0])) + uintptr(i)*b.typ.Size())
}

// Iterate calls f on each element stored in the buffer.
func (b *Buffer) Iterate(f func(i int, p unsafe.Pointer)) {
	var (
		size = uintptr(b.Size())
		p    = unsafe.Pointer(&b.data[0])
	)

	for pos, i := uintptr(0), 0; pos < size; pos += b.typ.Size() {
		f(i, unsafe.Pointer(uintptr(p)+pos))
		i++
	}
}

// String satisfies fmt.Stringer interface. It produces the same results as the
// specific type slices would produce with default formatting.
func (b *Buffer) String() string {
	vs, scf := []string(nil), b.typ.AsStringFunc()

	b.Iterate(func(_ int, p unsafe.Pointer) {
		vs = append(vs, scf(p))
	})

	return "[" + strings.Join(vs, " ") + "]"
}
