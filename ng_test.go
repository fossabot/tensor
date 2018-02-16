package ng_test

import (
	"testing"

	"github.com/ppknap/cxnn/ng"
)

func TestTensorIndexing(t *testing.T) {
	tests := []struct {
		Tensor *ng.Tensor
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
			Pos:    []int{0, 0},
			Val:    0,
		},
		{
			// 6 //
			Tensor: tensorEnum(6, 1),
			Pos:    []int{5, 0},
			Val:    5,
		},
		{
			// 7 //
			Tensor: tensorEnum(6, 1),
			Pos:    []int{4, 0},
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
	}

	for i, test := range tests {
		val := test.Tensor.At(test.Pos...)
		if val != test.Val {
			t.Errorf("want %v; got %v (i:%d)", test.Val, val, i)
		}
	}
}

// tensorEnum creates a new tensor and fills its internal buffer with numbers
// that indicate element position in the memory.
func tensorEnum(first, second int, rest ...int) *ng.Tensor {
	t := ng.NewTensor(first, second, rest...)

	for i, data := 0, t.Data(); i < len(data); i++ {
		data[i] = complex(float64(i), 0)
	}

	return t
}
