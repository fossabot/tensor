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
		"empty": {
			Tensor: tacvs.NewTensor(),
			Sum:    0,
		},
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
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
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

func TestTensorMin(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Min    complex128
	}{
		"empty": {
			Tensor: tacvs.NewTensor(),
			Min:    0,
		},
		"vector": {
			Tensor: tacvs.NewTensor(3).Fill([]complex128{-2, 4, 30}),
			Min:    -1,
		},
		"zero min": {
			Tensor: tacvs.NewTensor(2, 3, 4),
			Min:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 3, 4}),
			Min:    2,
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{1i, 2i, 3i, 4i}),
			Min:    2i,
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
			Min:    2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if min := test.Tensor.Sum(); min != test.Min {
				t.Fatalf("want min=%v; got %v", test.Min, min)
			}
		})
	}
}

func TestTensorMax(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Max    complex128
	}{
		"empty": {
			Tensor: tacvs.NewTensor(),
			Max:    0,
		},
		"vector": {
			Tensor: tacvs.NewTensor(3).Fill([]complex128{-2, 4, 30}),
			Max:    30,
		},
		"zero max": {
			Tensor: tacvs.NewTensor(2, 3, 4),
			Max:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 3, 4}),
			Max:    8,
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{1i, 4i, 3i, 2i}),
			Max:    4i,
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
			Max:    3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if max := test.Tensor.Max(); max != test.Max {
				t.Fatalf("want max=%v; got %v", test.Max, max)
			}
		})
	}
}

func TestTensorMean(t *testing.T) {
	tests := map[string]struct {
		Tensor *tacvs.Tensor
		Mean   complex128
	}{
		"empty": {
			Tensor: tacvs.NewTensor(),
			Mean:   0,
		},
		"vector": {
			Tensor: tacvs.NewTensor(3).Fill([]complex128{-2, 4, 31}),
			Mean:   11,
		},
		"zero max": {
			Tensor: tacvs.NewTensor(2, 3, 4),
			Mean:   0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 6, 4}),
			Mean:   5,
		},
		"complex": {
			Tensor: tacvs.NewTensor(2, 2).Fill([]complex128{2 + 1i, 1 - 4i, 4 + 3i, 9 + 2i}),
			Mean:   4 + 0.5i,
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
			Mean:   2.5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if mean := test.Tensor.Mean(); mean != test.Mean {
				t.Fatalf("want mean=%v; got %v", test.Mean, mean)
			}
		})
	}
}
