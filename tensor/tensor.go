package tensor

import (
	"fmt"

	"github.com/ppknap/tacvs/internal/core"
	"github.com/ppknap/tacvs/internal/index"
)

type Tensor struct {
	idx *index.Index
	buf *core.Buffer
}

func (t *Tensor) At(pos ...int) *Tensor {
	if !t.idx.Validate(pos) {
		panic(fmt.Sprintf("tensor: invalid position %v for %v", pos, t.idx))
	}

	return &Tensor{
		idx: t.idx.Scalar(pos),
		buf: t.buf,
	}
}

func (t *Tensor) View() *Tensor {
	return nil // TODO
}

func Delegate(res, src *Tensor) *Tensor {
	return nil
}

func Add(a, b *Tensor) *Tensor {
	return nil
} // TODO
