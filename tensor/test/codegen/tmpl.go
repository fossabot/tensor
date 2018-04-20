package main

import "text/template"

var tmpl = template.Must(template.New("").Funcs(funcMap).Parse(`// Code generated by codegen; DO NOT EDIT.

package tensor_test

import (
	"reflect"
	"testing"

	"github.com/ppknap/tacvs/tensor"
)

// Use reflect package in case it isn't used in tests.
var _ = reflect.TypeOf(&tensor.Tensor{}){{ range . }}{{ if .Pass }}

func TestTensor{{ .Name }}(t *testing.T) {
	tests := map[string]struct {
		Got, Want {{ .RTyp }}
	}{ {{- range .Pass }}
		"{{ .Name }}": {
			Got:  {{ .Expr }},
			Want: {{ .Want }},
		},{{ end }}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			{{ range $expr, $msg := makeComparators .RTyp }}if {{ $expr }} {
				t.Errorf({{ $msg }})
			}{{ end }}
		})
	}
}{{ end }}{{ if .Panic }}

func TestTensorPanic{{ .Name }}(t *testing.T) {
	tests := map[string]func() {
		{{- range .Panic }}
		"{{ .Name }}": func() { _ = {{ .Expr }} },{{ end }}
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			fn()
		})
	}
}{{ end }}{{ end }}
`))

var funcMap = template.FuncMap{
	// makeComparators creates a testing expressions based on tested type.
	"makeComparators": func(t string) map[string]string {
		switch t {
		case "bool", "byte", "int", "float64", "complex128", "string", "dtype.DType":
			return map[string]string{
				"test.Want != test.Got": `"want %v; got %v", test.Want, test.Got`,
			}
		case "*tensor.Tensor":
			return map[string]string{
				"ws, ts := test.Want.Shape(), test.Got.Shape(); reflect.DeepEqual(ws, ts)": `"want shape=%v; got %v", ws, ts`,
			}
		default:
			return map[string]string{
				"reflect.DeepEqual(test.Want, test.Got)": `"want %v; got %v", test.Want, test.Got`,
			}
		}
	},
}
