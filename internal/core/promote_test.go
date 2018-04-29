package core_test

import (
	"fmt"
	"testing"

	"github.com/ppknap/tensor/internal/core"
)

func TestPromote(t *testing.T) {
	tests := []struct {
		A, B, Pr core.DType
	}{
		// Bool.
		{A: core.Bool, B: core.Bool, Pr: core.Bool},
		{A: core.Bool, B: core.Int, Pr: core.Int},
		{A: core.Bool, B: core.Int8, Pr: core.Int8},
		{A: core.Bool, B: core.Int16, Pr: core.Int16},
		{A: core.Bool, B: core.Int32, Pr: core.Int32},
		{A: core.Bool, B: core.Int64, Pr: core.Int64},
		{A: core.Bool, B: core.Uint, Pr: core.Uint},
		{A: core.Bool, B: core.Uint8, Pr: core.Uint8},
		{A: core.Bool, B: core.Uint16, Pr: core.Uint16},
		{A: core.Bool, B: core.Uint32, Pr: core.Uint32},
		{A: core.Bool, B: core.Uint64, Pr: core.Uint64},
		{A: core.Bool, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Bool, B: core.Float32, Pr: core.Float32},
		{A: core.Bool, B: core.Float64, Pr: core.Float64},
		{A: core.Bool, B: core.Complex64, Pr: core.Complex64},
		{A: core.Bool, B: core.Complex128, Pr: core.Complex128},
		{A: core.Bool, B: core.String, Pr: core.String},
		// Int.
		{A: core.Int, B: core.Int, Pr: core.Int},
		{A: core.Int, B: core.Int8, Pr: core.Int},
		{A: core.Int, B: core.Int16, Pr: core.Int},
		{A: core.Int, B: core.Int32, Pr: core.Int},
		{A: core.Int, B: core.Int64, Pr: core.Int64},
		{A: core.Int, B: core.Uint, Pr: core.Float64},
		{A: core.Int, B: core.Uint8, Pr: core.Int},
		{A: core.Int, B: core.Uint16, Pr: core.Int},
		{A: core.Int, B: core.Uint32, Pr: core.Int64},
		{A: core.Int, B: core.Uint64, Pr: core.Float64},
		{A: core.Int, B: core.Uintptr, Pr: core.Float64},
		{A: core.Int, B: core.Float32, Pr: core.Float64},
		{A: core.Int, B: core.Float64, Pr: core.Float64},
		{A: core.Int, B: core.Complex64, Pr: core.Complex128},
		{A: core.Int, B: core.Complex128, Pr: core.Complex128},
		{A: core.Int, B: core.String, Pr: core.String},
		// Int8.
		{A: core.Int8, B: core.Int8, Pr: core.Int8},
		{A: core.Int8, B: core.Int16, Pr: core.Int16},
		{A: core.Int8, B: core.Int32, Pr: core.Int32},
		{A: core.Int8, B: core.Int64, Pr: core.Int64},
		{A: core.Int8, B: core.Uint, Pr: core.Float64},
		{A: core.Int8, B: core.Uint8, Pr: core.Int16},
		{A: core.Int8, B: core.Uint16, Pr: core.Int32},
		{A: core.Int8, B: core.Uint32, Pr: core.Int64},
		{A: core.Int8, B: core.Uint64, Pr: core.Float64},
		{A: core.Int8, B: core.Uintptr, Pr: core.Float64},
		{A: core.Int8, B: core.Float32, Pr: core.Float32},
		{A: core.Int8, B: core.Float64, Pr: core.Float64},
		{A: core.Int8, B: core.Complex64, Pr: core.Complex64},
		{A: core.Int8, B: core.Complex128, Pr: core.Complex128},
		{A: core.Int8, B: core.String, Pr: core.String},
		// Int16.
		{A: core.Int16, B: core.Int16, Pr: core.Int16},
		{A: core.Int16, B: core.Int32, Pr: core.Int32},
		{A: core.Int16, B: core.Int64, Pr: core.Int64},
		{A: core.Int16, B: core.Uint, Pr: core.Float64},
		{A: core.Int16, B: core.Uint8, Pr: core.Int16},
		{A: core.Int16, B: core.Uint16, Pr: core.Int32},
		{A: core.Int16, B: core.Uint32, Pr: core.Int64},
		{A: core.Int16, B: core.Uint64, Pr: core.Float64},
		{A: core.Int16, B: core.Uintptr, Pr: core.Float64},
		{A: core.Int16, B: core.Float32, Pr: core.Float32},
		{A: core.Int16, B: core.Float64, Pr: core.Float64},
		{A: core.Int16, B: core.Complex64, Pr: core.Complex64},
		{A: core.Int16, B: core.Complex128, Pr: core.Complex128},
		{A: core.Int16, B: core.String, Pr: core.String},
		// Int32.
		{A: core.Int32, B: core.Int32, Pr: core.Int32},
		{A: core.Int32, B: core.Int64, Pr: core.Int64},
		{A: core.Int32, B: core.Uint, Pr: core.Float64},
		{A: core.Int32, B: core.Uint8, Pr: core.Int32},
		{A: core.Int32, B: core.Uint16, Pr: core.Int32},
		{A: core.Int32, B: core.Uint32, Pr: core.Int64},
		{A: core.Int32, B: core.Uint64, Pr: core.Float64},
		{A: core.Int32, B: core.Uintptr, Pr: core.Float64},
		{A: core.Int32, B: core.Float32, Pr: core.Float64},
		{A: core.Int32, B: core.Float64, Pr: core.Float64},
		{A: core.Int32, B: core.Complex64, Pr: core.Complex128},
		{A: core.Int32, B: core.Complex128, Pr: core.Complex128},
		{A: core.Int32, B: core.String, Pr: core.String},
		// Int64.
		{A: core.Int64, B: core.Int64, Pr: core.Int64},
		{A: core.Int64, B: core.Uint, Pr: core.Float64},
		{A: core.Int64, B: core.Uint8, Pr: core.Int64},
		{A: core.Int64, B: core.Uint16, Pr: core.Int64},
		{A: core.Int64, B: core.Uint32, Pr: core.Int64},
		{A: core.Int64, B: core.Uint64, Pr: core.Float64},
		{A: core.Int64, B: core.Uintptr, Pr: core.Float64},
		{A: core.Int64, B: core.Float32, Pr: core.Float64},
		{A: core.Int64, B: core.Float64, Pr: core.Float64},
		{A: core.Int64, B: core.Complex64, Pr: core.Complex128},
		{A: core.Int64, B: core.Complex128, Pr: core.Complex128},
		{A: core.Int64, B: core.String, Pr: core.String},
		// Uint.
		{A: core.Uint, B: core.Uint, Pr: core.Uint},
		{A: core.Uint, B: core.Uint8, Pr: core.Uint},
		{A: core.Uint, B: core.Uint16, Pr: core.Uint},
		{A: core.Uint, B: core.Uint32, Pr: core.Uint},
		{A: core.Uint, B: core.Uint64, Pr: core.Uint64},
		{A: core.Uint, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Uint, B: core.Float32, Pr: core.Float64},
		{A: core.Uint, B: core.Float64, Pr: core.Float64},
		{A: core.Uint, B: core.Complex64, Pr: core.Complex128},
		{A: core.Uint, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uint, B: core.String, Pr: core.String},
		// Uint8.
		{A: core.Uint8, B: core.Uint8, Pr: core.Uint8},
		{A: core.Uint8, B: core.Uint16, Pr: core.Uint16},
		{A: core.Uint8, B: core.Uint32, Pr: core.Uint32},
		{A: core.Uint8, B: core.Uint64, Pr: core.Uint64},
		{A: core.Uint8, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Uint8, B: core.Float32, Pr: core.Float32},
		{A: core.Uint8, B: core.Float64, Pr: core.Float64},
		{A: core.Uint8, B: core.Complex64, Pr: core.Complex64},
		{A: core.Uint8, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uint8, B: core.String, Pr: core.String},
		// Uint16.
		{A: core.Uint16, B: core.Uint16, Pr: core.Uint16},
		{A: core.Uint16, B: core.Uint32, Pr: core.Uint32},
		{A: core.Uint16, B: core.Uint64, Pr: core.Uint64},
		{A: core.Uint16, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Uint16, B: core.Float32, Pr: core.Float32},
		{A: core.Uint16, B: core.Float64, Pr: core.Float64},
		{A: core.Uint16, B: core.Complex64, Pr: core.Complex64},
		{A: core.Uint16, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uint16, B: core.String, Pr: core.String},
		// Uint32.
		{A: core.Uint32, B: core.Uint32, Pr: core.Uint32},
		{A: core.Uint32, B: core.Uint64, Pr: core.Uint64},
		{A: core.Uint32, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Uint32, B: core.Float32, Pr: core.Float64},
		{A: core.Uint32, B: core.Float64, Pr: core.Float64},
		{A: core.Uint32, B: core.Complex64, Pr: core.Complex128},
		{A: core.Uint32, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uint32, B: core.String, Pr: core.String},
		// Uint64.
		{A: core.Uint64, B: core.Uint64, Pr: core.Uint64},
		{A: core.Uint64, B: core.Uintptr, Pr: core.Uint64},
		{A: core.Uint64, B: core.Float32, Pr: core.Float64},
		{A: core.Uint64, B: core.Float64, Pr: core.Float64},
		{A: core.Uint64, B: core.Complex64, Pr: core.Complex128},
		{A: core.Uint64, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uint64, B: core.String, Pr: core.String},
		// Uintptr.
		{A: core.Uintptr, B: core.Uintptr, Pr: core.Uintptr},
		{A: core.Uintptr, B: core.Float32, Pr: core.Float64},
		{A: core.Uintptr, B: core.Float64, Pr: core.Float64},
		{A: core.Uintptr, B: core.Complex64, Pr: core.Complex128},
		{A: core.Uintptr, B: core.Complex128, Pr: core.Complex128},
		{A: core.Uintptr, B: core.String, Pr: core.String},
		// Float32.
		{A: core.Float32, B: core.Float32, Pr: core.Float32},
		{A: core.Float32, B: core.Float64, Pr: core.Float64},
		{A: core.Float32, B: core.Complex64, Pr: core.Complex64},
		{A: core.Float32, B: core.Complex128, Pr: core.Complex128},
		{A: core.Float32, B: core.String, Pr: core.String},
		// Float64.
		{A: core.Float64, B: core.Float64, Pr: core.Float64},
		{A: core.Float64, B: core.Complex64, Pr: core.Complex128},
		{A: core.Float64, B: core.Complex128, Pr: core.Complex128},
		{A: core.Float64, B: core.String, Pr: core.String},
		// Complex64.
		{A: core.Complex64, B: core.Complex64, Pr: core.Complex64},
		{A: core.Complex64, B: core.Complex128, Pr: core.Complex128},
		{A: core.Complex64, B: core.String, Pr: core.String},
		// Complex128.
		{A: core.Complex128, B: core.Complex128, Pr: core.Complex128},
		{A: core.Complex128, B: core.String, Pr: core.String},
		// String.
		{A: core.String, B: core.String, Pr: core.String},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s_%s", test.A, test.B), func(t *testing.T) {
			pr := core.Promote(test.A, test.B)

			if pr != test.Pr {
				t.Errorf("want promoted type=%v; got %v", test.Pr, pr)
			}
		})
	}
}

func TestPromoteCommutativity(t *testing.T) {
	ts := []core.DType{
		core.Bool,
		core.Int,
		core.Int8,
		core.Int16,
		core.Int32,
		core.Int64,
		core.Uint,
		core.Uint8,
		core.Uint16,
		core.Uint32,
		core.Uint64,
		core.Uintptr,
		core.Float32,
		core.Float64,
		core.Complex64,
		core.Complex128,
		core.String,
	}

	for i, a := range ts {
		for j := i + 1; j < len(ts); j++ {
			var b = ts[j]
			t.Run(fmt.Sprintf("%[1]v_%[2]v==%[2]v_%[1]v", a, b), func(t *testing.T) {
				prA, prB := core.Promote(a, b), core.Promote(b, a)
				if prA != prB {
					t.Errorf("want a(%v)=b(%v)", prA, prB)
				}
			})
		}
	}
}
