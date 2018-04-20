// Package test uses numpy commands to generate tests for tensor package.
package test

//go:generate go run codegen/conv.go codegen/data.go codegen/exec.go codegen/main.go codegen/test.go codegen/tmpl.go -instances instance.json -methods method/*.json -out ../
//go:generate gofmt -s -w ../
