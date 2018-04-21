package index_test

import (
	"testing"

	"github.com/ppknap/tensor/internal/index"
)

func BenchmarkIterateRaw_1D_1E(b *testing.B)  { bckIterateRaw(b, []int{1}) }
func BenchmarkIterateAt_1D_1E(b *testing.B)   { bckIterateAt(b, []int{1}) }
func BenchmarkIterateRaw_1D_1kE(b *testing.B) { bckIterateRaw(b, []int{1000}) }
func BenchmarkIterateAt_1D_1kE(b *testing.B)  { bckIterateAt(b, []int{1000}) }
func BenchmarkIterateRaw_1D_1ME(b *testing.B) { bckIterateRaw(b, []int{1000000}) }
func BenchmarkIterateAt_1D_1ME(b *testing.B)  { bckIterateAt(b, []int{1000000}) }

func BenchmarkIterateRaw_2D_1E(b *testing.B)   { bckIterateRaw(b, []int{1, 1}) }
func BenchmarkIterateAt_2D_1E(b *testing.B)    { bckIterateAt(b, []int{1, 1}) }
func BenchmarkIterateRaw_2D_100E(b *testing.B) { bckIterateRaw(b, []int{100, 100}) }
func BenchmarkIterateAt_2D_100E(b *testing.B)  { bckIterateAt(b, []int{100, 100}) }
func BenchmarkIterateRaw_2D_1kE(b *testing.B)  { bckIterateRaw(b, []int{1000, 1000}) }
func BenchmarkIterateAt_2D_1kE(b *testing.B)   { bckIterateAt(b, []int{1000, 1000}) }

func BenchmarkIterateRaw_3D_1E(b *testing.B)   { bckIterateRaw(b, []int{1, 1, 1}) }
func BenchmarkIterateAt_3D_1E(b *testing.B)    { bckIterateAt(b, []int{1, 1, 1}) }
func BenchmarkIterateRaw_3D_10E(b *testing.B)  { bckIterateRaw(b, []int{10, 10, 10}) }
func BenchmarkIterateAt_3D_10E(b *testing.B)   { bckIterateAt(b, []int{10, 10, 10}) }
func BenchmarkIterateRaw_3D_100E(b *testing.B) { bckIterateRaw(b, []int{100, 100, 100}) }
func BenchmarkIterateAt_3D_100E(b *testing.B)  { bckIterateAt(b, []int{100, 100, 100}) }

func bckIterateRaw(b *testing.B, shape []int) {
	idx, sum := index.NewIndex(shape, index.IdxSchemeColMajor), 0
	for n := 0; n < b.N; n++ {
		idx.Iterate(func(pos []int) {
			sum = len(pos)
		})

	}
	_ = sum
}

func bckIterateAt(b *testing.B, shape []int) {
	idx, sum := index.NewIndex(shape, index.IdxSchemeColMajor), 0

	at := idx.At()
	for n := 0; n < b.N; n++ {
		idx.Iterate(func(pos []int) {
			sum = at(pos)
		})

	}
	_ = sum
}
