package tacvs_test

import (
	"reflect"
	"strconv"
	"strings"
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
		"each": {
			Tensor: tensorEnum(2, 2).Each(func(complex128) complex128 { return 2 }),
			Val:    2,
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

func TestTensorIm(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Data   []complex128
	}{
		"empty": {
			Tensor: &tacvs.Tensor{},
			Data:   nil,
		},
		"matrix": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{1 + 1i, 2 + 2i, 3 + 3i, 4 + 4i}),
			Data:   []complex128{1i, 2i, 3i, 4i},
		},
		"vector": {
			Tensor: tacvs.NewTensor(4).Fill([]complex128{1 + 1i, 2 + 2i, 3 + 3i, 4 + 4i}),
			Data:   []complex128{1i, 2i, 3i, 4i},
		},
		"matrix slice": {
			Tensor: tacvs.NewTensor(1, 2).Fill([]complex128{1 + 1i, 2 + 2i}).Slice(1, 1, 2),
			Data:   []complex128{1 + 1i, 2i},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.Im()

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
			}
		})
	}
}
func TestTensorApply(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		CallN  int
		Data   []complex128
	}{
		"matrix": {
			Tensor: tacvs.NewTensor(3, 3),
			CallN:  9,
			Data:   []complex128{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		"empty": {
			Tensor: &tacvs.Tensor{},
			CallN:  0,
			Data:   nil,
		},
		"vector": {
			Tensor: tacvs.NewTensor(4),
			CallN:  4,
			Data:   []complex128{1, 1, 1, 1},
		},
		"matrix horizontal slice": {
			Tensor: tacvs.NewTensor(3, 3).Slice(0, 1, 2),
			CallN:  3,
			Data:   []complex128{0, 1, 0, 0, 1, 0, 0, 1, 0},
		},
		"matrix vertical slice": {
			Tensor: tacvs.NewTensor(3, 3).Slice(1, 1, 2),
			CallN:  3,
			Data:   []complex128{0, 0, 0, 1, 1, 1, 0, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			callN, callM := 0, map[string]int{}

			tr := test.Tensor.Apply(func(tr *tacvs.Tensor, idx []int) {
				idxStr := make([]string, len(idx))
				for i := range idx {
					idxStr[i] = strconv.Itoa(idx[i])
				}

				callN++
				callM[strings.Join(idxStr, "|")]++

				tr.Set(1, idx...)
			})

			if callN != test.CallN {
				t.Fatalf("want calls number=%d; got %d", test.CallN, callN)
			}

			if callN != len(callM) {
				t.Fatalf("multiple calls with the same index: %v", callM)
			}

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
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
		"fill on view": {
			Tensor: tensorEnum(2, 2).Slice(0, 0),
			Vs:     []complex128{1, 2, 3, 4},
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
