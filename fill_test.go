// Code generated by testgen; DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tensor"
	"github.com/ppknap/tensor/dtype"

	"github.com/ppknap/tensor/internal/core"
)

// Use reflect, core, and dtype packages in case they aren't used in tests.
var _ = reflect.TypeOf(dtype.DType(0) == core.DType(0))

func TestTensorZeros(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Zeros(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Zeros(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Zeros(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Zeros(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Zeros(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Zeros(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Zeros(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Zeros(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Zeros(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Zeros(),
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

func TestTensorOnes(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Ones(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Ones(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Ones(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Ones(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Ones(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Ones(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Ones(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Ones(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Ones(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Ones(),
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

func TestTensorFill(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Fill(tensor.NewScalar(5)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Fill(tensor.NewScalar(5)),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Fill(tensor.NewScalar(5)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Fill(tensor.NewScalar(5)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Fill(tensor.NewScalar(5)),
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

func TestTensorArange(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Arange(tensor.NewScalar(10), tensor.NewScalar(4)),
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

func TestTensorLinspace(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Linspace(tensor.NewScalar(10), tensor.NewScalar(20)),
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

func TestTensorEye(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Eye(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Eye(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Eye(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Eye(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Eye(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Eye(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Eye(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Eye(),
			Want: tensor.New(3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Eye(),
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

func TestTensorPanicEye(t *testing.T) {
	tests := map[string]func(){
		"three dim tensor": func() { _ = enumerate(tensor.New(4, 3, 2)).Eye() },
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if e, ok := recover().(*core.Error); !ok || e == nil {
					t.Fatalf("test should have panicked with Error, but it did not")
				}
			}()

			fn()
		})
	}
}