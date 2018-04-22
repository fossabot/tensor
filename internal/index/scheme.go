package index

import "github.com/ppknap/tensor/internal/core"

// IdxScheme defines the order of dimmensions in continuos memory.
type IdxScheme Flags

const (
	// IdxSchemeColMajor represents column-major order where the vertical values
	// of a matrix lie side by side in continuos 1D array. This is also known as
	// Fortran order and is used in MatLab data indexing.
	IdxSchemeColMajor IdxScheme = 0
)

// Shape creates a shape array from provided strides.
func (s IdxScheme) Shape(strides []int) []int {
	if f, ok := schemeShapeFuncs[s]; ok {
		return f(strides)
	}

	panic(core.NewError("invalid strided indexing scheme"))
}

var schemeShapeFuncs = map[IdxScheme]func([]int) []int{
	IdxSchemeColMajor: colMajorShape,
}

func colMajorShape(strides []int) []int {
	if len(strides) == 0 {
		return nil
	}

	shape := []int{1}
	for i := 0; i < len(strides)-1; i++ {
		shape = append(shape, shape[len(shape)-1]*strides[i])
	}

	return shape
}

// Strides returns an array which contains data offsets on each dimension.
func (s IdxScheme) Strides(shape []int) []int {
	if f, ok := schemeStridesFuncs[s]; ok {
		return f(shape)
	}

	panic(core.NewError("invalid strided indexing scheme"))
}

var schemeStridesFuncs = map[IdxScheme]func([]int) []int{
	IdxSchemeColMajor: colMajorStrides,
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
	IdxSchemeColMajor: "column-major (F)",
}
