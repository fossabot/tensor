// Code generated by testgen; DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tensor"
)

// Use reflect package in case it isn't used in tests.
var _ = reflect.TypeOf(tensor.DType(0))

func TestTensorAdd(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Add((&tensor.Tensor{}), (&tensor.Tensor{})),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Add(tensor.New(), tensor.New()),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1)), enumerate(tensor.New(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(9)), enumerate(tensor.New(9))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1)), enumerate(tensor.New(1, 1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3)), enumerate(tensor.New(3, 3))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3).View()), enumerate(tensor.New(3, 3).View())),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 2)), enumerate(tensor.New(3, 2))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(4, 3, 2)), enumerate(tensor.New(4, 3, 2))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), enumerate(tensor.New(1, 1, 1, 1, 1, 1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Add((&tensor.Tensor{}), tensor.NewScalar(1)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Add(tensor.New(), tensor.NewScalar(1)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1)), tensor.NewScalar(1)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(9)), tensor.NewScalar(1)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1)), tensor.NewScalar(1)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3)), tensor.NewScalar(1)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(1)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 2)), tensor.NewScalar(1)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(1)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(1)),
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

func TestTensorSubtract(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Subtract((&tensor.Tensor{}), (&tensor.Tensor{})),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Subtract(tensor.New(), tensor.New()),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1)), enumerate(tensor.New(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(9)), enumerate(tensor.New(9))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1, 1)), enumerate(tensor.New(1, 1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 3)), enumerate(tensor.New(3, 3))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 3).View()), enumerate(tensor.New(3, 3).View())),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 2)), enumerate(tensor.New(3, 2))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(4, 3, 2)), enumerate(tensor.New(4, 3, 2))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), enumerate(tensor.New(1, 1, 1, 1, 1, 1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Subtract((&tensor.Tensor{}), tensor.NewScalar(1)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(tensor.New(), tensor.NewScalar(1)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1)), tensor.NewScalar(1)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(9)), tensor.NewScalar(1)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1, 1)), tensor.NewScalar(1)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 3)), tensor.NewScalar(1)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(1)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(3, 2)), tensor.NewScalar(1)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(1)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Subtract(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(1)),
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

func TestTensorMultiply(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Multiply((&tensor.Tensor{}), (&tensor.Tensor{})),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Multiply(tensor.New(), tensor.New()),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1)), enumerate(tensor.New(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(9)), enumerate(tensor.New(9))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1, 1)), enumerate(tensor.New(1, 1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 3)), enumerate(tensor.New(3, 3))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 3).View()), enumerate(tensor.New(3, 3).View())),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 2)), enumerate(tensor.New(3, 2))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(4, 3, 2)), enumerate(tensor.New(4, 3, 2))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), enumerate(tensor.New(1, 1, 1, 1, 1, 1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Multiply((&tensor.Tensor{}), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(tensor.New(), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1)), tensor.NewScalar(2)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(9)), tensor.NewScalar(2)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1, 1)), tensor.NewScalar(2)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 3)), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Multiply(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(2)),
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

func TestTensorDivide(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Divide((&tensor.Tensor{}), (&tensor.Tensor{})),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Divide(tensor.New(), tensor.New()),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1)), enumerate(tensor.New(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(9)), enumerate(tensor.New(9))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1, 1)), enumerate(tensor.New(1, 1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 3)), enumerate(tensor.New(3, 3))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 3).View()), enumerate(tensor.New(3, 3).View())),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 2)), enumerate(tensor.New(3, 2))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(4, 3, 2)), enumerate(tensor.New(4, 3, 2))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), enumerate(tensor.New(1, 1, 1, 1, 1, 1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Divide((&tensor.Tensor{}), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Divide(tensor.New(), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1)), tensor.NewScalar(2)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(9)), tensor.NewScalar(2)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1, 1)), tensor.NewScalar(2)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 3)), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Divide(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(2)),
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

func TestTensorMaximum(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Maximum((&tensor.Tensor{}), tensor.NewDelegate(nil).Add((&tensor.Tensor{}), tensor.NewScalar(1))),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Maximum(tensor.New(), tensor.NewDelegate(nil).Add(tensor.New(), tensor.NewScalar(1))),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1)), tensor.NewScalar(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(9)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(9)), tensor.NewScalar(1))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1, 1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1)), tensor.NewScalar(1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 3)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3)), tensor.NewScalar(1))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 3).View()), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(1))),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 2)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 2)), tensor.NewScalar(1))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(4, 3, 2)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(1))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Maximum((&tensor.Tensor{}), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(tensor.New(), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1)), tensor.NewScalar(2)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(9)), tensor.NewScalar(2)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1, 1)), tensor.NewScalar(2)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 3)), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Maximum(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(2)),
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

func TestTensorMinimum(t *testing.T) {
	tests := map[string]struct {
		Got, Want *tensor.Tensor
	}{
		"zero value same": {
			Got:  tensor.NewDelegate(nil).Minimum((&tensor.Tensor{}), tensor.NewDelegate(nil).Add((&tensor.Tensor{}), tensor.NewScalar(1))),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar same": {
			Got:  tensor.NewDelegate(nil).Minimum(tensor.New(), tensor.NewDelegate(nil).Add(tensor.New(), tensor.NewScalar(1))),
			Want: tensor.New(),
		},
		"vector with one element same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1)), tensor.NewScalar(1))),
			Want: tensor.New(1),
		},
		"vector with 9 elements same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(9)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(9)), tensor.NewScalar(1))),
			Want: tensor.New(9),
		},
		"matrix one element same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1, 1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1)), tensor.NewScalar(1))),
			Want: tensor.New(1, 1),
		},
		"square matrix same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 3)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3)), tensor.NewScalar(1))),
			Want: tensor.New(3, 3),
		},
		"square matrix view same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 3).View()), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(1))),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 2)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(3, 2)), tensor.NewScalar(1))),
			Want: tensor.New(3, 2),
		},
		"three dim tensor same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(4, 3, 2)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(1))),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element same": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewDelegate(nil).Add(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(1))),
			Want: tensor.New(1, 1, 1, 1, 1, 1),
		},
		"zero value scalar": {
			Got:  tensor.NewDelegate(nil).Minimum((&tensor.Tensor{}), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"new empty tensor aka scalar scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(tensor.New(), tensor.NewScalar(2)),
			Want: tensor.New(),
		},
		"vector with one element scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1)), tensor.NewScalar(2)),
			Want: tensor.New(1),
		},
		"vector with 9 elements scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(9)), tensor.NewScalar(2)),
			Want: tensor.New(9),
		},
		"matrix one element scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1, 1)), tensor.NewScalar(2)),
			Want: tensor.New(1, 1),
		},
		"square matrix scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 3)), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"square matrix view scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 3).View()), tensor.NewScalar(2)),
			Want: tensor.New(3, 3),
		},
		"matrix three rows two cols scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(3, 2),
		},
		"three dim tensor scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(4, 3, 2)), tensor.NewScalar(2)),
			Want: tensor.New(4, 3, 2),
		},
		"six dim tensor one element scalar": {
			Got:  tensor.NewDelegate(nil).Minimum(enumerate(tensor.New(1, 1, 1, 1, 1, 1)), tensor.NewScalar(2)),
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