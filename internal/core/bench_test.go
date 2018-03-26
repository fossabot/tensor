package core_test

import (
	"reflect"
	"testing"
)

func BenchmarkFillSliceInt1(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 1), 700)
}
func BenchmarkFillSliceReflect1(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 1), int(700))
}

func BenchmarkFillSliceInt1000(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 100), 700)
}
func BenchmarkFillSliceReflect1000(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 100), int(700))
}

func BenchmarkFillSliceInt1000000(b *testing.B) {
	benchmarkFillSliceInt(b, make([]int, 10000), 700)
}
func BenchmarkFillSliceReflect1000000(b *testing.B) {
	benchmarkFillSliceReflect(b, make([]int, 10000), int(700))
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
