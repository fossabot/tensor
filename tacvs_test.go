package tacvs_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs"
)

func TestTensorInfo(t *testing.T) {
	tests := []struct {
		Tensor   *tacvs.Tensor
		Shape    []int
		Size     int
		DataSize int
	}{
		{
			// 0 //
			Tensor:   tacvs.NewTensor(1, 1),
			Shape:    []int{1},
			Size:     1,
			DataSize: 1,
		},
		{
			// 1 //
			Tensor:   tacvs.NewTensor(1, 2, 3, 4),
			Shape:    []int{1, 2, 3, 4},
			Size:     24,
			DataSize: 24,
		},
		{
			// 2 //
			Tensor:   tacvs.NewTensor(1, 1, 1, 1),
			Shape:    []int{1},
			Size:     1,
			DataSize: 1,
		},
		{
			// 3 //
			Tensor:   tacvs.NewTensor(6, 1, 1, 1, 1),
			Shape:    []int{6},
			Size:     6,
			DataSize: 6,
		},
		{
			// 4 //
			Tensor:   tacvs.NewTensor(),
			Shape:    []int{},
			Size:     0,
			DataSize: 0,
		},
		{
			// 5 //
			Tensor:   &tacvs.Tensor{},
			Shape:    []int{},
			Size:     0,
			DataSize: 0,
		},
		{
			// 6 //
			Tensor:   tacvs.NewTensor(2),
			Shape:    []int{2},
			Size:     2,
			DataSize: 2,
		},
	}

	for i, test := range tests {
		if ndim, want := test.Tensor.NDim(), len(test.Shape); ndim != want {
			t.Errorf("want ndim=%v; got %v (i:%d)", ndim, want, i)
		}

		if shape := test.Tensor.Shape(); !reflect.DeepEqual(shape, test.Shape) {
			t.Errorf("want shape=%v; got %v (i:%d)", test.Shape, shape, i)
		}

		if size := test.Tensor.Size(); size != test.Size {
			t.Errorf("want size=%v; got %v (i:%d)", test.Size, size, i)
		}

		if dataSize := len(test.Tensor.Data()); dataSize != test.DataSize {
			t.Errorf("want data size=%v; got %v (i:%d)", test.DataSize, dataSize, i)
		}
	}
}

func TestTensorIndexing(t *testing.T) {
	tests := []struct {
		Tensor *tacvs.Tensor
		Pos    []int
		Val    complex128
	}{
		{
			// 0 //
			Tensor: tensorEnum(3, 3),
			Pos:    []int{0, 0},
			Val:    0,
		},
		{
			// 1 //
			Tensor: tensorEnum(3, 3),
			Pos:    []int{2, 0},
			Val:    2,
		},
		{
			// 2 //
			Tensor: tensorEnum(3, 3),
			Pos:    []int{2, 2},
			Val:    8,
		},
		{
			// 3 //
			Tensor: tensorEnum(3, 3),
			Pos:    []int{0, 2},
			Val:    6,
		},
		{
			// 4 //
			Tensor: tensorEnum(3, 3),
			Pos:    []int{1, 1},
			Val:    4,
		},
		{
			// 5 //
			Tensor: tensorEnum(6, 1),
			Pos:    []int{0},
			Val:    0,
		},
		{
			// 6 //
			Tensor: tensorEnum(6, 1),
			Pos:    []int{5},
			Val:    5,
		},
		{
			// 7 //
			Tensor: tensorEnum(6, 1),
			Pos:    []int{4},
			Val:    4,
		},
		{
			// 8 //
			Tensor: tensorEnum(1, 6),
			Pos:    []int{0, 0},
			Val:    0,
		},
		{
			// 9 //
			Tensor: tensorEnum(1, 6),
			Pos:    []int{0, 5},
			Val:    5,
		},
		{
			// 10 //
			Tensor: tensorEnum(1, 6),
			Pos:    []int{0, 4},
			Val:    4,
		},
		{
			// 11 //
			Tensor: tensorEnum(2, 2, 3),
			Pos:    []int{1, 0, 0},
			Val:    1,
		},
		{
			// 12 //
			Tensor: tensorEnum(2, 2, 3),
			Pos:    []int{0, 0, 1},
			Val:    4,
		},
		{
			// 13 //
			Tensor: tensorEnum(2, 2, 3),
			Pos:    []int{1, 0, 1},
			Val:    5,
		},
		{
			// 14 //
			Tensor: tensorEnum(2, 2, 3),
			Pos:    []int{1, 1, 2},
			Val:    11,
		},
		{
			// 15 //
			Tensor: tensorEnum(3, 3, 1, 1),
			Pos:    []int{1, 1, 0, 0, 0},
			Val:    4,
		},
		{
			// 16 //
			Tensor: tensorEnum(6, 1, 1, 1, 1),
			Pos:    []int{0, 0, 0, 0},
			Val:    0,
		},
	}

	for i, test := range tests {
		val := test.Tensor.At(test.Pos...)
		if val != test.Val {
			t.Errorf("want pos=%v; got %v (i:%d)", test.Val, val, i)
		}
	}
}

// tensorEnum creates a new tensor and fills its internal buffer with numbers
// that indicate element position in the memory.
func tensorEnum(shape ...int) *tacvs.Tensor {
	t := tacvs.NewTensor(shape...)

	for i, data := 0, t.Data(); i < len(data); i++ {
		data[i] = complex(float64(i), 0)
	}

	return t
}
