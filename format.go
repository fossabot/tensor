package tacvs

import (
	"fmt"
	"strings"
)

// DefaultMaxFmtElements defines the maximum number of tensor's printable
// elements during formations.
const DefaultMaxFmtElements = 8

// String returns a printable representation of a tensor. If any shape size is
// greater than DefaultMaxFmtElements, excess elements will be represented as
// ellipsis symbol.
func (t *Tensor) String() string {
	return ""
}

func addCont(vss [][]string) [][]string {
	const ellipsis = "⋯"

	for i, vs := range vss {
		if len(vss[i]) > 1 {
			vss[i] = append(vs[:len(vs)-1], ellipsis, vs[len(vs)-1])
		}
	}

	if len(vss) > 1 {
		ellVec := make([]string, len(vss[len(vss)-1]))
		for i := range ellVec {
			ellVec[i] = ellipsis
		}

		vss = append(vss[:len(vss)-1], ellVec, vss[len(vss)-1])
	}

	return vss
}

func fmtMatrix(mat [][]complex128) (res [][]string) {
	for _, row := range mat {
		if len(row) == 0 {
			continue
		}

		rows := make([]string, 0, len(row))
		for _, val := range row {
			rows = append(rows, fmtCplx(val))
		}

		res = append(res, rows)
	}

	return res
}

func fmtCplx(val complex128) string {
	switch {
	case val == 0:
		return "0"
	case real(val) == 0:
		return fmt.Sprintf("%vi", imag(val))
	case imag(val) == 0:
		return fmt.Sprintf("%v", real(val))
	}

	return strings.Trim(fmt.Sprintf("%v", val), "()")
}
