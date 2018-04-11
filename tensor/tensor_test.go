package tensor_test

import (
	"testing"

	"github.com/ppknap/tacvs/tensor"
)

func TestTensorNewPanic(t *testing.T) {
	tests := map[string]func(){
		"negative axis size": func() {
			_ = tensor.New(3, -2)
		},
	}

	for name, fn := range tests {
		t.Run(name, testPanic(fn))
	}
}

func testPanic(fn func()) func(*testing.T) {
	return func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("test should have panicked, but it did not")
			}
		}()

		fn()
	}
}
