// Code generated by testgen; DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tensor"
)

// Use reflect package in case it isn't used in tests.
var _ = reflect.TypeOf(tensor.DType(0))

func TestTensorExp(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Exp(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Exp(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Exp(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Exp(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Exp(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Exp(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Exp(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Exp(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Exp(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Exp(),
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

func TestTensorPow(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Pow(tensor.NewScalar(5)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Pow(tensor.NewScalar(5)),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Pow(tensor.NewScalar(5)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Pow(tensor.NewScalar(5)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Pow(tensor.NewScalar(5)),
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

func TestTensorSqrt(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Sqrt(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Sqrt(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Sqrt(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Sqrt(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Sqrt(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Sqrt(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Sqrt(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Sqrt(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Sqrt(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Sqrt(),
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

func TestTensorLog(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Log(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Log(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Log(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Log(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Log(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Log(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Log(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Log(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Log(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Log(),
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

func TestTensorLog10(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value": {
			Got:  (&tensor.Tensor{}).Log10(),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar": {
			Got:  tensor.New().Log10(),
			Want: tensor.New(),
		},
		"vector with one element": {
			Got:  enumerate(tensor.New(1)).Log10(),
			Want: tensor.New(1),
		},
		"vector with 9 elements": {
			Got:  enumerate(tensor.New(9)).Log10(),
			Want: tensor.New(9),
		},
		"matrix one element": {
			Got:  enumerate(tensor.New(1, 1)).Log10(),
			Want: tensor.New(1, 1),
		},
		"square matrix": {
			Got:  enumerate(tensor.New(3, 3)).Log10(),
			Want: tensor.New(3, 3),
		},
		"square matrix view": {
			Got:  enumerate(tensor.New(3, 3).View()).Log10(),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols": {
			Got:  enumerate(tensor.New(3, 2)).Log10(),
			Want: tensor.New(3, 2),
		},
		"three dim tensor": {
			Got:  enumerate(tensor.New(4, 3, 2)).Log10(),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element": {
			Got:  enumerate(tensor.New(1, 1, 1, 1, 1, 1)).Log10(),
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
