package index

// Index represents an N-dimensional view on one dimensional array.
type Index struct {
	scheme IdxScheme
	shape  []int
	stride []int
	offset int
}

// NewIndex creates a new Index instance. If scheme is nil, the column-major
// order scheme will be used.
func NewIndex(shape []int, scheme IdxScheme) *Index {
	return nil
}

// NDim returns the number of dimmensions represented by index.
func (idx *Index) NDim() int { return 0 }

// Size returns the number of elements safely accessible by index.
func (idx *Index) Size() int { return 0 }

// At returns array index for a given coordinates.
func (idx *Index) At(pos []int) int {
	return 0
}

// Validate checks if provided position is in index shape boundaries.
func (idx *Index) Validate(pos []int) bool {
	return false
}

// Strides returns offsets to step in each dimmension when traversing 1D array.
func (idx *Index) Strides() []int { return nil } // copy

// Shape returns an array representing dimmension sizes of the index.
func (idx *Index) Shape() []int { return nil } // copy

// Slice gets a subset of the index indices and creates a new Index instance
// which will compute a valid array index using new shape coordinates.
func (idx *Index) Slice(dim, offset int, to ...int) *Index { return nil } // shalllow copy

// String satisfies fmt.Stringer interface. It returns some basic info about
// index object.
func (idx *Index) String() string { return "" }
