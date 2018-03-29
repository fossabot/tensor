package index

// IdxScheme defines the order of dimmensions in continuos memory.
type IdxScheme uint

const (
	// IdxSchemeColMajor represents column-major order where the vertical values
	// of a matrix lie side by side in continuos 1D array. This is also known as
	// Fortran order and is used in MatLab data indexing.
	IdxSchemeColMajor IdxScheme = iota
)

// Strides returns an array which contains data offsets on each dimension.
func (s IdxScheme) Strides(shape []int) []int {
	if f, ok := schemeFuncs[s]; ok {
		return f(shape)
	}

	panic("index: invalid strided indexing scheme")
}

var schemeFuncs = map[IdxScheme]func([]int) []int{
	IdxSchemeColMajor: colMajor,
}

func colMajor(shape []int) []int {
	return nil
}

// String returns the name of index ordering scheme.
func (s IdxScheme) String() string {
	if name, ok := schemeNames[s]; ok {
		return name
	}

	return "unknown"
}

var schemeNames = map[IdxScheme]string{
	IdxSchemeColMajor: "bool",
}
