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
