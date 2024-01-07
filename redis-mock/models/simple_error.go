package models

type SimpleError struct {
	Message string
}

func NewSimpleError(message string) *SimpleError {
	return &SimpleError{
		Message: message,
	}
}
