package models_test

import (
	"testing"

	"github.com/redis-mock/models"
)

func Test_ListStore_Lget_GetNonExistingKeyValue_Return_EmptyListAndFalse(t *testing.T) {
	listStore := models.NewListStore()
	list, ok := listStore.Lget("nonexist")
	if list.Len() != 0 {
		t.Fatalf("expecting empty list but get %v", list)
	}

	if ok != false {
		t.Fatalf("expecting false but get %v", ok)
	}
}

func Test_ListStore_Lget_GetExistingKeyValue_Return_Value(t *testing.T) {
	value := make([]string, 0, 2)
	value = append(value, "1")
	value = append(value, "2")
	listStore := models.NewListStore()
	listStore.Lpush("exist", value)
	list, ok := listStore.Lget("exist")
	if list.Len() != 2 {
		t.Fatalf("expecting list with range 2 but get %v", list.Len())
	}

	head := list.Front()
	for i := len(value) - 1; i >= 0; i-- {
		if head.Value.(string) != value[i] {
			t.Fatalf("list not match, at position %d expected %v but get %v", i, value[i], head.Value.(string))
		}
		head = head.Next()
	}

	if ok != true {
		t.Fatalf("expecting true but get %v", ok)
	}
}

func Test_ListStore_Lpush_AddNewKey_If_KeyNotExist(t *testing.T) {
	value := make([]string, 0, 2)
	value = append(value, "1")
	value = append(value, "2")
	listStore := models.NewListStore()
	listStore.Lpush("exist", value)
	_, ok := listStore.Lget("exist")
	if ok != true {
		t.Fatalf("expecting true but get %v", ok)
	}
}

func Test_ListStore_LRange_Start_0_Stop_2_Return_Whole_List(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", 0, 2)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"1", "2", "3"}
	for i, v := range output {
		if v != expected[i] {
			t.Fatalf("expected %v at position %d but actual %v", expected[i], i, v)
		}
	}
}

func Test_ListStore_LRange_Start_0_Stop_Negative1_Return_Whole_List(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", 0, -1)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"1", "2", "3"}
	for i, v := range output {
		if v != expected[i] {
			t.Fatalf("expected %v at position %d but actual %v", expected[i], i, v)
		}
	}
}

func Test_ListStore_LRange_Start_2_Stop_0_Return_Whole_List(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", 0, 2)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"1", "2", "3"}
	for i, v := range output {
		if v != expected[i] {
			t.Fatalf("expected %v at position %d but actual %v", expected[i], i, v)
		}
	}
}

func Test_ListStore_LRange_Start_Negative1000_Stop_1000_Return_Whole_List(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", -1000, 1000)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"1", "2", "3"}
	for i, v := range output {
		if v != expected[i] {
			t.Fatalf("expected %v at position %d but actual %v", expected[i], i, v)
		}
	}
}

func Test_ListStore_LRange_Start_1000_Stop_Negative1000_Return_Whole_List(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", 1000, -1000)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"1", "2", "3"}
	for i, v := range output {
		if v != expected[i] {
			t.Fatalf("expected %v at position %d but actual %v", expected[i], i, v)
		}
	}
}

func Test_ListStore_LRange_Start_1_Stop_1_Return_2(t *testing.T) {
	value := make([]string, 0, 3)
	// final order should be "1" "2" "3"
	value = append(value, "3")
	value = append(value, "2")
	value = append(value, "1")

	listStore := models.NewListStore()
	listStore.Lpush("list", value)
	output, err := listStore.Lrange("list", 1, 1)
	if err != nil {
		t.Fatalf("error occurs %v", err)
	}

	expected := []string{"2"}
	for i, v := range output {
		if v != expected[len(expected)-i-1] {
			t.Fatalf("expected %v at position %d but actual %v", expected[len(expected)-i-1], i, v)
		}
	}
}
