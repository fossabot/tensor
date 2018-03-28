package index

type Index struct {
	scheme IdxScheme
	shape  []int
	stride []int
	offset int
}

func NewIndex(shape []int, scheme IdxScheme) *Index {
	return nil
}

func (idx *Index) NDim() int { return 0 }

func (idx *Index) Size() int { return 0 }

func (idx *Index) Strides() []int { return nil } // copy

func (idx *Index) Shape() []int { return nil } // copy

func (idx *Index) Slice(dim, offset int) *Index { return nil } // shalllow copy

func (idx *Index) String() string { return "" }
