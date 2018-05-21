package index

import (
	"fmt"

	"github.com/ppknap/tensor/internal/errorc"
)

// Index represents an N-dimensional view on one dimensional array. It does not
// have any limit checks.
type Index struct {
	flags  Flags
	shape  []int
	stride []int
	offset int
}

// NewIndex creates a new Index instance. If scheme is 0, the column-major order
// scheme will be used.
func NewIndex(shape []int, scheme IdxScheme) *Index {
	if scheme == 0 {
		scheme = DefaultIdxScheme
	}

	// Shape with no length is treated as scalar.
	if len(shape) == 0 {
		shape = nil
	}

	// Negative axes are not allowed.
	for i := range shape {
		if shape[i] < 0 {
			panic(errorc.New("invalid shape %v", shape))
		}
	}

	return &Index{
		flags:  Flags(scheme),
		shape:  shape,
		stride: scheme.Strides(shape),
		offset: 0,
	}
}

// CopyNoView creates a copy of called index. It removes the index offset and
// view flag.
func (idx *Index) CopyNoView() *Index {
	return NewIndex(idx.shape, idx.flags.IdxScheme())
}

// NDim returns the number of dimmensions represented by index.
func (idx *Index) NDim() int {
	return len(idx.shape)
}

// Size returns the number of elements safely accessible by index.
func (idx *Index) Size() int {
	if len(idx.shape) == 0 {
		return 1
	}

	var size = 1
	for i := range idx.shape {
		size *= idx.shape[i]
	}

	return size
}

// At returns a function that computes array index for a given coordinates.
func (idx *Index) At() func([]int) int {
	if idx.Size() == 1 {
		return func([]int) int { return idx.offset }
	}

	return func(pos []int) (offset int) {
		for i := range pos {
			offset += pos[i] * idx.stride[i]
		}

		return idx.offset + offset
	}
}

// Validate checks if provided position is in index shape boundaries.
func (idx *Index) Validate(pos []int) bool {
	if idx == nil || len(pos) != len(idx.shape) {
		return false
	}

	for i := range idx.shape {
		if pos[i] < 0 || pos[i] >= idx.shape[i] {
			return false
		}
	}

	return true
}

// Strides returns offsets to step in each dimmension when traversing 1D array.
func (idx *Index) Strides() []int {
	return cloneInts(idx.stride)
}

// Shape returns an array representing dimmension sizes of the index.
func (idx *Index) Shape() []int {
	return cloneInts(idx.shape)
}

// Flags returns index properties.
func (idx *Index) Flags() Flags {
	return idx.flags
}

// EqShape returns true when shapes of provided indexes are identical.
func (idx *Index) EqShape(b *Index) bool {
	return EqShape(idx.shape, b.shape)
}

// MergeShape creates a shape that can fit shapes from both indexes.
func (idx *Index) MergeShape(b *Index) []int {
	if idx.shape == nil && b.shape == nil {
		return nil
	}

	long, short := idx.shape, b.shape
	if len(long) < len(short) {
		long, short = short, long
	}
	offset := len(long) - len(short)

	ms := make([]int, len(long))
	for i := len(long) - 1; i >= 0; i-- {
		if j := i - offset; j >= 0 && short[j] > long[i] {
			ms[i] = short[j]
		} else {
			ms[i] = long[i]
		}
	}

	return ms
}

// Iterate walks over N-dimensional index calling f with every possible indices.
// Slice given as function argument is reused by iterator logic and must not be
// modified. The last dimmension is iterated over first.
func (idx *Index) Iterate(f func(pos []int) bool) {
	// Scalars and zero slices support.
	if idx == nil {
		f(nil)
		return
	}

	if len(idx.shape) == 0 {
		f(idx.shape)
		return
	}

	pos, ok := make([]int, len(idx.shape)), true
	for i := 0; i >= 0 && ok; {
		switch {
		case pos[i] == idx.shape[i]:
			pos[i] = 0
			if i--; i >= 0 {
				pos[i]++
			}
		case i < len(pos)-1:
			i++
		case pos[i] < idx.shape[i]:
			ok = f(pos)
			pos[i]++
		}
	}
}

// Slice gets a subset of the index indices and creates a new Index instance
// which will compute a valid array index using new shape coordinates.
func (idx *Index) Slice(dim, from int, to ...int) *Index {
	limit := idx.shape[dim]
	if len(to) > 0 {
		limit = to[0]
	}

	shape := cloneInts(idx.shape)
	shape[dim] = limit - from

	return &Index{
		flags:  idx.flags.WithView(true),
		shape:  shape,
		stride: idx.stride,
		offset: idx.offset + from*idx.stride[dim],
	}
}

// View returns index which does not represent object which owns its data.
func (idx *Index) View() *Index {
	if idx == nil {
		return nil
	}

	return &Index{
		flags:  idx.flags.WithView(true),
		shape:  idx.shape,
		stride: idx.stride,
		offset: idx.offset,
	}
}

// Base creates a new index with original shape retrieved from strides.
func (idx *Index) Base() *Index {
	if idx == nil || !idx.flags.IsView() {
		return nil
	}

	return &Index{
		flags:  idx.flags.WithView(false),
		shape:  idx.flags.IdxScheme().Shape(idx.stride, idx.Size()),
		stride: idx.stride,
		offset: 0,
	}
}

// Scalar creates a 0-dimensional index representing a scalar object.
func (idx *Index) Scalar(pos []int) *Index {
	return &Index{
		flags:  idx.flags.WithView(true),
		shape:  nil,
		stride: nil,
		offset: idx.At()(pos),
	}
}

// String satisfies fmt.Stringer interface. It returns some basic info about
// index object.
func (idx *Index) String() string {
	return fmt.Sprintf("index %v, scheme %q", idx.shape, idx.flags.IdxScheme())
}

func cloneInts(slice []int) []int {
	if len(slice) == 0 {
		return nil
	}

	cp := make([]int, len(slice))
	copy(cp, slice)

	return cp
}
