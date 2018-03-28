package index

const (
	IdxSchemeColMajor IdxScheme = iota
)

type IdxScheme uint

type idxSchemeFunc func([]int) []int

func (s IdxScheme) Strides(shape []int) []int {
	return nil
}

var schemeFuncs = map[IdxScheme]idxSchemeFunc{
	IdxSchemeColMajor: columnMajor,
}

func columnMajor(shape []int) []int {
	return nil
}

func (s IdxScheme) String() string {
	return ""
}

var schemeNames = map[IdxScheme]string{
	IdxSchemeColMajor: "bool",
}
