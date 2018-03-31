package index_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs/internal/index"
)

func TestIndexDimensions(t *testing.T) {
	tests := map[string]struct {
		Index   *index.Index
		NDim    int
		Size    int
		Strides []int
		Shape   []int
	}{
		"aaa": {
			Index:   nil,
			NDim:    0,
			Size:    0,
			Strides: nil,
			Shape:   nil,
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
			Index:  nil,
			Pos:    nil,
			Valid:  false,
			Offset: 0,
		},
	}

	for i, test := range tests {
		if valid := test.Index.Validate(test.Pos); valid != test.Valid {
			t.Errorf("want valid=%t; got %t (i:%d)", test.Valid, valid, i)
		}

		if offset := test.Index.At(test.Pos); offset != test.Offset {
			t.Errorf("want offset=%d; got %d (i:%d)", test.Offset, offset, i)
		}
	}
}
