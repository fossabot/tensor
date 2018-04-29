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
		{A: core.Int, B: core.Bool, Pr: core.Int},
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
		{A: core.Int8, B: core.Bool, Pr: core.Int8},
		{A: core.Int8, B: core.Int, Pr: core.Int},
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

		// Complex64.
		{A: core.Complex64, B: core.Bool, Pr: core.Complex64},
		{A: core.Complex64, B: core.Int, Pr: core.Complex128},
		{A: core.Complex64, B: core.Int8, Pr: core.Complex64},
		{A: core.Complex64, B: core.Int16, Pr: core.Complex64},
		{A: core.Complex64, B: core.Int32, Pr: core.Complex128},
		{A: core.Complex64, B: core.Int64, Pr: core.Complex128},
		{A: core.Complex64, B: core.Uint, Pr: core.Complex128},
		{A: core.Complex64, B: core.Uint8, Pr: core.Complex64},
		{A: core.Complex64, B: core.Uint16, Pr: core.Complex64},
		{A: core.Complex64, B: core.Uint32, Pr: core.Complex128},
		{A: core.Complex64, B: core.Uint64, Pr: core.Complex128},
		{A: core.Complex64, B: core.Uintptr, Pr: core.Complex128},
		{A: core.Complex64, B: core.Float32, Pr: core.Complex64},
		{A: core.Complex64, B: core.Float64, Pr: core.Complex128},
		{A: core.Complex64, B: core.Complex64, Pr: core.Complex64},
		{A: core.Complex64, B: core.Complex128, Pr: core.Complex128},
		{A: core.Complex64, B: core.String, Pr: core.String},
		// Complex128.
		{A: core.Complex128, B: core.Bool, Pr: core.Complex128},
		{A: core.Complex128, B: core.Int, Pr: core.Complex128},
		{A: core.Complex128, B: core.Int8, Pr: core.Complex128},
		{A: core.Complex128, B: core.Int16, Pr: core.Complex128},
		{A: core.Complex128, B: core.Int32, Pr: core.Complex128},
		{A: core.Complex128, B: core.Int64, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uint, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uint8, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uint16, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uint32, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uint64, Pr: core.Complex128},
		{A: core.Complex128, B: core.Uintptr, Pr: core.Complex128},
		{A: core.Complex128, B: core.Float32, Pr: core.Complex128},
		{A: core.Complex128, B: core.Float64, Pr: core.Complex128},
		{A: core.Complex128, B: core.Complex64, Pr: core.Complex128},
		{A: core.Complex128, B: core.Complex128, Pr: core.Complex128},
		{A: core.Complex128, B: core.String, Pr: core.String},
		// String.
		{A: core.String, B: core.Bool, Pr: core.String},
		{A: core.String, B: core.Int, Pr: core.String},
		{A: core.String, B: core.Int8, Pr: core.String},
		{A: core.String, B: core.Int16, Pr: core.String},
		{A: core.String, B: core.Int32, Pr: core.String},
		{A: core.String, B: core.Int64, Pr: core.String},
		{A: core.String, B: core.Uint, Pr: core.String},
		{A: core.String, B: core.Uint8, Pr: core.String},
		{A: core.String, B: core.Uint16, Pr: core.String},
		{A: core.String, B: core.Uint32, Pr: core.String},
		{A: core.String, B: core.Uint64, Pr: core.String},
		{A: core.String, B: core.Uintptr, Pr: core.String},
		{A: core.String, B: core.Float32, Pr: core.String},
		{A: core.String, B: core.Float64, Pr: core.String},
		{A: core.String, B: core.Complex64, Pr: core.String},
		{A: core.String, B: core.Complex128, Pr: core.String},
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
