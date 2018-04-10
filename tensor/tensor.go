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
	if t.idx != nil && t.idx.IsView() {
		return t
	}

	return &Tensor{
		idx: t.idx.View(),
		buf: t.buf,
	}
}

func (t *Tensor) Delegate() *Delegate {
	return NewDelegate(t)
}

type Delegate struct {
	dst *Tensor
}

func NewDelegate(dst *Tensor) *Delegate {
	return &Delegate{dst: dst}
}

func (d *Delegate) Add(a, b *Tensor) *Tensor {
	if a == nil || b == nil {
		panic("tensor: nil argument provided")
	}

	//	var dst = a
	//	math.Binary()
	return nil
}

// View
// Scheme
// Size(l) && Size(r) || Size(l) && Size(r) == 1
