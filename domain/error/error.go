package error

import "fmt"

func InInput(err error) bool {
	input, ok := err.(input)
	return ok && input.input()
}

type input interface {
	input() bool
}

func InInternal(err error) bool {
	internal, ok := err.(internal)
	return ok && internal.internal()
}

type internal interface {
	internal() bool
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

func newBareError(status status, msg string) *bareError {
	return &bareError{
		status: status,
		msg:    msg,
	}
}

type bareError struct {
	status status
	msg    string
}

const (
	_ status = iota
	statusInput
	statusInternal
)

type status int
