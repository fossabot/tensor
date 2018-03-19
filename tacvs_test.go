package tacvs_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs"
)

func TestTensorInitializationPanic(t *testing.T) {
	tests := map[string]struct {
		Init func()
	}{
		"negative axis size": {
			Init: func() {
				_ = tacvs.NewTensor(3, -2)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("test should have panicked, but it did not")
				}
			}()

			test.Init()
		})
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
			Pos:    []int{0},
			Val:    0,
		},
		{
			// 17 //
			Tensor: tensorEnum(2, 2, 3),
			Pos:    []int{1},
			Val:    1,
		},
	}

	for i, test := range tests {
		val := test.Tensor.At(test.Pos...)
		if val != test.Val {
			t.Errorf("want pos=%v; got %v (i:%d)", test.Val, val, i)
		}
	}
}

func TestTensorIndexingPanic(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Pos    []int
	}{
		"too many nonzero indexes": {
			Tensor: tensorEnum(2, 2),
			Pos:    []int{0, 1, 1},
		},
		"negative index": {
			Tensor: tensorEnum(2),
			Pos:    []int{-1},
		},
		"nil index": {
			Tensor: tensorEnum(2),
			Pos:    nil,
		},
		"out of range index": {
			Tensor: tensorEnum(2, 2),
			Pos:    []int{3, 0},
		},
		"out of range second index": {
			Tensor: tensorEnum(2, 2),
			Pos:    []int{0, 3},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("test should have panicked, but it did not")
				}
			}()

			_ = test.Tensor.At(test.Pos...)
		})
	}
}

func TestTensorSplit(t *testing.T) {
	// nVal2Pos stores slice index and expected values at given positions.
	type nVal2Pos []map[complex128][]int

	tests := []struct {
		Tensor   *tacvs.Tensor
		Dim      int
		Shape    []int
		ValAtPos nVal2Pos
	}{
		{
			// 0 //
			Tensor: tensorEnum(3, 3),
			Dim:    0,
			Shape:  []int{1, 3},
			ValAtPos: nVal2Pos{
				{
					0: []int{0},
					6: []int{0, 2},
				},
				{
					1: []int{0},
					7: []int{0, 2},
				},
				{
					2: []int{0},
					8: []int{0, 2},
				},
			},
		},
		{
			// 1 //
			Tensor: tensorEnum(3),
			Dim:    0,
			Shape:  []int{1, 1},
			ValAtPos: nVal2Pos{
				{
					0: []int{0},
				},
				{
					1: []int{0},
				},
				{
					2: []int{0},
				},
			},
		},
		{
			// 2 //
			Tensor: tensorEnum(3),
			Dim:    1,
			Shape:  []int{3, 1},
			ValAtPos: nVal2Pos{
				{
					0: []int{0},
					1: []int{1},
					2: []int{2},
				},
			},
		},
		{
			// 3 //
			Tensor: tensorEnum(1, 4),
			Dim:    0,
			Shape:  []int{1, 4},
			ValAtPos: nVal2Pos{
				{
					0: []int{0, 0},
					1: []int{0, 1},
					2: []int{0, 2},
					3: []int{0, 3},
				},
			},
		},
		{
			// 4 //
			Tensor: tensorEnum(1, 4),
			Dim:    1,
			Shape:  []int{1, 1},
			ValAtPos: nVal2Pos{
				{
					0: []int{0},
				},
				{
					1: []int{0},
				},
				{
					2: []int{0},
				},
				{
					3: []int{0},
				},
			},
		},
		{
			// 5 //
			Tensor: tensorEnum(3, 2, 2),
			Dim:    2,
			Shape:  []int{3, 2},
			ValAtPos: nVal2Pos{
				{
					2: []int{2, 0},
					3: []int{0, 1},
					5: []int{2, 1},
				},
				{
					6:  []int{0, 0},
					8:  []int{2, 0},
					11: []int{2, 1},
				},
			},
		},
	}

	for i, test := range tests {
		slices := test.Tensor.Split(test.Dim)

		if ls, want := len(slices), len(test.ValAtPos); ls != want {
			t.Errorf("want slice length=%d; got %d (i:%d)", want, ls, i)
			continue
		}

		if len(slices) == 0 {
			t.Fatalf("split results in zero tensor (i:%d)", i)
		}

		if shape := slices[0].Shape(); !reflect.DeepEqual(shape, test.Shape) {
			t.Errorf("want tensor shape=%v; got %v (i:%d)", test.Shape, shape, i)
		}

		for j, val2pos := range test.ValAtPos {
			for want, pos := range val2pos {
				if val := slices[j].At(pos...); val != want {
					t.Errorf("want val=%v; got %v (i:%d,j:%d,pos:%v)", want, val, i, j, pos)
				}
			}
		}
	}
}

func TestTensorSplitPanic(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Dim    int
	}{
		"invalid dimension": {
			Tensor: tensorEnum(2, 2),
			Dim:    3,
		},
		"negative dimension": {
			Tensor: tensorEnum(2, 2),
			Dim:    -1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("test should have panicked, but it did not")
				}
			}()

			_ = test.Tensor.Split(test.Dim)
		})
	}
}

func TestTensorSlice(t *testing.T) {
	// val2Pos stores expected value at given position.
	type val2Pos map[complex128][]int

	tests := []struct {
		Tensor   *tacvs.Tensor
		Args     []int
		Shape    []int
		ValAtPos val2Pos
	}{
		{
			// 0 //
			Tensor: tensorEnum(4, 2),
			Args:   []int{0, 1},
			Shape:  []int{3, 2},
			ValAtPos: val2Pos{
				1: []int{0},
				7: []int{2, 1},
			},
		},
		{
			// 1 //
			Tensor: tensorEnum(4, 1),
			Args:   []int{0, 0, 3},
			Shape:  []int{3, 1},
			ValAtPos: val2Pos{
				0: []int{0},
				2: []int{2},
			},
		},
		{
			// 2 //
			Tensor: tensorEnum(2, 1),
			Args:   []int{0, 0, 2},
			Shape:  []int{2, 1},
			ValAtPos: val2Pos{
				0: []int{0},
				1: []int{1},
			},
		},
		{
			// 3 //
			Tensor: tensorEnum(2, 3),
			Args:   []int{1, 1, 2},
			Shape:  []int{2, 1},
			ValAtPos: val2Pos{
				2: []int{0},
				3: []int{1},
			},
		},
		{
			// 4 //
			Tensor:   tensorEnum(1, 2),
			Args:     []int{1, 1, 1},
			Shape:    []int{1, 0},
			ValAtPos: nil,
		},
		{
			// 5 //
			Tensor: tensorEnum(1, 3),
			Args:   []int{1, 2},
			Shape:  []int{1, 1},
			ValAtPos: val2Pos{
				2: []int{0},
			},
		},
		{
			// 6 //
			Tensor: tensorEnum(2, 2, 2),
			Args:   []int{2, 1},
			Shape:  []int{2, 2},
			ValAtPos: val2Pos{
				4: []int{0, 0},
				7: []int{1, 1},
			},
		},
		{
			// 7 //
			Tensor: tensorEnum(2, 1, 4),
			Args:   []int{2, 1, 3},
			Shape:  []int{2, 1, 2},
			ValAtPos: val2Pos{
				2: []int{0, 0, 0},
				3: []int{1, 0, 0},
				4: []int{0, 0, 1},
				5: []int{1, 0, 1},
			},
		},
	}

	for i, test := range tests {
		slice := test.Tensor.Slice(test.Args[0], test.Args[1], test.Args[2:]...)
		if slice == nil {
			t.Fatalf("want non-nil slice (i:%d)", i)
		}

		if shape := slice.Shape(); !reflect.DeepEqual(shape, test.Shape) {
			t.Errorf("want tensor shape=%v; got %v (i:%d)", test.Shape, shape, i)
		}

		for want, pos := range test.ValAtPos {
			if val := slice.At(pos...); val != want {
				t.Errorf("want val=%v; got %v (i:%d,pos:%v)", want, val, i, pos)
			}
		}
	}
}

func TestTensorSlicePanic(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Args   []int
	}{
		"out of range from": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, 5},
		},
		"out of range to": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, 1, 7},
		},
		"invalid slice index": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, 1, 0},
		},
		"too many arguments": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, 1, 2, 3},
		},
		"negative dimension": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{-1, 1},
		},
		"negative from": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, -1},
		},
		"negative to": {
			Tensor: tensorEnum(2, 2),
			Args:   []int{0, 0, -1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("test should have panicked, but it did not")
				}
			}()

			_ = test.Tensor.Slice(test.Args[0], test.Args[1], test.Args[2:]...)
		})
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
