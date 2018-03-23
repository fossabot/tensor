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
		"slice": {
			Tensor: tensorEnum(3, 3).Slice(0, 0).Full(99),
			Val:    99,
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

func TestTensorArrange(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Start  complex128
		Step   complex128
		Data   []complex128
	}{
		"empty": {
			Tensor: &tacvs.Tensor{},
			Start:  3,
			Step:   1,
			Data:   nil,
		},
		"vector": {
			Tensor: tacvs.NewTensor(5),
			Start:  0,
			Step:   1,
			Data:   []complex128{0, 1, 2, 3, 4},
		},
		"matrix": {
			Tensor: tacvs.NewTensor(2, 2),
			Start:  1,
			Step:   7,
			Data:   []complex128{1, 8, 15, 22},
		},
		"matrix slice": {
			Tensor: tacvs.NewTensor(3, 3).Slice(1, 1, 2),
			Start:  1,
			Step:   4,
			Data:   []complex128{0, 0, 0, 1, 5, 9, 0, 0, 0},
		},
		"constant": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  3,
			Step:   3,
			Data:   []complex128{3, 6, 9, 12, 15, 18},
		},
		"zero step": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  3,
			Step:   0,
			Data:   []complex128{3, 3, 3, 3, 3, 3},
		},
		"negative": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  4,
			Step:   -1,
			Data:   []complex128{4, 3, 2, 1, 0, -1},
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  -2 + 10i,
			Step:   1 - 1i,
			Data:   []complex128{-2 + 10i, -1 - 9i, 0 - 8i, 1 - 7i, 2 - 6i, 3 - 5i},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.Arrange(test.Start, test.Step)

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
			}
		})
	}
}

func TestTensorLinspace(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Start  complex128
		End    complex128
		Data   []complex128
	}{
		"empty": {
			Tensor: &tacvs.Tensor{},
			Start:  3,
			End:    10,
			Data:   nil,
		},
		"vector": {
			Tensor: tacvs.NewTensor(5),
			Start:  0,
			End:    10,
			Data:   []complex128{0, 2.5, 5, 7.5, 10},
		},
		"matrix": {
			Tensor: tacvs.NewTensor(2, 2),
			Start:  1,
			End:    7,
			Data:   []complex128{1, 3, 5, 7},
		},
		"matrix slice": {
			Tensor: tacvs.NewTensor(3, 3).Slice(1, 1, 2),
			Start:  1,
			End:    9,
			Data:   []complex128{0, 0, 0, 1, 5, 9, 0, 0, 0},
		},
		"constant": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  3,
			End:    3,
			Data:   []complex128{3, 3, 3, 3, 3, 3},
		},
		"negative": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  10,
			End:    -15,
			Data:   []complex128{10, 5, 0, -5, -10, -15},
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 3),
			Start:  -2 + 10i,
			End:    3 - 15i,
			Data:   []complex128{-2 + 10i, -1 + 5i, 0, 1 - 5i, 2 - 10i, 3 - 15i},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.Linspace(test.Start, test.End)

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
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

func TestTensorRandom(t *testing.T) {
	tr := tensorEnum(100, 100, 100).Random(nil)

	for i, val := range tr.Data() {
		if re := real(val); re < 0 || re >= 1 {
			t.Fatalf("want real part in [0, 1); got %v (i:%d)", re, i)
		}

		if im := imag(val); im != 0 {
			t.Fatalf("want imaginary part=0; got %v (i:%d)", im, i)
		}
	}
}

type testRandSource int64

func (trs testRandSource) Int63() int64 { return int64(trs) }
func (trs testRandSource) Seed(int64)   {}

func TestTensorRandomSource(t *testing.T) {
	tr := tensorEnum(20, 10, 5).Random(testRandSource(123456))

	first := tr.At(0, 0, 0)
	if imag(first) != 0 {
		t.Fatalf("want values to be real; got %v", first)
	}

	for _, val := range tr.Data() {
		if val != first {
			t.Fatalf("want the same values; got %v!=%v", first, val)
		}
	}
}

func TestTensorRe(t *testing.T) {
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
			Data:   []complex128{1, 2, 3, 4},
		},
		"vector": {
			Tensor: tacvs.NewTensor(4).Fill([]complex128{1 + 1i, 2 + 2i, 3 + 3i, 4 + 4i}),
			Data:   []complex128{1, 2, 3, 4},
		},
		"matrix slice": {
			Tensor: tacvs.NewTensor(1, 2).Fill([]complex128{1 + 1i, 2 + 2i}).Slice(1, 1, 2),
			Data:   []complex128{1 + 1i, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.Re()

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
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

func TestTensorCumSum(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Data   []complex128
	}{
		"empty": {
			Tensor: tacvs.NewTensor(),
			Data:   nil,
		},
		"vector": {
			Tensor: tacvs.NewTensor(3).Fill([]complex128{-2, 4, 30}),
			Data:   []complex128{-2, 2, 32},
		},
		"zero max": {
			Tensor: tacvs.NewTensor(2, 3),
			Data:   []complex128{0, 0, 0, 0, 0, 0},
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 3, 4}),
			Data:   []complex128{8, 10, 13, 17},
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{1i, 4i, 3i, 2i}),
			Data:   []complex128{1i, 5i, 8i, 10i},
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
			Data:   []complex128{0, 0, 2, 5},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tr := test.Tensor.CumSum()

			if data := tr.Data(); !reflect.DeepEqual(data, test.Data) {
				t.Fatalf("want data=%v; got %v", test.Data, data)
			}
		})
	}
}