package index_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tensor/internal/index"
)

func TestIndexDimensions(t *testing.T) {
	tests := map[string]struct {
		Index   *index.Index
		NDim    int
		Size    int
		Strides []int
		Shape   []int
		IsView  bool
	}{
		"scalar": {
			Index:   index.NewIndex(nil, index.DefaultIdxScheme),
			NDim:    0,
			Size:    1,
			Strides: nil,
			Shape:   nil,
			IsView:  false,
		},
		"vector": {
			Index:   index.NewIndex([]int{5}, index.DefaultIdxScheme),
			NDim:    1,
			Size:    5,
			Strides: []int{1},
			Shape:   []int{5},
			IsView:  false,
		},
		"vector zero": {
			Index:   index.NewIndex([]int{0}, index.DefaultIdxScheme),
			NDim:    1,
			Size:    0,
			Strides: []int{1},
			Shape:   []int{0},
			IsView:  false,
		},
		"vector slice": {
			Index:   index.NewIndex([]int{6}, index.DefaultIdxScheme).Slice(0, 4),
			NDim:    1,
			Size:    2,
			Strides: []int{1},
			Shape:   []int{2},
			IsView:  true,
		},
		"matrix row": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor),
			NDim:    2,
			Size:    9,
			Strides: []int{1, 3},
			Shape:   []int{3, 3},
			IsView:  false,
		},
		"matrix column": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor),
			NDim:    2,
			Size:    9,
			Strides: []int{1, 3},
			Shape:   []int{3, 3},
			IsView:  false,
		},
		"matrix zero row": {
			Index:   index.NewIndex([]int{0, 0}, index.IdxSchemeRowMajor),
			NDim:    2,
			Size:    0,
			Strides: []int{1, 1},
			Shape:   []int{0, 0},
			IsView:  false,
		},
		"matrix zero column": {
			Index:   index.NewIndex([]int{0, 0}, index.IdxSchemeColMajor),
			NDim:    2,
			Size:    0,
			Strides: []int{1, 1},
			Shape:   []int{0, 0},
			IsView:  false,
		},
		"matrix slice row": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).Slice(0, 1, 2),
			NDim:    2,
			Size:    3,
			Strides: []int{1, 3},
			Shape:   []int{1, 3},
			IsView:  true,
		},
		"matrix slice column": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).Slice(0, 1, 2),
			NDim:    2,
			Size:    3,
			Strides: []int{1, 3},
			Shape:   []int{1, 3},
			IsView:  true,
		},
		"matrix slice of slice row": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).Slice(1, 2).Slice(0, 1),
			NDim:    2,
			Size:    2,
			Strides: []int{1, 3},
			Shape:   []int{2, 1},
			IsView:  true,
		},
		"matrix slice of slice column": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).Slice(1, 2).Slice(0, 1),
			NDim:    2,
			Size:    2,
			Strides: []int{1, 3},
			Shape:   []int{2, 1},
			IsView:  true,
		},
		"tensor row": {
			Index:   index.NewIndex([]int{2, 2, 2}, index.IdxSchemeRowMajor),
			NDim:    3,
			Size:    8,
			Strides: []int{1, 2, 4},
			Shape:   []int{2, 2, 2},
			IsView:  false,
		},
		"tensor column": {
			Index:   index.NewIndex([]int{2, 2, 2}, index.IdxSchemeColMajor),
			NDim:    3,
			Size:    8,
			Strides: []int{1, 2, 4},
			Shape:   []int{2, 2, 2},
			IsView:  false,
		},
		"tensor slice row": {
			Index:   index.NewIndex([]int{2, 2, 2}, index.IdxSchemeRowMajor).Slice(2, 1),
			NDim:    3,
			Size:    4,
			Strides: []int{1, 2, 4},
			Shape:   []int{2, 2, 1},
			IsView:  true,
		},
		"tensor slice column": {
			Index:   index.NewIndex([]int{2, 2, 2}, index.IdxSchemeColMajor).Slice(2, 1),
			NDim:    3,
			Size:    4,
			Strides: []int{1, 2, 4},
			Shape:   []int{2, 2, 1},
			IsView:  true,
		},
		"scalar from matrix row": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).Scalar([]int{1, 1}),
			NDim:    0,
			Size:    1,
			Strides: nil,
			Shape:   nil,
			IsView:  true,
		},
		"scalar from matrix column": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).Scalar([]int{1, 1}),
			NDim:    0,
			Size:    1,
			Strides: nil,
			Shape:   nil,
			IsView:  true,
		},
		"matrix view row": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).View(),
			NDim:    2,
			Size:    9,
			Strides: []int{1, 3},
			Shape:   []int{3, 3},
			IsView:  true,
		},
		"matrix view column": {
			Index:   index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).View(),
			NDim:    2,
			Size:    9,
			Strides: []int{1, 3},
			Shape:   []int{3, 3},
			IsView:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if nDim := test.Index.NDim(); nDim != test.NDim {
				t.Errorf("want nDim=%d; got %d", test.NDim, nDim)
			}

			if size := test.Index.Size(); size != test.Size {
				t.Errorf("want size=%d; got %d", test.Size, size)
			}

			if strides := test.Index.Strides(); !reflect.DeepEqual(strides, test.Strides) {
				t.Errorf("want strides=%v; got %v", test.Strides, strides)
			}

			if shape := test.Index.Shape(); !reflect.DeepEqual(shape, test.Shape) {
				t.Errorf("want shape=%v; got %v", test.Shape, shape)
			}

			if isView := test.Index.Flags().IsView(); isView != test.IsView {
				t.Errorf("want isView=%t; got %t", test.IsView, isView)
			}
		})
	}
}

func TestIndexOffset(t *testing.T) {
	tests := []struct {
		Index  *index.Index
		Pos    []int
		Valid  bool
		Offset int
	}{
		{
			// 0 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{0, 0},
			Valid:  true,
			Offset: 0,
		},
		{
			// 1 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{0, 0},
			Valid:  true,
			Offset: 0,
		},
		{
			// 2 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{3, 1},
			Valid:  true,
			Offset: 7,
		},
		{
			// 3 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{3, 1},
			Valid:  true,
			Offset: 7,
		},
		{
			// 4 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{3, 2},
			Valid:  false,
			Offset: -1,
		},
		{
			// 5 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{3, 2},
			Valid:  false,
			Offset: -1,
		},
		{
			// 6 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{4, 1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 7 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{4, 1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 8 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{1, 1, 1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 9 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{1, 1, 1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 10 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 11 //
			Index:  index.NewIndex([]int{4, 2}, index.IdxSchemeColMajor),
			Pos:    []int{1},
			Valid:  false,
			Offset: -1,
		},
		{
			// 12 //
			Index:  index.NewIndex([]int{4}, index.IdxSchemeRowMajor),
			Pos:    []int{2},
			Valid:  true,
			Offset: 2,
		},
		{
			// 13 //
			Index:  index.NewIndex([]int{4}, index.IdxSchemeColMajor),
			Pos:    []int{2},
			Valid:  true,
			Offset: 2,
		},
		{
			// 14 //
			Index:  index.NewIndex([]int{}, index.IdxSchemeRowMajor),
			Pos:    []int{0},
			Valid:  false,
			Offset: -1,
		},
		{
			// 15 //
			Index:  index.NewIndex([]int{}, index.IdxSchemeColMajor),
			Pos:    []int{0},
			Valid:  false,
			Offset: -1,
		},
		{
			// 16 //
			Index:  index.NewIndex([]int{2, 3, 2}, index.IdxSchemeRowMajor),
			Pos:    []int{1, 1, 1},
			Valid:  true,
			Offset: 9,
		},
		{
			// 17 //
			Index:  index.NewIndex([]int{2, 3, 2}, index.IdxSchemeColMajor),
			Pos:    []int{1, 1, 1},
			Valid:  true,
			Offset: 9,
		},
		{
			// 18 //
			Index:  index.NewIndex([]int{5, 3}, index.IdxSchemeRowMajor).Slice(0, 2, 4),
			Pos:    []int{0, 2},
			Valid:  true,
			Offset: 12,
		},
		{
			// 19 //
			Index:  index.NewIndex([]int{5, 3}, index.IdxSchemeColMajor).Slice(0, 2, 4),
			Pos:    []int{0, 2},
			Valid:  true,
			Offset: 12,
		},
		{
			// 20 //
			Index:  index.NewIndex([]int{5, 3}, index.IdxSchemeRowMajor).Slice(0, 2, 4).Slice(1, 1),
			Pos:    []int{0, 0},
			Valid:  true,
			Offset: 7,
		},
		{
			// 21 //
			Index:  index.NewIndex([]int{5, 3}, index.IdxSchemeColMajor).Slice(0, 2, 4).Slice(1, 1),
			Pos:    []int{0, 0},
			Valid:  true,
			Offset: 7,
		},
		{
			// 22 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).Scalar([]int{1, 1}),
			Pos:    nil,
			Valid:  true,
			Offset: 4,
		},
		{
			// 23 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).Scalar([]int{1, 1}),
			Pos:    nil,
			Valid:  true,
			Offset: 4,
		},
		{
			// 24 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor).Scalar([]int{1, 2}),
			Pos:    []int{},
			Valid:  true,
			Offset: 7,
		},
		{
			// 25 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor).Scalar([]int{1, 2}),
			Pos:    []int{},
			Valid:  true,
			Offset: 7,
		},
		{
			// 26 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeRowMajor),
			Pos:    []int{0},
			Valid:  false,
			Offset: -1,
		},
		{
			// 27 //
			Index:  index.NewIndex([]int{3, 3}, index.IdxSchemeColMajor),
			Pos:    []int{0},
			Valid:  false,
			Offset: -1,
		},
		{
			// 28 //
			Index:  nil,
			Pos:    []int{0},
			Valid:  false,
			Offset: -1,
		},
	}

	for i, test := range tests {
		valid := test.Index.Validate(test.Pos)
		if valid != test.Valid {
			t.Errorf("want valid=%t; got %t (i:%d)", test.Valid, valid, i)
		}

		if !valid {
			continue
		}

		if offset := test.Index.At()(test.Pos); offset != test.Offset {
			t.Errorf("want offset=%d; got %d (i:%d)", test.Offset, offset, i)
		}
	}
}

func TestIndexIterate(t *testing.T) {
	tests := map[string]struct {
		Index   *index.Index
		Indices [][]int
	}{
		"scalar": {
			Index:   index.NewIndex([]int{}, index.IdxSchemeColMajor),
			Indices: nil,
		},
		"vector": {
			Index: index.NewIndex([]int{5}, index.IdxSchemeColMajor),
			Indices: [][]int{
				{0}, {1}, {2}, {3}, {4},
			},
		},
		"vector one element": {
			Index: index.NewIndex([]int{1}, index.IdxSchemeColMajor),
			Indices: [][]int{
				{0},
			},
		},
		"square matrix": {
			Index: index.NewIndex([]int{2, 2}, index.IdxSchemeColMajor),
			Indices: [][]int{
				{0, 0}, {0, 1}, {1, 0}, {1, 1},
			},
		},
		"view zero size": {
			Index:   index.NewIndex([]int{2, 2}, index.IdxSchemeColMajor).Slice(0, 2),
			Indices: nil,
		},
		"tensor": {
			Index: index.NewIndex([]int{3, 2, 1}, index.IdxSchemeColMajor),
			Indices: [][]int{
				{0, 0, 0}, {0, 1, 0}, {1, 0, 0}, {1, 1, 0}, {2, 0, 0}, {2, 1, 0},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var indices [][]int

			test.Index.Iterate(func(idx []int) {
				cp := make([]int, len(idx))
				copy(cp, idx)

				indices = append(indices, cp)
			})

			if !reflect.DeepEqual(indices, test.Indices) {
				t.Errorf("want indices=%v; got %v", test.Indices, indices)
			}
		})
	}
}
