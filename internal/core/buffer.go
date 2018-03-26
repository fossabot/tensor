package core

import (
	"strings"
	"unsafe"
)

// DefaultBufferDType is a default type a Buffer object will get if its type is
// not set explicitly or it is not already inherited from the first inserted value.
const DefaultBufferDType = Int64

type Buffer struct {
	n    int
	data []byte
	typ  DType
}

// NewBuffer creates a new Buffer instance with provided number of elements.
func NewBuffer(n int) *Buffer {
	return &Buffer{
		n: n,
	}
}

// Len returns the number of items the buffer can store.
func (b *Buffer) Len() int {
	return b.n
}

// Size returns the number of bytes used to store buffer's data.
func (b *Buffer) Size() int {
	return b.Len() * int(b.typ.Size())
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
	if b.Len() == 0 {
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
		typ.Setraw(
			unsafe.Pointer(uintptr(unsafe.Pointer(&dst[0]))+pos),
			typ.Convert(b.typ, unsafe.Pointer(uintptr(unsafe.Pointer(&b.data[0]))+oldpos)),
		)

		pos += typ.Size()
	}

	return pos
}

func (b *Buffer) Setval(i int, v interface{}) {

}

func (b *Buffer) Setraw(i int, typ DType, p unsafe.Pointer) {

}

func (b *Buffer) At(i int) unsafe.Pointer {
	return nil
}

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

// func Tst() {
// 	a := NewBuffer(1).AsType(Int64)

// 	conv := a.DType().AsStringFunc()

// 	b := NewBuffer(4).AsType(Int64)
// 	b.Iterate(func(i int, lhv unsafe.Pointer) {
// 		b.Converted(a.DType(), a.At(0), func(rhv unsafe.Pointer) {
// 			if b.DType() == Int64 {
// 				*(*int64)(lhv) += *(int64)(rhv)
// 			}
// 		})
// 	})
// }

// String satisfies fmt.Stringer interface. It produces the same results as the
// specific type slices would produce with default formatting.
func (b *Buffer) String() string {
	vs, scf := []string(nil), b.typ.AsStringFunc()

	b.Iterate(func(_ int, p unsafe.Pointer) {
		vs = append(vs, scf(p))
	})

	return "[" + strings.Join(vs, " ") + "]"
}
