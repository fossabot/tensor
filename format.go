package tacvs

// DefaultMaxFmtElements defines the maximum number of tensor's printable
// elements during formations.
const DefaultMaxFmtElements = 8

// String returns a printable representation of a tensor. If any shape size is
// greater than DefaultMaxFmtElements, excess elements will be represented as
// ellipsis symbol.
func (t *Tensor) String() string {
	return ""
}
