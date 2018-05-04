package buffer_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/ppknap/tensor/internal/buffer"
	"github.com/ppknap/tensor/internal/core"
)

func BenchmarkRawInt_1(b *testing.B)       { bckRawInt(b, 1) }
func BenchmarkBufferInt_1(b *testing.B)    { bckBufferInt(b, 1) }
func BenchmarkReflectInt_1(b *testing.B)   { bckReflectInt(b, 1) }
func BenchmarkInterfaceInt_1(b *testing.B) { bckInterfaceInt(b, 1) }

func BenchmarkRawInt_1k(b *testing.B)       { bckRawInt(b, 1000) }
func BenchmarkBufferInt_1k(b *testing.B)    { bckBufferInt(b, 1000) }
func BenchmarkReflectInt_1k(b *testing.B)   { bckReflectInt(b, 1000) }
func BenchmarkInterfaceInt_1k(b *testing.B) { bckInterfaceInt(b, 1000) }

func BenchmarkRawInt_1M(b *testing.B)       { bckRawInt(b, 1000000) }
func BenchmarkBufferInt_1M(b *testing.B)    { bckBufferInt(b, 1000000) }
func BenchmarkReflectInt_1M(b *testing.B)   { bckReflectInt(b, 1000000) }
func BenchmarkInterfaceInt_1M(b *testing.B) { bckInterfaceInt(b, 1000000) }

func bckRawInt(b *testing.B, size int) {
	var buf = make([]int, size)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			buf[i] = (size + i) % size / 2
		}
	}
}

func bckBufferInt(b *testing.B, size int) { bckBuffer(b, buffer.New(size).AsType(core.Int), size) }

func bckBuffer(b *testing.B, buf *buffer.Buffer, v interface{}) {
	dt, src := core.Destruct(v)
	if dt != buf.DType() {
		panic("invalid benchmark argument types")
	}

	var setter func(i int, src, dst unsafe.Pointer)
	switch dt {
	case core.Int:
		setter = func(i int, dst, src unsafe.Pointer) {
			*(*int)(dst) = (*(*int)(src) + i) % *(*int)(src) / 2
		}
	default:
		panic("unsupported type: " + dt.String())
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf.Iterate(func(i int, dst unsafe.Pointer) {
			setter(i, dst, src)
		})
	}
}

func bckReflectInt(b *testing.B, size int)   { bckReflect(b, make([]int, size), size) }
func bckInterfaceInt(b *testing.B, size int) { bckReflect(b, make([]interface{}, size), size) }

func bckReflect(b *testing.B, d interface{}, v interface{}) {
	if kind := reflect.TypeOf(d).Kind(); kind != reflect.Slice {
		panic("invalid benchmark data type")
	}

	var setter = func() { panic("invalid benchmark input") }

	s := reflect.ValueOf(d)
	switch v := v.(type) {
	case int:
		setter = func() {
			for i := 0; i < s.Len(); i++ {
				s.Index(i).Set(reflect.ValueOf((v + i) % v / 2))
			}
		}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		setter()
	}
}
