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
	if t.Size() == 0 {
		return "[]"
	}

	var maxElem = DefaultMaxFmtElements
	if t.FmtMaxElems > 1 {
		maxElem = t.FmtMaxElems
	}

	type info struct {
		group  []int
		isCont bool
		data   [][]complex128
	}

	group2data := make(map[int]*info)
	t.Apply(func(t *Tensor, idx []int) {
		key, group, midx := computeKeyIdx(maxElem, t.shape, idx)
		if key < 0 {
			return
		}

		nfo := group2data[key]
		if nfo == nil {
			data, isCont := makeMatrixBuf(maxElem, t.shape)
			nfo = &info{
				group:  group,
				isCont: isCont,
				data:   data,
			}

			group2data[key] = nfo
		}

		nfo.data[midx[0]][midx[1]] = t.At(idx...)
	})

	var (
		keys = make([]int, 0, len(group2data))
		fmtd = make([]string, 0, len(group2data))
	)

	for key, nfo := range group2data {
		keys = append(keys, key)

		fmtm := fmtMatrix(nfo.data)
		if nfo.isCont {
			fmtm = addCont(fmtm)
		}

		fmtd = append(fmtd, tabData(fmtm))
	}

	if len(fmtd) == 1 {
		return fmtd[0]
	}

	return ""
}

func makeMatrixBuf(maxElem int, shape []int) (ret [][]complex128, isCont bool) {
	var dims [2]int

	switch len(shape) {
	case 0:
		panic("tensor: matrix buffer from empty shape")
	case 1:
		dims[0] = min(shape[0], maxElem)
		dims[1] = 1
		isCont = shape[0] > dims[0]
	default:
		dims[0] = min(shape[0], maxElem)
		dims[1] = min(shape[1], maxElem)
		isCont = shape[0] > dims[0] || shape[1] > dims[1]
	}

	ret = make([][]complex128, dims[0])
	for i := range ret {
		ret[i] = make([]complex128, dims[1])
	}

	return ret, isCont
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func computeKeyIdx(maxElem int, size, idx []int) (key int, group, midx []int) {
	if len(size) != len(idx) {
		panic(fmt.Sprintf("invalid index length (%v!=%v)", size, idx))
	}

	for i := range size {
		if idx[i] > maxElem-2 && (idx[i] != size[i]-1) {
			return -1, nil, nil
		}
	}

	if len(idx) > 1 {
		group = idx[2:]
	}

	key = sliceToIdx(size, group)

	switch len(idx) {
	case 0:
		midx = []int{0, 0}
	case 1:
		midx = []int{idx[0], 0}
	default:
		midx = []int{idx[0], idx[1]}
	}

	for i := range midx {
		if midx[i] >= maxElem {
			midx[i] = maxElem - 1
		}
	}

	return key, group, midx
}

func sliceToIdx(size, group []int) int {
	if len(group) == 0 {
		return 0
	}

	val, mul := 0, 1
	for i := range group {
		val += group[i] * mul
		mul *= size[len(size)-len(group)+i]
	}

	return val
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

	return strings.TrimSpace(strings.Replace(ret[1:], "\n ", "\n", -1))
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
