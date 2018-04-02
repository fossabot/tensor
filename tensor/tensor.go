package tensor

import (
	"fmt"
	"unsafe"

	"github.com/ppknap/tacvs/dtype"
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

func Delegate(res, src *Tensor) *Tensor {
	return nil
}

func Add(a, b *Tensor) *Tensor {
	return nil
}

func add(dt dtype.DType, ptrDst, ptrLhv, ptrRhv unsafe.Pointer) {
	switch dt {
	case dtype.Bool:
		*(*bool)(ptrDst) = *(*bool)(ptrLhv) || *(*bool)(ptrRhv)
	case dtype.Int:
		*(*int)(ptrDst) = *(*int)(ptrLhv) + *(*int)(ptrRhv)
	case dtype.Int64:
		*(*int64)(ptrDst) = *(*int64)(ptrLhv) + *(*int64)(ptrRhv)
	case dtype.String:
		*(*string)(ptrDst) = *(*string)(ptrLhv) + *(*string)(ptrRhv)
	default:
		panic("core: unsupported type: " + dt.String())
	}
}
