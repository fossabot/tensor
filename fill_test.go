package tacvs_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs"
)

func TestTensorFull(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Val    complex128
	}{
		"zeros": {
			Tensor: tensorEnum(2, 3, 4).Zeros(),
			Val:    0,
		},
		"ones": {
			Tensor: tensorEnum(2, 2).Ones(),
			Val:    1,
		},
		"real full": {
			Tensor: tensorEnum(2, 2).Full(34),
			Val:    34,
		},
		"complex full": {
			Tensor: tensorEnum(2, 2).Full(1 + 45i),
			Val:    1 + 45i,
		},
		"empty full": {
			Tensor: tacvs.NewTensor().Full(34),
			Val:    0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			data := test.Tensor.Data()
			buf := make([]complex128, len(data))

			for i := range buf {
				buf[i] = test.Val
			}

			if !reflect.DeepEqual(buf, data) {
				t.Errorf("want val=%v; got %v", buf, data)
			}
		})
	}
}

func TestTensorEye(t *testing.T) {
	// val2Pos stores expected value at given positions.
	type val2Pos map[complex128][][]int

	tests := map[string]struct {
		Tensor   *tacvs.Tensor
		Sum      complex128
		ValAtPos val2Pos
	}{
		"empty": {
			Tensor:   tacvs.NewTensor().Eye(),
			Sum:      0,
			ValAtPos: nil,
		},
		"scalar": {
			Tensor: tensorEnum(1).Eye(),
			Sum:    1,
			ValAtPos: val2Pos{
				1: [][]int{{0}},
			},
		},
		"vector": {
			Tensor: tensorEnum(5).Eye(),
			Sum:    1,
			ValAtPos: val2Pos{
				1: [][]int{{0}},
			},
		},
		"matrix": {
			Tensor: tensorEnum(3, 3).Eye(),
			Sum:    3,
			ValAtPos: val2Pos{
				1: [][]int{{0, 0}, {1, 1}, {2, 2}},
			},
		},
		"tensor": {
			Tensor: tensorEnum(4, 4, 4).Eye(),
			Sum:    4,
			ValAtPos: val2Pos{
				1: [][]int{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if sum := test.Tensor.Sum(); sum != test.Sum {
				t.Fatalf("want element sum=%v; got %v", test.Sum, sum)
			}

			for want, poss := range test.ValAtPos {
				for _, pos := range poss {
					if val := test.Tensor.At(pos...); val != want {
						t.Errorf("want val=%v; got %v (pos:%v)", want, val, pos)
					}
				}
			}
		})
	}
}

func TestTensorFill(t *testing.T) {
	// val2Pos stores expected value at given position.
	type val2Pos map[complex128][]int

	tests := map[string]struct {
		Tensor   *tacvs.Tensor
		Vs       []complex128
		ValAtPos val2Pos
	}{
		"column vector": {
			Tensor: tensorEnum(3, 1),
			Vs:     []complex128{4, 5, 6},
			ValAtPos: val2Pos{
				4: []int{0},
				5: []int{1},
				6: []int{2},
			},
		},
		"row vector": {
			Tensor: tensorEnum(1, 4),
			Vs:     []complex128{5, 6, 7, 8},
			ValAtPos: val2Pos{
				5: []int{0, 0},
				6: []int{0, 1},
				7: []int{0, 2},
				8: []int{0, 3},
			},
		},
		"matrix": {
			Tensor: tensorEnum(2, 2),
			Vs:     []complex128{7, 8, 9, 0},
			ValAtPos: val2Pos{
				7: []int{0, 0},
				8: []int{1, 0},
				9: []int{0, 1},
				0: []int{1, 1},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.Fill(test.Vs)

			for want, pos := range test.ValAtPos {
				if val := tr.At(pos...); val != want {
					t.Errorf("want val=%v; got %v (pos:%v)", want, val, pos)
				}
			}
		})
	}
}

func TestTensorFillPanic(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Vs     []complex128
	}{
		"invalid source size": {
			Tensor: tensorEnum(2, 2),
			Vs:     []complex128{1, 2, 3},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("test should have panicked, but it did not")
				}
			}()

			_ = test.Tensor.Fill(test.Vs)
		})
	}
}
