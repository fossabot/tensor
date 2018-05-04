package tensor

import (
	"github.com/ppknap/tensor/internal/core"
	"github.com/ppknap/tensor/internal/index"
)

// DType describes the internal byte representation of tensor's elements. One
// should not create instances of this type. The zero DType is not a valid type.
type DType = core.DType

const (
	Bool       = core.Bool       // bool type.
	Int        = core.Int        // int type.
	Int8       = core.Int8       // int8 type.
	Int16      = core.Int16      // int16 type.
	Int32      = core.Int32      // int32 type.
	Int64      = core.Int64      // int64 type.
	Uint       = core.Uint       // uint type.
	Uint8      = core.Uint8      // uint8 type.
	Uint16     = core.Uint16     // uint16 type.
	Uint32     = core.Uint32     // uint32 type.
	Uint64     = core.Uint64     // uint64 type.
	Uintptr    = core.Uintptr    // uintptr type.
	Float32    = core.Float32    // float32 type.
	Float64    = core.Float64    // float64 type.
	Complex64  = core.Complex64  // complex64 type.
	Complex128 = core.Complex128 // complex128 type.
	String     = core.String     // string type.
)

// Error satisfies error interface. If any invalid operation in tensor occurs,
// the object of this type will be panicked. Thus, one can revover it and
// examine what went wrong.
type Error = core.Error

// Tensor represents an organized multidimensional array of fixed-size elements.
type Tensor struct {
	idx *index.Index
	buf *core.Buffer
}

// New creates a new tensor with a given shape. Empty shape creates a scalar.
func New(shape ...int) *Tensor {
	var idx = index.NewIndex(shape, index.DefaultIdxScheme)

	return &Tensor{
		idx: idx,
		buf: core.NewBuffer(idx.Size()),
	}
}

// NewScalar creates a 0-dimensional tensor from a given value. The returned
// object's data type will be inherited from a given argument.
func NewScalar(scalar interface{}) *Tensor {
	t := &Tensor{
		idx: index.NewIndex(nil, index.DefaultIdxScheme),
		buf: core.NewBuffer(1),
	}

	typ, p := core.Destruct(scalar)
	t.buf.AsType(typ).Setptr()(0, typ, p)

	return t
}

func (t *Tensor) init() {
	if t.idx == nil && t.buf == nil {
		*t = *New()
	}
}
