package stale_test

import (
	"math/cmplx"
	"testing"

	"github.com/ppknap/tensor/stale"
)

func TestTensorSum(t *testing.T) {
	tests := map[string]struct {
		Tensor *stale.Tensor
		Sum    complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Sum:    0,
		},
		"zero sum": {
			Tensor: stale.NewTensor(2, 3, 4),
			Sum:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2),
			Sum:    6,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{1i, 2i, 3i, 4i}),
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
		Tensor *stale.Tensor
		Min    complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Min:    0,
		},
		"vector": {
			Tensor: stale.NewTensor(3).Fill([]complex128{-2, 4, 30}),
			Min:    -1,
		},
		"zero min": {
			Tensor: stale.NewTensor(2, 3, 4),
			Min:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 3, 4}),
			Min:    2,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{1i, 2i, 3i, 4i}),
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
		Tensor *stale.Tensor
		Max    complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Max:    0,
		},
		"vector": {
			Tensor: stale.NewTensor(3).Fill([]complex128{-2, 4, 30}),
			Max:    30,
		},
		"zero max": {
			Tensor: stale.NewTensor(2, 3, 4),
			Max:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 3, 4}),
			Max:    8,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{1i, 4i, 3i, 2i}),
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
		Tensor *stale.Tensor
		Mean   complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Mean:   0,
		},
		"vector": {
			Tensor: stale.NewTensor(3).Fill([]complex128{-2, 4, 31}),
			Mean:   11,
		},
		"zero mean": {
			Tensor: stale.NewTensor(2, 3, 4),
			Mean:   0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 6, 4}),
			Mean:   5,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{2 + 1i, 1 - 4i, 4 + 3i, 9 + 2i}),
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

func TestTensorMedian(t *testing.T) {
	tests := map[string]struct {
		Tensor *stale.Tensor
		Median complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Median: 0,
		},
		"vector": {
			Tensor: stale.NewTensor(5).Fill([]complex128{-2, 4, 31, -1, 20}),
			Median: 4,
		},
		"zero median": {
			Tensor: stale.NewTensor(2, 3, 4),
			Median: 0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{8, 2, 5, 4}),
			Median: 4.5,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{2 + 1i, 1 - 4i, 4 + 3i, 9 + 2i}),
			Median: 3 + 2i,
		},
		"slice": {
			Tensor: tensorEnum(3, 2).Slice(1, 1),
			Median: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if median := test.Tensor.Median(); median != test.Median {
				t.Fatalf("want median=%v; got %v", test.Median, median)
			}
		})
	}
}

func TestTensorStd(t *testing.T) {
	tests := map[string]struct {
		Tensor *stale.Tensor
		Std    complex128
	}{
		"empty": {
			Tensor: stale.NewTensor(),
			Std:    cmplx.NaN(),
		},
		"vector": {
			Tensor: stale.NewTensor(5).Fill([]complex128{-2, 4, 30, -10, 34}),
			Std:    17.6,
		},
		"zero std": {
			Tensor: stale.NewTensor(2, 3, 4),
			Std:    0,
		},
		"matrix": {
			Tensor: tensorEnum(2, 2).Fill([]complex128{18, 7, 18, 15}),
			Std:    4.5,
		},
		"complex": {
			Tensor: stale.NewTensor(2, 2).Fill([]complex128{8 + 2i, 2 - 2i, 6 + 4i, 4i}),
			Std:    4,
		},
		"slice": {
			Tensor: tensorEnum(2, 2).Slice(1, 1),
			Std:    0.5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if std := test.Tensor.Std(); std != test.Std {
				t.Fatalf("want std=%v; got %v", test.Std, std)
			}
		})
	}
}
