package main

import (
	"encoding/json"
	"os"
)

// instance represents an object on which tests are called.
type instance struct {
	Name    string `json:"name"`    // Tensor name.
	Tensor  string `json:"tensor"`  // Go tensor initialization.
	NDArray string `json:"ndarray"` // Python's ndarray initialization.
}

func loadInstances(path string) (instances []*instance, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&instances); err != nil {
		return nil, err
	}

	return instances, nil
}

// method describes the call on a single instance.
type method struct {
	Name  string  `json:"name"`  // Method name.
	RTyp  string  `json:"rtyp"`  // Function return type.
	Calls []*call `json:"calls"` // Go to python method call.
}

// call describes corresponding Go and Python method calls.
type call struct {
	Dsc string `json:"dsc"` // Additional description.
	Go  string `json:"go"`  // Call in Go.
	Py  string `json:"py"`  // Call in Python.
}

func loadMethods(path string) (methods []*method, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&methods); err != nil {
		return nil, err
	}

	return methods, nil
}
