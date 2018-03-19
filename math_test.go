package tacvs_test

import (
	"testing"

	"github.com/ppknap/tacvs"
)

func TestTensorSum(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Sum    complex128
	}{
		"zero sum": {
			Tensor: tacvs.NewTensor(2, 3, 4),
			Sum:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2),
			Sum:    6,
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{1i, 2i, 3i, 4i}),
			Sum:    10i,
		},
		"empty": {
			Tensor: tacvs.NewTensor(),
			Sum:    0,
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 0),
			Sum:    5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if sum := test.Tensor.Sum(); sum != test.Sum {
				t.Fatalf("want sum=%v; got %v", test.Sum, sum)
			}
		})
	}
}
