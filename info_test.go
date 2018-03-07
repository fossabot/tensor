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
