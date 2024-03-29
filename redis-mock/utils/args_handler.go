package utils

import (
	"fmt"
	"strconv"

	"github.com/redis-mock/models"
)

var serializer = new(models.RespSerializer)

func HandleDefault(args []string, store *models.StringStore) (response string, err error) {
	response, err = serializer.Serialize(*models.NewSimpleString("OK"))
	return response, err
}

func HandlePing(args []string, store *models.StringStore) (response string, err error) {
	response, err = serializer.Serialize(*models.NewSimpleString("PONG"))
	return response, err
}

func HandleSet(args []string, store *models.StringStore) (response string, err error) {
	store.Set(args[0], args[1])
	response, err = serializer.Serialize(*models.NewSimpleString("OK"))
	return response, err
}

func HandleGet(args []string, store *models.StringStore) (response string, err error) {
	getResult, exist := store.Get(args[0])
	if !exist {
		response = serializer.SerializeErrors(fmt.Sprintf("get key %v does not exist", args[0]))
	} else {
		response, err = serializer.Serialize(*models.NewSimpleString(getResult))
	}

	return response, err
}

func HandleExists(args []string, store *models.StringStore) (response string, err error) {
	response = serializer.SerializeIntegers(int64(0))
	if len(args) == 1 {
		_, exist := store.Get(args[0])
		if exist {
			response = serializer.SerializeIntegers(int64(1))
			return response, nil
		}
	} else {
		for _, v := range args {
			_, exist := store.Get(v)
			if exist {
				response = serializer.SerializeIntegers(int64(2))
				return response, nil
			}
		}
	}
	return response, nil
}

func HandleDelete(args []string, store *models.StringStore) (response string, err error) {
	response = serializer.SerializeIntegers(int64(0))
	if len(args) == 1 {
		ok := store.Del(args[0])
		if ok {
			response = serializer.SerializeIntegers(int64(1))
			return response, nil
		}
	} else {
		for _, v := range args {
			ok := store.Del(v)
			if ok {
				response = serializer.SerializeIntegers(int64(2))
				return response, nil
			}
		}
	}
	return response, nil
}

func HandleIncr(args []string, store *models.StringStore) (response string, err error) {
	key := args[0]
	value, exist := store.Get(key)
	if !exist {
		response = serializer.SerializeErrors("Error: key not exist")
		return response, nil
	}

	valueInInt, err := strconv.Atoi(value)
	if err != nil {
		response = serializer.SerializeErrors("Error: key is not type int")
		return response, nil
	}

	valueInInt += 1
	store.Set(key, strconv.Itoa(valueInInt))
	response = serializer.SerializeIntegers(int64(valueInInt))
	return response, nil
}

func HandleDecr(args []string, store *models.StringStore) (response string, err error) {
	key := args[0]
	value, exist := store.Get(key)
	if !exist {
		response = serializer.SerializeErrors("Error: key not exist")
		return response, nil
	}

	valueInInt, err := strconv.Atoi(value)
	if err != nil {
		response = serializer.SerializeErrors("Error: key is not type int")
		return response, nil
	}

	valueInInt -= 1
	store.Set(key, strconv.Itoa(valueInInt))
	response = serializer.SerializeIntegers(int64(valueInInt))
	return response, nil
}
