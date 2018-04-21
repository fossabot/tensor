package main

import (
	"fmt"
	"strconv"
	"strings"
)

var typeToExpr = map[string]func(string) string{
	"bool":           skipEmpty(boolToBool),
	"byte":           skipEmpty(numToByte),
	"int":            skipEmpty(numToInt),
	"float64":        skipEmpty(numToFloat64),
	"complex128":     skipEmpty(numToCmplx),
	"string":         skipEmpty(nil),
	"[]int":          skipEmpty(tupleToIntSlice),
	"dtype.DType":    skipEmpty(dtypeToDType),
	"*tensor.Tensor": skipEmpty(ndarrayToTensor),
}

// boolToBool checks and converts Python's boolean string to Go format.
func boolToBool(output string) string {
	if b, err := strconv.ParseBool(output); err == nil {
		return fmt.Sprintf("%t", b)
	}
	panic("invalid boolean: " + output)
}

// numToByte checks if provided byte string is valid.
func numToByte(output string) string {
	if n, err := strconv.ParseUint(output, 10, 8); err == nil {
		return fmt.Sprintf("%v", n)
	}
	panic("invalid byte: " + output)
}

// numToInt checks if provided integer string is valid.
func numToInt(output string) string {
	if _, err := strconv.Atoi(output); err == nil {
		return output
	}
	panic("invalid integer: " + output)
}

// numToFloat64 checks if provided float string is valid.
func numToFloat64(output string) string {
	if f, err := strconv.ParseFloat(output, 64); err == nil {
		return fmt.Sprintf("%v", f)
	}
	panic("invalid float: " + output)
}

// numToCmplx converts Python's complex number to Go complex type string.
func numToCmplx(output string) string {
	return strings.TrimSpace(strings.Replace(output, "j", "i", -1))
}

// tupleToIntSlice converts Python's tuple to Go int slice string.
func tupleToIntSlice(output string) string {
	toks := strings.Split(strings.Trim(output, "(,)"), ", ")
	if len(toks) == 1 && toks[0] == "" {
		return "nil"
	}

	for i := range toks {
		_ = numToInt(toks[i])
	}

	return "[]int{" + strings.Join(toks, ", ") + "}"
}

// dtypeToDType takes numpy's dtype name and converts it to coresponding DType
// object expression.
func dtypeToDType(output string) string {
	typ := strings.Trim(output, "<>=")

	if strings.HasPrefix(typ, "U") {
		return "dtype.String"
	}

	return "dtype." + strings.Title(typ)
}

// ndarrayToTensor takes info about Python's ndarray and creates corresponding
// Tensor object expression.
func ndarrayToTensor(output string) string {
	output = strings.TrimSpace(output)
	if output == "nil" {
		return "nil"
	}

	toks := strings.FieldsFunc(output, func(r rune) bool { return r == '\n' })

	return fmt.Sprintf("tensor.New(%s)", strings.Trim(toks[0], "(,)"))
}

// typeToExpr i a set of functions that format and validate given output to
// Go expression based on expected type.
func skipEmpty(f func(string) string) func(string) string {
	return func(output string) string {
		if output == "" {
			return ""
		}

		if f == nil {
			return output
		}

		return f(output)
	}
}
