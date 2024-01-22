package models

import (
	"errors"
	"fmt"

	"github.com/redis-mock/constants"
)

type Serializer interface {
	Serialize(message any) (string, error)
	SerializeSimpleStrings(message string) string
	SerializeBulkStrings(message string) string
	SerializeBulkStringsForNull() string
	SerializeErrors(message string) string
	SerializeIntegers(message int64) string
	SerializeArrays(messages []any) string
}

type RespSerializer struct {
}

func (respSerializer *RespSerializer) Serialize(message any) (string, error) {
	serializedOutput := ""
	switch v := any(message).(type) {

	case SimpleString:
		serializedOutput = respSerializer.SerializeSimpleStrings(v.Message)
	case BulkString:
		serializedOutput = respSerializer.SerializeBulkStrings(v.Message)
	case SimpleError:
		serializedOutput = respSerializer.SerializeErrors(v.Message)
	case []any:
		serializedOutput = respSerializer.SerializeArrays(v)
	case int64:
		serializedOutput = respSerializer.SerializeIntegers(v)
	case int32:
		serializedOutput = respSerializer.SerializeIntegers(int64(v))
	case int16:
		serializedOutput = respSerializer.SerializeIntegers(int64(v))
	case int8:
		serializedOutput = respSerializer.SerializeIntegers(int64(v))
	case int:
		serializedOutput = respSerializer.SerializeIntegers(int64(v))
	case nil:
		serializedOutput = respSerializer.SerializeBulkStringsForNull()
	default:
		return "", errors.New("Message type is not available")
	}

	return serializedOutput, nil
}

func (respSerializer *RespSerializer) SerializeSimpleStrings(message string) string {
	prefix := constants.SIMPLE_STRINGS_PREFIX
	crlf := constants.CRLF

	return prefix + message + crlf
}

func (respSerializer *RespSerializer) SerializeBulkStrings(message string) string {
	prefix := constants.BULK_STRINGS_PREFIX
	length := len(message)
	crlf := constants.CRLF
	return fmt.Sprintf("%s%d%s%s%s", prefix, length, crlf, message, crlf)
}

func (respSerializer *RespSerializer) SerializeBulkStringsForNull() string {
	prefix := constants.BULK_STRINGS_PREFIX
	content := "-1"
	crlf := constants.CRLF

	return prefix + content + crlf
}

func (respSerializer *RespSerializer) SerializeErrors(message string) string {
	prefix := constants.ERRORS_PREFIX
	crlf := constants.CRLF

	return prefix + message + crlf
}

func (respSerializer *RespSerializer) SerializeIntegers(message int64) string {
	prefix := constants.INTEGERS_PREFIX

	crlf := constants.CRLF

	return fmt.Sprintf("%s%d%s", prefix, message, crlf)
}

func (respSerializer *RespSerializer) SerializeArrays(messages []any) string {
	prefix := constants.ARRAYS_PREFIX
	length := len(messages)
	crlf := constants.CRLF
	output := fmt.Sprintf("%s%v%s", prefix, length, crlf)
	for _, message := range messages {
		serializedOutput, _ := respSerializer.Serialize(message)
		output += serializedOutput
	}
	return output
}
