package tensor

import (
	"github.com/ppknap/tacvs/internal/core"
	"github.com/ppknap/tacvs/internal/index"
)

type Tensor struct {
	idx *index.Index
	buf *core.Buffer
} // TODO

func (t *Tensor) At(idx ...int) *Tensor {
	return &Tensor{
		//	idx:
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
