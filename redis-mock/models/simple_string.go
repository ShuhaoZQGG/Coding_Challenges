package models

type SimpleString struct {
	Message string
}

func NewSimpleString(message string) *SimpleString {
	return &SimpleString{
		Message: message,
	}
}
