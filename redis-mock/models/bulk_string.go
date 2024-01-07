package models

type BulkString struct {
	Message string
}

func NewBulkString(message string) *BulkString {
	return &BulkString{
		Message: message,
	}
}
