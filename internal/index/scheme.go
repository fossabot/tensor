package index

import (
	"github.com/ppknap/tensor/internal/errorc"
)

// IdxScheme defines the order of dimmensions in continuous memory.
type IdxScheme Flags

// DefaultIdxScheme defines the default indexing scheme when not set explicitly.
const DefaultIdxScheme = IdxSchemeRowMajor

const (
	// IdxSchemeRowMajor represents row major order where the horizontal values
	// of a matrix lie side by side in continuous 1D array. It is also known as
	// C element order.
	IdxSchemeRowMajor IdxScheme = 1 << iota
	// IdxSchemeColMajor represents column-major order where the vertical values
	// of a matrix lie side by side in continuous 1D array. This is also known as
	// Fortran order and is used in MatLab data indexing.
	IdxSchemeColMajor
)

// Shape creates a shape array from provided strides.
func (s IdxScheme) Shape(strides []int, size int) []int {
	if f, ok := schemeShapeFuncs[s]; ok {
		return f(strides, size)
	}

	panic(errorc.New("invalid strided indexing scheme"))
}

var schemeShapeFuncs = map[IdxScheme]func([]int, int) []int{
	IdxSchemeRowMajor: rowMajorShape,
	IdxSchemeColMajor: colMajorShape,
}

func rowMajorShape(strides []int, size int) []int {
	if len(strides) == 0 {
		return nil
	}

	var shape = []int{size / strides[0]}
	for i := 0; i < len(strides)-1; i++ {
		shape = append(shape, strides[i]/strides[i+1])
	}

	return shape
}

func colMajorShape(strides []int, size int) []int {
	if len(strides) == 0 {
		return nil
	}

	var shape []int
	for i := 1; i < len(strides); i++ {
		shape = append(shape, strides[i]/strides[i-1])
	}

	return append(shape, size/strides[len(strides)-1])
}

// Strides returns an array which contains data offsets on each dimension.
func (s IdxScheme) Strides(shape []int) []int {
	if f, ok := schemeStridesFuncs[s]; ok {
		return f(shape)
	}

	panic(errorc.New("invalid strided indexing scheme"))
}

var schemeStridesFuncs = map[IdxScheme]func([]int) []int{
	IdxSchemeRowMajor: rowMajorStrides,
	IdxSchemeColMajor: colMajorStrides,
}

func rowMajorStrides(shape []int) []int {
	if len(shape) == 0 {
		return nil
	}

	strides := make([]int, len(shape))
	for i := len(shape) - 1; i >= 0; i-- {
		if j := i + 1; j < len(shape) && shape[j] != 0 {
			strides[i] = strides[j] * shape[j]
		} else {
			strides[i] = 1
		}
	}

	return strides
}

func colMajorStrides(shape []int) []int {
	if len(shape) == 0 {
		return nil
	}

	strides := make([]int, len(shape))
	for i := range shape {
		if j := i - 1; j >= 0 && shape[j] != 0 {
			strides[i] = strides[j] * shape[j]
		} else {
			strides[i] = 1
		}
	}

	return strides
}

// String returns the name of index ordering scheme.
func (s IdxScheme) String() string {
	if name, ok := schemeNames[s]; ok {
		return name
	}

	return "unknown"
}

var schemeNames = map[IdxScheme]string{
	IdxSchemeRowMajor: "row-major (C)",
	IdxSchemeColMajor: "column-major (F)",
}
