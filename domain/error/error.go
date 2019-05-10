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
		bareError: bareErrorf(statusInput, format, a...),
	}
}

type ValidationError struct {
	*bareError
}

func bareErrorf(status status, format string, a ...interface{}) *bareError {
	return &bareError{
		status: status,
		msg:    fmt.Sprintf(format, a...),
	}
}

type bareError struct {
	status status
	msg    string
}

func (e *bareError) input() bool {
	return e.status == statusInput
}

func (e *bareError) internal() bool {
	return e.status == statusInternal
}

func (e *bareError) Error() string {
	return e.msg
}

const (
	_ status = iota
	statusInput
	statusInternal
)

type status int
