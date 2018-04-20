// Package test uses numpy commands to generate tests for tensor package.
package main

//go:generate go run conv.go data.go exec.go main.go test.go tmpl.go -instances data/instance.json -methods data/method/*.json -out ../
//go:generate gofmt -s -w ../
