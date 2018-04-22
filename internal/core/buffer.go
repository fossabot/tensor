package core

import (
	"strings"
	"unsafe"
)

// DefaultBufferDType is a default type used by buffer type when its data type
// was not set explicitly.
const DefaultBufferDType = Float64

// Buffer stores a set of data elements. Its size is predefined thus, it is like
// constant size array that can store objects of different type. The ones that
// do not include pointers are stored in contiguous memory segment. This
// property, and the fact that only one type is allowed at a time, makes Buffer
// different from a slice of empty interfaces. Dynamic types are kept indirectly
// by storing their unsafe pointers in a dedicated slice.
//
// If buffer type is not  set explicitly, it defaults to DefaultBufferDType.
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
func (b *Buffer) Size() int { return b.n }

// NBytes returns the number of bytes used to store buffer's data. For dynamic
// types only the object pointer size is counted.
func (b *Buffer) NBytes() int {
	if b.typ == 0 {
		return b.Size() * int(DefaultBufferDType.Size())
	}

	return b.Size() * int(b.typ.Size())
}

// DType returns the underlying buffer's data type.
func (b *Buffer) DType() DType {
	if b.typ != 0 {
		return b.typ
	}

	return DefaultBufferDType
}

// Setval sets interface value to a given position in the buffer. Conversion
// between types may occur when v and buffer types differ.
func (b *Buffer) Setval() func(int, interface{}) {
	setptr := b.Setptr()

	return func(i int, v interface{}) {
		typ, p := Destruct(v)
		setptr(i, typ, p)
	}
}

// Setptr sets value under p to a given position in a buffer. Conversion might
// happen when types differ.
func (b *Buffer) Setptr() func(int, DType, unsafe.Pointer) {
	b.init()

	if !b.typ.IsDynamic() {
		atFunc := b.At()
		return func(i int, typ DType, p unsafe.Pointer) {
			b.typ.Setraw(atFunc(i), b.typ.Convert(typ, p))
		}
	}

	return func(i int, typ DType, p unsafe.Pointer) {
		b.pts[i] = b.typ.Convert(typ, p)
	}
}

// At gets element at a given position. If buffer size is one, any index
// provided to created function will result in a first buffer's element
// returned. This logic simplifies access operation to scalars.
func (b *Buffer) At() func(int) unsafe.Pointer {
	b.init()

	if b.Size() == 1 {
		if b.typ.IsDynamic() {
			return func(int) unsafe.Pointer { return b.pts[0] }
		}

		first := unsafe.Pointer(&b.data[0])
		return func(int) unsafe.Pointer { return first }
	}

	if !b.typ.IsDynamic() {
		first := unsafe.Pointer(&b.data[0])
		return func(i int) unsafe.Pointer {
			return unsafe.Pointer(uintptr(first) + uintptr(i)*b.typ.Size())
		}
	}

	return func(i int) unsafe.Pointer { return b.pts[i] }
}

// Iterate calls f on each element stored in the buffer.
func (b *Buffer) Iterate(f func(i int, p unsafe.Pointer)) {
	b.init()

	size := uintptr(b.Size())

	p := unsafe.Pointer(&b.data[0])
	if b.typ.IsDynamic() {
		p = unsafe.Pointer(&b.pts[0])
	}

	for pos, i := uintptr(0), 0; pos < size; pos += b.typ.Size() {
		f(i, unsafe.Pointer(uintptr(p)+pos))
		i++
	}
}

// String satisfies fmt.Stringer interface. It produces the same results as the
// specific type slices would produce with default formatting.
func (b *Buffer) String() string {
	b.init()

	vs, scf := []string(nil), b.typ.AsStringFunc()

	b.Iterate(func(_ int, p unsafe.Pointer) {
		vs = append(vs, scf(p))
	})

	return "[" + strings.Join(vs, " ") + "]"
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

	var size = uintptr(b.Size())

	// Unallocated buffers or type with no size.
	if (b.data == nil && b.pts == nil) || typ.Size() == 0 {
		b.typ = typ

		if !typ.IsDynamic() {
			b.data = make([]byte, size*typ.Size())
			return b
		}

		// Dynamic type storage requires initialization.
		b.pts = make([]unsafe.Pointer, size)
		for i := range b.pts {
			b.pts[i] = typ.Zero()
		}

		return b
	}

	switch {
	case b.typ.IsDynamic() && typ.IsDynamic(): // Dynamic to Dynamic conversion.
		// Reuse existing pointer buffer.
		b.Iterate(func(i int, p unsafe.Pointer) {
			b.pts[i] = typ.Convert(b.typ, p)
		})

	case b.typ.IsDynamic() && !typ.IsDynamic(): // Dynamic to Static conversion.
		// Convert data and save it in a new buffer.
		b.data = make([]byte, size*typ.Size())

		var first = unsafe.Pointer(&b.data[0])
		b.Iterate(func(i int, p unsafe.Pointer) {
			typ.Setraw(
				unsafe.Pointer(uintptr(first)+uintptr(i)*typ.Size()),
				typ.Convert(b.typ, *(*unsafe.Pointer)(p)),
			)
		})
		b.pts = nil

	case !b.typ.IsDynamic() && typ.IsDynamic(): // Static to Dynamic conversion.
		// Convert to dynamic pointer and save it to pointers buffer. Its data
		// will be tracked by GC thus, not silently deallocated.
		b.pts = make([]unsafe.Pointer, b.Size())

		b.Iterate(func(i int, p unsafe.Pointer) {
			b.pts[i] = typ.Convert(b.typ, p)
		})
		b.data = nil

	case b.typ.Size() >= typ.Size(): // Static to Static conversion.
		// Reuse existing buffer when replacing type size is smaller or equal
		// than the size of type which is being replaced.
		newpos := uintptr(0)
		first := unsafe.Pointer(&b.data[0])
		b.Iterate(func(i int, p unsafe.Pointer) {
			typ.Setraw(
				unsafe.Pointer(uintptr(first)+newpos),
				typ.Convert(b.typ, p),
			)
			newpos += typ.Size()
		})
		b.data = b.data[:newpos]

	default: // Static to Static conversion.
		// Allocate a new buffer since the old one will not be able to store
		// existing data after conversion.
		data := make([]byte, uintptr(b.Size())*typ.Size())
		first := unsafe.Pointer(&data[0])
		b.Iterate(func(i int, p unsafe.Pointer) {
			typ.Setraw(
				unsafe.Pointer(uintptr(first)+uintptr(i)*typ.Size()),
				typ.Convert(b.typ, p),
			)
		})
		b.data = data
	}

	b.typ = typ

	return b
}

func (b *Buffer) init() {
	if b.typ == 0 {
		b.AsType(DefaultBufferDType)
	}
}
