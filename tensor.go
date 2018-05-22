package tensor

import (
	"reflect"

	"github.com/ppknap/tensor/internal/buffer"
	"github.com/ppknap/tensor/internal/dtype"
	"github.com/ppknap/tensor/internal/errorc"
	"github.com/ppknap/tensor/internal/index"
)

// DType describes the internal byte representation of tensor's elements. One
// should not create instances of this type. The zero DType is not a valid type.
type DType = dtype.DType

// Group of all supported data types.
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

// NewVector creates a new 1-dimensional tensor from a given array. The type and
// size of a newly created object will be inherited from a given argument.
func NewVector(slice interface{}) *Tensor {
	if slice == nil {
		panic(errorc.New("nil argument provided"))
	}

	v := reflect.Indirect(reflect.ValueOf(slice))
	if k := v.Kind(); k != reflect.Slice && k != reflect.Array {
		panic(errorc.New("argument type is not array-like (got %v)", k))
	}

	t := New(v.Len()).AsType(dtype.FromKind(v.Type().Elem().Kind()))
	t.Each(func(pos []int, _ *Tensor) {
		t.ItemSet(NewScalar(v.Index(pos[0]).Interface()), pos...)
	})

	return t
}

// NewMatrix creates a new 2-dimensional tensor from a given two dimensional
// array. Inner arrays will be treated as rows and their lengths must match.
// The type and shape of a newly created object will be inherited from a given
// argument.
func NewMatrix(slice2d interface{}) *Tensor {
	if slice2d == nil {
		panic(errorc.New("nil argument provided"))
	}

	v := reflect.Indirect(reflect.ValueOf(slice2d))
	if k := v.Kind(); k != reflect.Slice && k != reflect.Array {
		panic(errorc.New("argument type is not array-like (got %v)", k))
	}

	if k := v.Type().Elem().Kind(); k != reflect.Slice && k != reflect.Array {
		panic(errorc.New("inner type is not array-like (got %v)", k))
	}

	ins := make([]reflect.Value, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		var in = v.Index(i)

		if l := in.Len(); len(ins) > 0 && ins[0].Len() != l {
			panic(errorc.New("inner lengths do not match (got %d != %d)", ins[0].Len(), l))
		}

		ins = append(ins, in)
	}

	var innerLen int
	if len(ins) > 0 {
		innerLen = ins[0].Len()
	}

	t := New(v.Len(), innerLen).AsType(dtype.FromKind(v.Type().Elem().Elem().Kind()))
	t.Each(func(pos []int, _ *Tensor) {
		t.ItemSet(NewScalar(ins[pos[0]].Index(pos[1]).Interface()), pos...)
	})

	return t
}

// Delegate creates a new delegate with its destination set to a called tensor.
func (t *Tensor) Delegate() *Delegate {
	return NewDelegate(t)
}

func (t *Tensor) mustScalar(name string) {
	if t.Size() != 1 {
		panic(errorc.New("%v value must be a scalar (shape %v)", name, t.Shape()))
	}
}

func (t *Tensor) init() {
	if t.idx == nil && t.buf == nil {
		*t = *New()
	}
}

var defaultDelegate = NewDelegate(nil)

// Add adds elements from tensors 'a' and 'b' element-wise. A new tensor with
// the computed result will be returned. This function allows to use either
// tensors with scalars or tensors that have equal shapes.
func Add(a, b *Tensor) *Tensor { return defaultDelegate.Add(a, b) }

// Subtract substracts elements 'b' from tensor 'a' element-wise. A new tensor
// with the computed result will be returned. This function allows to use either
// tensors with scalars or tensors that have equal shapes.
func Subtract(a, b *Tensor) *Tensor { return defaultDelegate.Subtract(a, b) }

// Multiply multiplies elements from tensors 'a' and 'b' element-wise. A new
// tensor with the computed result will be returned. This function allows to use
// either tensors with scalars or tensors that have equal shapes.
func Multiply(a, b *Tensor) *Tensor { return defaultDelegate.Multiply(a, b) }

// Divide divides elements 'a' by elements from tensor 'b' element-wise. A new
// tensor with the computed result will be returned. This function allows to use
// either tensors with scalars or tensors that have equal shapes.
func Divide(a, b *Tensor) *Tensor { return defaultDelegate.Divide(a, b) }

// Mod computes remainders of elements 'a' divided by elements from tensor 'b'
// element-wise. A new tensor with the computed result will be returned. This
// function allows to use either tensors with scalars or tensors that have equal
// shapes.
func Mod(a, b *Tensor) *Tensor { return defaultDelegate.Mod(a, b) }
