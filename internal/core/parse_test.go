package core

import "testing"

func TestStrAsComplex(t *testing.T) {
	tests := []struct {
		Str string
		Val complex128
	}{
		{Str: "", Val: 0},                        // 0 //
		{Str: "1i", Val: 1i},                     // 1 //
		{Str: "3j", Val: 3i},                     // 2 //
		{Str: "4e5", Val: 4e5},                   // 3 //
		{Str: "-4-4i", Val: -4 - 4i},             // 4 //
		{Str: "5e4-0.4e3j", Val: 5e4 - 0.4e3i},   // 5 //
		{Str: "-1.3e1+1.4i", Val: -1.3e1 + 1.4i}, // 6 //
		{Str: "-.2e1-2e3j", Val: -0.2e1 - 2e3i},  // 7 //
		{Str: "(-2.3-4.2i)", Val: -2.3 - 4.2i},   // 8 //
	}

	for i, test := range tests {
		val := strAsComplex(test.Str)

		if val != test.Val {
			t.Errorf("want val=%v; got %v (i:%d)", test.Val, val, i)
		}
	}
}
