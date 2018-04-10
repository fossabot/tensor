package core

import (
	"fmt"
)

// Error satisfies buildin.error interface. All panics raised by tensor package
// should carry an instance of this type.
type Error struct {
	msg string
}

// NewError creates an error with message formated from provided arguments.
func NewError(format string, a ...interface{}) *Error {
	return &Error{msg: fmt.Sprintf(format, a...)}
}

// Error returns stored error message.
func (e *Error) Error() string {
	return "tensor: " + string(e.msg)
}
