package main

import (
	"fmt"

	"github.com/ppknap/tensor"
)

func main() {
	fmt.Printf("Type %T\n", tensor.NewScalar(5))
	fmt.Printf("Float vec: %v\n---\n", tensor.NewVector([]float64{5, 3, 56}))
	fmt.Printf("Float vec dbg on: % #v\n---\n", tensor.NewVector([]float64{5, 3, 56}))
	fmt.Printf("Float vec dbg: %#v\n---\n", tensor.NewVector([]float64{5, 3, 56}))
	fmt.Printf("Float vec one elem: %#+1v\n---\n", tensor.NewVector([]float64{5, 3, 56, 4}))
	fmt.Printf("Float vec precission: %#1.5v\n---\n", tensor.NewVector([]float64{5, 3, 56}))
	fmt.Printf("Cmpx mat: %#v\n---\n", tensor.New(3, 5).Fill(tensor.NewScalar(4+5i)))
	fmt.Printf("Scalar:  %#v\n---\n", tensor.NewScalar(4+5i))
}
