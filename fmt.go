package tensor

import (
	"fmt"
	"strconv"
)

func (t *Tensor) Format(f fmt.State, c rune) {

	fmt.Fprintf(f, "%s", stateFmtStr(f, c, "+-# 0"))
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
