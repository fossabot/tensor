package tensor

import (
	"github.com/ppknap/tensor/internal/buffer"
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
)

// DType describes the internal byte representation of tensor's elements. One
// should not create instances of this type. The zero DType is not a valid type.
type DType = dtype.DType

const (
	Bool       = dtype.Bool       // bool type.
	Int        = dtype.Int        // int type.
	Int8       = dtype.Int8       // int8 type.
	Int16      = dtype.Int16      // int16 type.
	Int32      = dtype.Int32      // int32 type.
	Int64      = dtype.Int64      // int64 type.
	Uint       = dtype.Uint       // uint type.
	Uint8      = dtype.Uint8      // uint8 type.
	Uint16     = dtype.Uint16     // uint16 type.
	Uint32     = dtype.Uint32     // uint32 type.
	Uint64     = dtype.Uint64     // uint64 type.
	Uintptr    = dtype.Uintptr    // uintptr type.
	Float32    = dtype.Float32    // float32 type.
	Float64    = dtype.Float64    // float64 type.
	Complex64  = dtype.Complex64  // complex64 type.
	Complex128 = dtype.Complex128 // complex128 type.
	String     = dtype.String     // string type.
)

// Error satisfies error interface. If any invalid operation in tensor occurs,
// the object of this type will be panicked. Thus, one can revover it and
// examine what went wrong.
type Error = errorc.Error

// Tensor represents an organized multidimensional array of fixed-size elements.
type Tensor struct {
	idx *index.Index
	buf *buffer.Buffer
}

// New creates a new tensor with a given shape. Empty shape creates a scalar.
func New(shape ...int) *Tensor {
	var idx = index.NewIndex(shape, index.DefaultIdxScheme)

	return &Tensor{
		idx: idx,
		buf: buffer.New(idx.Size()),
	}
}

// NewScalar creates a 0-dimensional tensor from a given value. The returned
// object's data type will be inherited from a given argument.
func NewScalar(scalar interface{}) *Tensor {
	t := &Tensor{
		idx: index.NewIndex(nil, index.DefaultIdxScheme),
		buf: buffer.New(1),
	}

	typ, p := dtype.Destruct(scalar)
	t.buf.AsType(typ).Setptr()(0, typ, p)

	return t
}

func (t *Tensor) init() {
	if t.idx == nil && t.buf == nil {
		*t = *New()
	}
}
