package tensor

import (
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

// Format satisfies fmt.Formatter interface. This method allows to examine
// tensor content as well as format underlying elements. There are a number of
// format strings one can use e.g.:
//
//  `%v`   - display up to 1000 tensor elements in each dimmension with their
//           default format flags,
//  `%3v`  - limit the number of displayed elements in each dimmension,
//  `%.5f` - display floating point elements with the precision set to 5 runes,
//  `%-v`  - elements with dim > 1 tensors will be left justified along the columns,
//  `%#X`  - alternative format for underlying elements,
//  `%#v`  - special "debug" format. It displays useful information about the
//           tensor. This includes tensor "head" with maximum 10 elements.
//
// Note that formatting flags '-' and ' '(space) are not passed to underlying
// element formatters. Also, the 'width' formatting value sets the number of
// displayed elements instead of the minimum number of runes to display.
func (t *Tensor) Format(f fmt.State, c rune) {
	t.init()

	if c == 'v' && f.Flag('#') {
		t.fmtDebugStr(f)
		return
	}

	// Normal: pass "+#0". "- " are ignored. handle "-"
	fmt.Fprintf(f, "%s", stateFmtStr(f, c, "+-# 0"))
}

func (t *Tensor) fmtDebugStr(w io.Writer) {
	type Debug struct {
		addr   uintptr
		size   int
		shape  []int
		dtype  string
		scheme string
		owner  bool
		head10 []interface{}
	}

	dbg := Debug{
		addr:   (uintptr)(unsafe.Pointer(t)),
		size:   t.Size(),
		shape:  t.Shape(),
		dtype:  t.DType().String(),
		scheme: t.idx.Flags().IdxScheme().String(),
		owner:  t.IsOwner(),
	}

	var (
		elements       = 10
		ifcFn          = t.buf.DType().AsInterfaceFunc()
		tBufAt, tIdxAt = t.buf.At(), t.idx.At()
	)

	t.idx.Iterate(func(pos []int) bool {
		dbg.head10 = append(dbg.head10, ifcFn(tBufAt(tIdxAt(pos))))

		elements--
		return elements > 0
	})

	fmt.Fprintf(w, "%#v", dbg)
}

// stateFmtStr reverses format state to its string representation. It ignores
// flags that are not in filter string.
func stateFmtStr(f fmt.State, c rune, filter string) string {
	runes := []rune{'%'}
	for _, r := range filter {
		if f.Flag(int(r)) {
			runes = append(runes, r)
		}
	}

	if w, ok := f.Width(); ok {
		runes = append(runes, []rune(strconv.Itoa(w))...)
	}

	if p, ok := f.Precision(); ok {
		runes = append(runes, '.')
		if p != 0 {
			runes = append(runes, []rune(strconv.Itoa(p))...)
		}
	}

	return string(append(runes, c))
}
