package tacvs

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
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

func tabData(vss [][]string) string {
	if len(vss) == 0 {
		return "[]"
	}

	lineFn := func(data []string, l, r string) string {
		return l + "\t" + strings.Join(data, "\t") + "\t" + r
	}

	buf := &bytes.Buffer{}
	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.AlignRight)
	for i := range vss {
		switch {
		case len(vss) == 1:
			fmt.Fprintln(w, lineFn(vss[i], "[", "]"))
		case i == 0:
			fmt.Fprintln(w, lineFn(vss[i], "⎡", "⎤"))
		case i == len(vss)-1:
			fmt.Fprintln(w, lineFn(vss[i], "⎣", "⎦"))
		default:
			fmt.Fprintln(w, lineFn(vss[i], "⎢", "⎥"))
		}
	}

	w.Flush()
	ret := buf.String()

	return strings.Replace(ret[1:], "\n ", "\n", -1)
}

func addCont(vss [][]string) [][]string {
	const (
		ellipsis   = "⋯"
		elliSpaces = " ⋯ "
	)

	for i, vs := range vss {
		if len(vss[i]) > 1 {
			vss[i] = append(vs[:len(vs)-1], elliSpaces, vs[len(vs)-1])
		}
	}

	if len(vss) > 1 {
		ellVec := make([]string, len(vss[len(vss)-1]))
		for i := range ellVec {
			if i == len(ellVec)-2 {
				ellVec[i] = elliSpaces
			} else {
				ellVec[i] = ellipsis
			}
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
