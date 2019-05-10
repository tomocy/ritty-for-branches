package error

import "fmt"

type input interface {
	input() bool
}

func ValidationErrorf(format string, a ...interface{}) *ValidationError {
	return &ValidationError{
		msg: fmt.Sprintf(format, a...),
	}
}

type ValidationError struct {
	msg string
}

func (e *ValidationError) Error() string {
	return e.msg
}
