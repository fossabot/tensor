// Code generated by testgen; DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tensor"
)

// Use reflect package in case it isn't used in tests.
var _ = reflect.TypeOf(tensor.DType(0))

func TestTensorEach(t *testing.T) {
	f := func() func([]int, *tensor.Tensor) {
		i := 1
		return func(_ []int, t *tensor.Tensor) {
			i++
			t.ItemSet(tensor.NewScalar(i))
		}
	}

	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Each(f()),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Each(f()),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Each(f()),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Each(f()),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Each(f()),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Each(f()),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Each(f()),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Each(f()),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Each(f()),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Each(f()),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.Want == nil && test.Got == nil {
				return
			}

			if test.Want == nil && test.Got != nil {
				t.Fatalf("want result to be nil, got %v", test.Got)
			}
			if test.Want != nil && test.Got == nil {
				t.Fatalf("want result to be not nil")
			}
			if err := checkTensor(test.Want, test.Got); err != nil {
				t.Errorf("want err=nil; got %v", err)
			}
		})
	}
}

func TestTensorItemAt(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1),
			Want: tensor.New(),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1),
			Want: tensor.New(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.Want == nil && test.Got == nil {
				return
			}

			if test.Want == nil && test.Got != nil {
				t.Fatalf("want result to be nil, got %v", test.Got)
			}
			if test.Want != nil && test.Got == nil {
				t.Fatalf("want result to be not nil")
			}
			if err := checkTensor(test.Want, test.Got); err != nil {
				t.Errorf("want err=nil; got %v", err)
			}
		})
	}
}

func TestTensorPanicItemAt(t *testing.T) {
	tests := map[string]func(){
		"zero value":                  func() { _ = (&tensor.Tensor{}).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"new empty tensor aka scalar": func() { _ = tensor.New().FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"vector with one element":     func() { _ = enumerate(tensor.New(1)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"vector with 9 elements":      func() { _ = enumerate(tensor.New(9)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"matrix one element":          func() { _ = enumerate(tensor.New(1, 1)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"matrix three rows two cols":  func() { _ = enumerate(tensor.New(3, 2)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"three dim tensor":            func() { _ = enumerate(tensor.New(4, 3, 2)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1) },
		"six dim tensor one element": func() {
			_ = enumerate(tensor.New(1, 1, 1, 1, 1, 1)).FillBuf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).ItemAt(1, 1)
		},
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if e, ok := recover().(*tensor.Error); !ok || e == nil {
					t.Fatalf("test should have panicked with Error, but it did not")
				}
			}()

			fn()
		})
	}
}

func TestTensorItemSet(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).ItemSet(tensor.NewScalar(5.0), 1, 1),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).ItemSet(tensor.NewScalar(5.0), 1, 1),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).ItemSet(tensor.NewScalar(5.0), 1, 1),
			Want: tensor.New(3, 2),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.Want == nil && test.Got == nil {
				return
			}

			if test.Want == nil && test.Got != nil {
				t.Fatalf("want result to be nil, got %v", test.Got)
			}
			if test.Want != nil && test.Got == nil {
				t.Fatalf("want result to be not nil")
			}
			if err := checkTensor(test.Want, test.Got); err != nil {
				t.Errorf("want err=nil; got %v", err)
			}
		})
	}
}

func TestTensorPanicItemSet(t *testing.T) {
	tests := map[string]func(){
		"zero value":                  func() { _ = (&tensor.Tensor{}).ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"new empty tensor aka scalar": func() { _ = tensor.New().ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"vector with one element":     func() { _ = enumerate(tensor.New(1)).ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"vector with 9 elements":      func() { _ = enumerate(tensor.New(9)).ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"matrix one element":          func() { _ = enumerate(tensor.New(1, 1)).ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"three dim tensor":            func() { _ = enumerate(tensor.New(4, 3, 2)).ItemSet(tensor.NewScalar(5.0), 1, 1) },
		"six dim tensor one element":  func() { _ = enumerate(tensor.New(1, 1, 1, 1, 1, 1)).ItemSet(tensor.NewScalar(5.0), 1, 1) },
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if e, ok := recover().(*tensor.Error); !ok || e == nil {
					t.Fatalf("test should have panicked with Error, but it did not")
				}
			}()

			fn()
		})
	}
}
