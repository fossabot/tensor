package dtype

import (
	"github.com/ppknap/tensor/internal/core"
)

type DType = core.DType

const (
	Bool       = core.Bool
	Int        = core.Int
	Int16      = core.Int16
	Int64      = core.Int64
	Uint       = core.Uint
	Uint8      = core.Uint8
	Uint16     = core.Uint16
	Uint32     = core.Uint32
	Uint64     = core.Uint64
	Float32    = core.Float32
	Float64    = core.Float64
	Complex64  = core.Complex64
	Complex128 = core.Complex128
	String     = core.String
)
