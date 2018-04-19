package main

import (
	"fmt"
	"strconv"
	"strings"
)

var typeToExpr = map[string]func(string) string{
	"bool":           skipEmpty(boolToBool),
	"int":            skipEmpty(intToInt),
	"[]int":          skipEmpty(tupleToIntSlice),
	"*tensor.Tensor": skipEmpty(ndarrayToTensor),
}

// boolToBool checks and converts Python's boolean string to Go format.
func boolToBool(output string) string {
	if b, err := strconv.ParseBool(output); err == nil {
		return fmt.Sprintf("%t", b)
	}
	panic("invalid boolean: " + output)
}

// intToInt checks if provided integer string is valid.
func intToInt(output string) string {
	if _, err := strconv.Atoi(output); err == nil {
		return output
	}
	panic("invalid integer: " + output)
}

// tupleToIntSlice converts Python's tuple to Go int slice string.
func tupleToIntSlice(output string) string {
	toks := strings.Split(strings.Trim(output, "(,)"), ", ")
	if len(toks) == 1 && toks[0] == "" {
		return "nil"
	}

	for i := range toks {
		_ = intToInt(toks[i])
	}

	return "[]int{" + strings.Join(toks, ", ") + "}"
}

// ndarrayToTensor takes info about Python's ndarray and creates corresponding
// Tensor object expression.
func ndarrayToTensor(output string) string {
	return "tensor.New()"
}

// typeToExpr i a set of functions that format and validate given output to
// Go expression based on expected type.
func skipEmpty(f func(string) string) func(string) string {
	return func(output string) string {
		if output == "" {
			return ""
		}
		return f(output)
	}
}
