package core

import (
	"fmt"
)

// Error satisfies buildin.error interface. All panics raised by tensor package
// should carry an instance of this type.
type Error string

// Error returns stored error message.
func (e Error) Error() string {
	return "tensor: " + string(e)
}

// Panic panics with an error formated from provided arguments.
func Panic(format string, a ...interface{}) {
	panic(Error(fmt.Sprintf(format, a...)))
}
