package dtype

import (
	"fmt"
	"testing"
)

func TestPromote(t *testing.T) {
	tests := []struct {
		A, B, Pr DType
	}{
		// Bool.
		{A: Bool, B: Bool, Pr: Bool},
		{A: Bool, B: Int, Pr: Int},
		{A: Bool, B: Int8, Pr: Int8},
		{A: Bool, B: Int16, Pr: Int16},
		{A: Bool, B: Int32, Pr: Int32},
		{A: Bool, B: Int64, Pr: Int64},
		{A: Bool, B: Uint, Pr: Uint},
		{A: Bool, B: Uint8, Pr: Uint8},
		{A: Bool, B: Uint16, Pr: Uint16},
		{A: Bool, B: Uint32, Pr: Uint32},
		{A: Bool, B: Uint64, Pr: Uint64},
		{A: Bool, B: Uintptr, Pr: Uintptr},
		{A: Bool, B: Float32, Pr: Float32},
		{A: Bool, B: Float64, Pr: Float64},
		{A: Bool, B: Complex64, Pr: Complex64},
		{A: Bool, B: Complex128, Pr: Complex128},
		{A: Bool, B: String, Pr: String},
		// Int.
		{A: Int, B: Int, Pr: Int},
		{A: Int, B: Int8, Pr: Int},
		{A: Int, B: Int16, Pr: Int},
		{A: Int, B: Int32, Pr: Int},
		{A: Int, B: Int64, Pr: Int64},
		{A: Int, B: Uint, Pr: Float64},
		{A: Int, B: Uint8, Pr: Int},
		{A: Int, B: Uint16, Pr: Int},
		{A: Int, B: Uint32, Pr: Int64},
		{A: Int, B: Uint64, Pr: Float64},
		{A: Int, B: Uintptr, Pr: Float64},
		{A: Int, B: Float32, Pr: Float64},
		{A: Int, B: Float64, Pr: Float64},
		{A: Int, B: Complex64, Pr: Complex128},
		{A: Int, B: Complex128, Pr: Complex128},
		{A: Int, B: String, Pr: String},
		// Int8.
		{A: Int8, B: Int8, Pr: Int8},
		{A: Int8, B: Int16, Pr: Int16},
		{A: Int8, B: Int32, Pr: Int32},
		{A: Int8, B: Int64, Pr: Int64},
		{A: Int8, B: Uint, Pr: Float64},
		{A: Int8, B: Uint8, Pr: Int16},
		{A: Int8, B: Uint16, Pr: Int32},
		{A: Int8, B: Uint32, Pr: Int64},
		{A: Int8, B: Uint64, Pr: Float64},
		{A: Int8, B: Uintptr, Pr: Float64},
		{A: Int8, B: Float32, Pr: Float32},
		{A: Int8, B: Float64, Pr: Float64},
		{A: Int8, B: Complex64, Pr: Complex64},
		{A: Int8, B: Complex128, Pr: Complex128},
		{A: Int8, B: String, Pr: String},
		// Int16.
		{A: Int16, B: Int16, Pr: Int16},
		{A: Int16, B: Int32, Pr: Int32},
		{A: Int16, B: Int64, Pr: Int64},
		{A: Int16, B: Uint, Pr: Float64},
		{A: Int16, B: Uint8, Pr: Int16},
		{A: Int16, B: Uint16, Pr: Int32},
		{A: Int16, B: Uint32, Pr: Int64},
		{A: Int16, B: Uint64, Pr: Float64},
		{A: Int16, B: Uintptr, Pr: Float64},
		{A: Int16, B: Float32, Pr: Float32},
		{A: Int16, B: Float64, Pr: Float64},
		{A: Int16, B: Complex64, Pr: Complex64},
		{A: Int16, B: Complex128, Pr: Complex128},
		{A: Int16, B: String, Pr: String},
		// Int32.
		{A: Int32, B: Int32, Pr: Int32},
		{A: Int32, B: Int64, Pr: Int64},
		{A: Int32, B: Uint, Pr: Float64},
		{A: Int32, B: Uint8, Pr: Int32},
		{A: Int32, B: Uint16, Pr: Int32},
		{A: Int32, B: Uint32, Pr: Int64},
		{A: Int32, B: Uint64, Pr: Float64},
		{A: Int32, B: Uintptr, Pr: Float64},
		{A: Int32, B: Float32, Pr: Float64},
		{A: Int32, B: Float64, Pr: Float64},
		{A: Int32, B: Complex64, Pr: Complex128},
		{A: Int32, B: Complex128, Pr: Complex128},
		{A: Int32, B: String, Pr: String},
		// Int64.
		{A: Int64, B: Int64, Pr: Int64},
		{A: Int64, B: Uint, Pr: Float64},
		{A: Int64, B: Uint8, Pr: Int64},
		{A: Int64, B: Uint16, Pr: Int64},
		{A: Int64, B: Uint32, Pr: Int64},
		{A: Int64, B: Uint64, Pr: Float64},
		{A: Int64, B: Uintptr, Pr: Float64},
		{A: Int64, B: Float32, Pr: Float64},
		{A: Int64, B: Float64, Pr: Float64},
		{A: Int64, B: Complex64, Pr: Complex128},
		{A: Int64, B: Complex128, Pr: Complex128},
		{A: Int64, B: String, Pr: String},
		// Uint.
		{A: Uint, B: Uint, Pr: Uint},
		{A: Uint, B: Uint8, Pr: Uint},
		{A: Uint, B: Uint16, Pr: Uint},
		{A: Uint, B: Uint32, Pr: Uint},
		{A: Uint, B: Uint64, Pr: Uint64},
		{A: Uint, B: Uintptr, Pr: Uintptr},
		{A: Uint, B: Float32, Pr: Float64},
		{A: Uint, B: Float64, Pr: Float64},
		{A: Uint, B: Complex64, Pr: Complex128},
		{A: Uint, B: Complex128, Pr: Complex128},
		{A: Uint, B: String, Pr: String},
		// Uint8.
		{A: Uint8, B: Uint8, Pr: Uint8},
		{A: Uint8, B: Uint16, Pr: Uint16},
		{A: Uint8, B: Uint32, Pr: Uint32},
		{A: Uint8, B: Uint64, Pr: Uint64},
		{A: Uint8, B: Uintptr, Pr: Uintptr},
		{A: Uint8, B: Float32, Pr: Float32},
		{A: Uint8, B: Float64, Pr: Float64},
		{A: Uint8, B: Complex64, Pr: Complex64},
		{A: Uint8, B: Complex128, Pr: Complex128},
		{A: Uint8, B: String, Pr: String},
		// Uint16.
		{A: Uint16, B: Uint16, Pr: Uint16},
		{A: Uint16, B: Uint32, Pr: Uint32},
		{A: Uint16, B: Uint64, Pr: Uint64},
		{A: Uint16, B: Uintptr, Pr: Uintptr},
		{A: Uint16, B: Float32, Pr: Float32},
		{A: Uint16, B: Float64, Pr: Float64},
		{A: Uint16, B: Complex64, Pr: Complex64},
		{A: Uint16, B: Complex128, Pr: Complex128},
		{A: Uint16, B: String, Pr: String},
		// Uint32.
		{A: Uint32, B: Uint32, Pr: Uint32},
		{A: Uint32, B: Uint64, Pr: Uint64},
		{A: Uint32, B: Uintptr, Pr: Uintptr},
		{A: Uint32, B: Float32, Pr: Float64},
		{A: Uint32, B: Float64, Pr: Float64},
		{A: Uint32, B: Complex64, Pr: Complex128},
		{A: Uint32, B: Complex128, Pr: Complex128},
		{A: Uint32, B: String, Pr: String},
		// Uint64.
		{A: Uint64, B: Uint64, Pr: Uint64},
		{A: Uint64, B: Uintptr, Pr: Uint64},
		{A: Uint64, B: Float32, Pr: Float64},
		{A: Uint64, B: Float64, Pr: Float64},
		{A: Uint64, B: Complex64, Pr: Complex128},
		{A: Uint64, B: Complex128, Pr: Complex128},
		{A: Uint64, B: String, Pr: String},
		// Uintptr.
		{A: Uintptr, B: Uintptr, Pr: Uintptr},
		{A: Uintptr, B: Float32, Pr: Float64},
		{A: Uintptr, B: Float64, Pr: Float64},
		{A: Uintptr, B: Complex64, Pr: Complex128},
		{A: Uintptr, B: Complex128, Pr: Complex128},
		{A: Uintptr, B: String, Pr: String},
		// Float32.
		{A: Float32, B: Float32, Pr: Float32},
		{A: Float32, B: Float64, Pr: Float64},
		{A: Float32, B: Complex64, Pr: Complex64},
		{A: Float32, B: Complex128, Pr: Complex128},
		{A: Float32, B: String, Pr: String},
		// Float64.
		{A: Float64, B: Float64, Pr: Float64},
		{A: Float64, B: Complex64, Pr: Complex128},
		{A: Float64, B: Complex128, Pr: Complex128},
		{A: Float64, B: String, Pr: String},
		// Complex64.
		{A: Complex64, B: Complex64, Pr: Complex64},
		{A: Complex64, B: Complex128, Pr: Complex128},
		{A: Complex64, B: String, Pr: String},
		// Complex128.
		{A: Complex128, B: Complex128, Pr: Complex128},
		{A: Complex128, B: String, Pr: String},
		// String.
		{A: String, B: String, Pr: String},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s_%s", test.A, test.B), func(t *testing.T) {
			pr := Promote(test.A, test.B)

			if pr != test.Pr {
				t.Errorf("want promoted type=%v; got %v", test.Pr, pr)
			}
		})
	}
}

func TestPromoteCommutativity(t *testing.T) {
	ts := []DType{
		Bool,
		Int,
		Int8,
		Int16,
		Int32,
		Int64,
		Uint,
		Uint8,
		Uint16,
		Uint32,
		Uint64,
		Uintptr,
		Float32,
		Float64,
		Complex64,
		Complex128,
		String,
	}

	for i, a := range ts {
		for j := i + 1; j < len(ts); j++ {
			var b = ts[j]
			t.Run(fmt.Sprintf("%[1]v_%[2]v==%[2]v_%[1]v", a, b), func(t *testing.T) {
				prA, prB := Promote(a, b), Promote(b, a)
				if prA != prB {
					t.Errorf("want a(%v)=b(%v)", prA, prB)
				}
			})
		}
	}
}
