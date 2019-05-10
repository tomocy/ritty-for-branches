package error

func NewValidationError() *ValidationError {
	return new(ValidationError)
}

type ValidationError struct{}
