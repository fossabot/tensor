package main

import (
	"fmt"

	"github.com/ppknap/tacvs/tensor"
)

func main() {
	t := tensor.Tensor{}
	fmt.Println(t.At(3))
}
