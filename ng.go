package ng

import "fmt"

type Tensor struct {
	data  []complex128
	shape []int
}

func NewTensor(first, second int, rest ...int) *Tensor {
	shape := append([]int{first, second}, rest...)

	size := 1
	for _, n := range shape {
		if n == 0 {
			panic(fmt.Sprintf("zero rank in shape: %v", shape))
		}

		size *= n
	}

	return &Tensor{
		data:  make([]complex128, size),
		shape: shape,
	}
}

func (t *Tensor) At(idx ...int) complex128 {
	if len(idx) != len(t.shape) {
		panic(fmt.Sprintf("invalid tensor index %v for shape %v", idx, t.shape))
	}

	var pos int
	for k := 0; k < len(t.shape); k++ {
		stride := 1
		for j := 0; j < k; j++ {
			stride *= t.shape[j]
		}

		pos += stride * idx[k]
	}

	return t.data[pos]
}

func (t *Tensor) Data() []complex128 {
	return t.data
}

func (t *Tensor) Slice(dim, from, to int) []*Tensor {
	return nil
}
