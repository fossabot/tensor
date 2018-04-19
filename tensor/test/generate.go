// Package test uses numpy commands to generate tests for tensor package.
package test

//go:generate go run -tags codegen ../../private/model/cli/gen-endpoints/main.go -model ./endpoints.json -out ../../aws/endpoints/defaults.go
//go:generate gofmt -s -w ../
