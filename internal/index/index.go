package index

import (
	"fmt"
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
		scheme = IdxSchemeColMajor
	}

	return &Index{
		flags:  Flags(scheme),
		shape:  shape,
		stride: scheme.Strides(shape),
		offset: 0,
	}
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

// At returns array index for a given coordinates.
func (idx *Index) At(pos []int) (offset int) {
	for i := range pos {
		offset += pos[i] * idx.stride[i]
	}

	return idx.offset + offset
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

// Iterate walks over N-dimensional index calling f with every possible indices.
// Slice given as function argument is reused by iterator logic and must not be
// modified. The last dimmension is iterated over first.
//
// TODO: https://play.golang.org/p/cbQoG4zlNGG
func (idx *Index) Iterate(f func(pos []int)) {
	// Scalars and zero slices are not supported by indexing iterator.
	if len(idx.shape) == 0 || idx.Size() == 0 {
		return
	}

	pos := make([]int, len(idx.shape))
	for i := 0; i >= 0; {
		switch {
		case pos[i] == idx.shape[i]:
			pos[i] = 0
			if i--; i >= 0 {
				pos[i]++
			}
		case i < len(pos)-1:
			i++
		case pos[i] < idx.shape[i]:
			f(pos)
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

// Scalar creates a 0-dimensional index representing a scalar object.
func (idx *Index) Scalar(pos []int) *Index {
	return &Index{
		flags:  idx.flags.WithView(true),
		shape:  nil,
		stride: nil,
		offset: idx.At(pos),
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
