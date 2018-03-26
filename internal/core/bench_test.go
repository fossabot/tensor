package core_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/ppknap/tacvs/internal/core"
)

func BenchmarkFillSliceInt_1(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 1), 700)
}
func BenchmarkFillSliceReflect_1(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 1), int(700))
}
func BenchmarkFillSliceDTypeInt_1(b *testing.B) {
	benchmarkFillSliceDType(b, core.NewBuffer(1).AsType(core.Int64), int64(700))
}

func BenchmarkFillSliceInt_1000(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 100), 700)
}
func BenchmarkFillSliceReflect_1000(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 100), int(700))
}
func BenchmarkFillSliceDTypeInt_1000(b *testing.B) {
	benchmarkFillSliceDType(b, core.NewBuffer(100).AsType(core.Int64), int64(700))
}

func BenchmarkFillSliceInt_1000000(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 1000000), 700)
}
func BenchmarkFillSliceReflect_1000000(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 1000000), int(700))
}
func BenchmarkFillSliceDTypeInt_1000000(b *testing.B) {
	benchmarkFillSliceDType(b, core.NewBuffer(1000000).AsType(core.Int64), int64(700))
}

func benchmarkFillSliceInt(b *testing.B, d []int, v int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(d); i++ {
			d[i] = (v + i) % v / 2
		}
	}
}

func benchmarkFillSliceReflect(b *testing.B, d interface{}, v interface{}) {
	for n := 0; n < b.N; n++ {
		switch reflect.TypeOf(d).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(d)
			switch v := v.(type) {
			case int:
				for i := 0; i < s.Len(); i++ {
					s.Index(i).Set(reflect.ValueOf((v + 1) % v / 2))
				}
			default:
				panic("invalid benchmark input value")
			}
		default:
			panic("invalid benchmark data type")
		}
	}
}

func benchmarkFillSliceDType(b *testing.B, buf *core.Buffer, v interface{}) {
	typ, src := core.Destruct(v)
	if typ != buf.DType() {
		panic("invalid benchmark argument types")
	}

	for n := 0; n < b.N; n++ {
		buf.Iterate(func(_ int, dst unsafe.Pointer) {
			typ.Setraw(dst, src)
		})
	}
}
