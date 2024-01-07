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

}

func Test_RespParser_Serialize_Errors(t *testing.T) {

}

func Test_RespParser_Serialize_Arrays(t *testing.T) {

}

func Test_RespParser_Deserialize_Simple_Strings(t *testing.T) {
	t.Log("x")
}

func Test_RespParser_Deserialize_Bulk_Strings(t *testing.T) {

}

func Test_RespParser_Deserialize_integers(t *testing.T) {

}

func Test_RespParser_Deserialize_Errors(t *testing.T) {

}

func Test_RespParser_Deserialize_Arrays(t *testing.T) {

}
