package errorc

import (
	"fmt"
)

// Error satisfies buildin.error interface. It represents a common error type.
// All panics raised by tensor package should carry on an instance of this type.
type Error struct {
	msg string
}

// New creates an error with message formated from provided arguments.
func New(format string, a ...interface{}) *Error {
	return &Error{msg: fmt.Sprintf(format, a...)}
}

// Error returns stored error message.
func (e *Error) Error() string {
	return "tensor: " + string(e.msg)
}
