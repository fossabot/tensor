package stale_test

import (
	"strings"
	"testing"

	"github.com/ppknap/tensor/stale"
)

func TestTensorString(t *testing.T) {
	tests := map[string]struct {
		Tensor      *stale.Tensor
		FmtMaxElems int
		String      string
	}{
		"empty tensor": {
			Tensor: stale.NewTensor(),
			String: "[]",
		},
		"one real element": {
			Tensor: stale.NewTensor(1).Fill([]complex128{
				2,
			}),
			String: "[ 2]",
		},
		"one imaginary element": {
			Tensor: stale.NewTensor(1).Fill([]complex128{
				2i,
			}),
			String: "[ 2i]",
		},
		"one complex element": {
			Tensor: stale.NewTensor(1).Fill([]complex128{
				1 + 2i,
			}),
			String: "[ 1+2i]",
		},
		"transposed vector mixed elements": {
			Tensor: stale.NewTensor(1, 5).Fill([]complex128{
				1 + 2i, 2i, 1 + 3i, -5i, 7,
			}),
			String: "[ 1+2i 2i 1+3i -5i 7]",
		},
		"vector mixed elements": {
			Tensor: stale.NewTensor(5, 1).Fill([]complex128{
				76, 1 + 2i, -21i, 7, 4.56,
			}),
			String: `⎡   76⎤
					 ⎢ 1+2i⎥
					 ⎢ -21i⎥
					 ⎢    7⎥
					 ⎣ 4.56⎦`,
		},
		"matrix mixed elements": {
			Tensor: stale.NewTensor(5, 4).Fill([]complex128{
				76, 1 + 2i, -21i, 7, 4.76,
				-3.2 - 2i, 9i, 1i, 90 + 3i, 0,
				5, -3 + 1i, -2 + 4i, 123, 8 + 3i,
				3i, 2.34 + 1i, 4i, 0.24, 2.2i,
			}),
			String: `⎡   76 -3.2-2i     5      3i⎤
					 ⎢ 1+2i      9i -3+1i 2.34+1i⎥
					 ⎢ -21i      1i -2+4i      4i⎥
					 ⎢    7   90+3i   123    0.24⎥
					 ⎣ 4.76       0  8+3i    2.2i⎦`,
		},
		"transposed vector mixed elements limit": {
			Tensor: stale.NewTensor(1, 5).Fill([]complex128{
				1 + 2i, 2i, 0, 0, 7,
			}),
			FmtMaxElems: 3,
			String:      "[ 1+2i 2i  ⋯  7]",
		},
		"vector mixed elements limit": {
			Tensor: stale.NewTensor(5).Fill([]complex128{
				76, 1 + 2i, 0, 0, 7,
			}),
			FmtMaxElems: 3,
			String: `⎡   76⎤
				     ⎢ 1+2i⎥
					 ⎢    ⋯⎥
					 ⎣    7⎦`,
		},
		"matrix mixed elements limit": {
			Tensor: stale.NewTensor(6, 5).Fill([]complex128{
				76, 1 + 2i, -21i, 0, 0, 4.76,
				-3.2 - 2i, 9i, 1i, 0, 0, 0,
				0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0,
				3i, 2.34 + 1i, 4i, 0, 0, 2.2i,
			}),
			FmtMaxElems: 3,
			String: `⎡   76 -3.2-2i  ⋯       3i⎤
					 ⎢ 1+2i      9i  ⋯  2.34+1i⎥
					 ⎢    ⋯       ⋯  ⋯        ⋯⎥
					 ⎣ 4.76       0  ⋯     2.2i⎦`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Tensor.FmtMaxElems = test.FmtMaxElems

			var (
				str  = test.Tensor.String()
				want = trimBlockSpace(test.String)
			)

			if str != want {
				t.Fatalf("want str=\n%q; got:\n%q", want, str)
			}
		})
	}
}

func trimBlockSpace(s string) string {
	fs := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })
	for i := range fs {
		fs[i] = strings.TrimSpace(fs[i])
	}

	return strings.Join(fs, "\n")
}
