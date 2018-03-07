package tacvs_test

import (
	"testing"

	"github.com/ppknap/tacvs"
)

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
