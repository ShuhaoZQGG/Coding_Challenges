package utils_test

import (
	"testing"

	"github.com/redis-mock/models"
	"github.com/redis-mock/utils"
)

func Test_RespParser_Serialize_Simple_Strings(t *testing.T) {
	inputs := []string{"OK", "hello world"}
	expectedResults := []string{"+OK\r\n", "+hello world\r\n"}
	for i, v := range inputs {
		message := models.NewSimpleString(v)
		actualResult, err := utils.Serialize[models.SimpleString](*message)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expectedResult := expectedResults[i]
		if expectedResult != actualResult {
			t.Fatalf("expectedResult does not match actualResult, %v vs %v", expectedResult, actualResult)
		}
	}
}

func Test_RespParser_Serialize_Bulk_Strings(t *testing.T) {
	inputs := []string{"hello", "", "-1"}
	expectedResults := []string{"$5\r\nhello\r\n", "$0\r\n\r\n", "$2\r\n-1\r\n"}
	for i, v := range inputs {
		message := models.NewBulkString(v)
		actualResult, err := utils.Serialize[models.BulkString](*message)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expectedResult := expectedResults[i]
		if expectedResult != actualResult {
			t.Fatalf("expectedResult does not match actualResult, %v vs %v", expectedResult, actualResult)
		}
	}
}

func Test_RespParser_Serialize_integers(t *testing.T) {
	inputs := []int64{12, 5, -7}
	expectedResults := []string{":12\r\n", ":5\r\n", ":-7\r\n"}
	for i, v := range inputs {
		actualResult, err := utils.Serialize[int64](v)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expectedResult := expectedResults[i]
		if expectedResult != actualResult {
			t.Fatalf("expectedResult does not match actualResult, %v vs %v", expectedResult, actualResult)
		}
	}
}

func Test_RespParser_Serialize_Errors(t *testing.T) {
	inputs := []string{"Error message", "ERR unknown command 'asdf'", "WRONGTYPE Operation against a key holding the wrong kind of value"}
	expectedResults := []string{"-Error message\r\n", "-ERR unknown command 'asdf'\r\n", "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n"}
	for i, v := range inputs {
		message := models.NewSimpleError(v)
		actualResult, err := utils.Serialize[models.SimpleError](*message)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expectedResult := expectedResults[i]
		if expectedResult != actualResult {
			t.Fatalf("expectedResult does not match actualResult, %v vs %v", expectedResult, actualResult)
		}
	}
}

func Test_RespParser_Serialize_Arrays(t *testing.T) {
	inputs := [][]any{
		{*models.NewBulkString("hello"), *models.NewBulkString("world")},
		{},
		{*models.NewBulkString("ping")},
		{*models.NewBulkString("get"), *models.NewBulkString("key")},
		{*models.NewBulkString("echo"), *models.NewBulkString("hello world")},
		{1, 2, 3, 4, *models.NewBulkString("hello")},
	}
	expectedResults := []any{
		"*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
		"*0\r\n",
		"*1\r\n$4\r\nping\r\n",
		"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
		"*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
		"*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r\n",
	}
	for i, v := range inputs {
		actualResult, err := utils.Serialize[[]any](v)
		if err != nil {
			t.Fatalf("%v", err)
		}
		expectedResult := expectedResults[i]
		if expectedResult != actualResult {
			t.Fatalf("expectedResult does not match actualResult, %v vs %v", expectedResult, actualResult)
		}
	}
}

func Test_RespParser_Deserialize_Simple_Strings(t *testing.T) {
}

func Test_RespParser_Deserialize_Bulk_Strings(t *testing.T) {

}

func Test_RespParser_Deserialize_integers(t *testing.T) {

}

func Test_RespParser_Deserialize_Errors(t *testing.T) {

}

func Test_RespParser_Deserialize_Arrays(t *testing.T) {

}
