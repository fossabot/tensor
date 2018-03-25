package core

import (
	"unsafe"
)

const DefaultBufferDType DType = Int64

type Buffer struct {
	n    int
	data []byte
	typ  DType
}

func NewBuffer(n int) *Buffer {
	return &Buffer{
		n: n,
	}
}

func (b *Buffer) Len() int {
	return b.n
}

func (b *Buffer) Size() int {
	return b.Len() * int(b.typ.Size())
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
	if b.Len() == 0 {
		b.typ = typ
		return b
	}

	// Unallocated or empty buffer or type with no size.
	if len(b.data) == 0 || typ.Size() == 0 {
		b.data = make([]byte, b.Size())
		b.typ = typ
		return b
	}

	if b.typ.Size() >= typ.Size() {
		// Reuse existing buffer when replacing type size is smaller or equal
		// than the size of type which is being replaced.
		b.data = b.data[:b.transfer(typ, b.data)]
	} else {
		// Allocate a new buffer since the old one will not be able to store
		// existing data after conversion.
		data := make([]byte, uintptr(b.Len())*typ.Size())
		b.transfer(typ, data)
		b.data = data
	}

	b.typ = typ

	return b
}

// transfer copies data from buffer to provided destination. It makes all
// necessary conversions between object type and provided one.
func (b *Buffer) transfer(typ DType, dst []byte) (pos uintptr) {
	var size = uintptr(b.Size())
	for oldpos := uintptr(0); oldpos < size; oldpos += b.typ.Size() {
		setraw(typ,
			unsafe.Pointer(uintptr(unsafe.Pointer(&dst[0]))+pos),
			convert(typ, b.typ, unsafe.Pointer(uintptr(unsafe.Pointer(&b.data[0]))+oldpos)),
		)

		pos += typ.Size()
	}

	return pos
}

func (b *Buffer) DSet(i int, v interface{}) {

}

func (b *Buffer) DAt(i int, f func(p unsafe.Pointer)) {

}

func (b *Buffer) CSet(i int, v interface{}) error {

}

func (b *Buffer) CAt(i int, f func(p unsafe.Pointer)) error {
	return nil
}

func (b *Buffer) Iterate(f func(i int, p unsafe.Pointer)) {}
