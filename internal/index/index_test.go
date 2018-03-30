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
			idx := test.Index

			if nDim := idx.NDim(); nDim != test.NDim {
				t.Errorf("want nDim=%d; got %d", test.NDim, nDim)
			}

			if size := idx.Size(); size != test.Size {
				t.Errorf("want size=%d; got %d", test.Size, size)
			}

			if strides := idx.Strides(); !reflect.DeepEqual(strides, test.Strides) {
				t.Errorf("want strides=%v; got %v", test.Strides, strides)
			}

			if shape := idx.Shape(); !reflect.DeepEqual(shape, test.Shape) {
				t.Errorf("want shape=%v; got %v", test.Shape, shape)
			}
		})
	}
}
