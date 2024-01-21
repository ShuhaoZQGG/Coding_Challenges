package configs_test

import (
	"strings"
	"testing"

	"github.com/redis-mock/configs"
)

func Test_GetConfig_Throw_Error_If_Config_NotExist(t *testing.T) {
	_, err := configs.GetConfig("nonexist.yaml")
	if err == nil {
		t.Fatalf("Expected get err but not get err")
	}

	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Fatalf("Expected no such file or directory but not another error")
	}
}

func Test_GetConfig_Throw_Error_If_Config_Extension_NotYaml(t *testing.T) {
	_, err := configs.GetConfig("existbutwrongextension.json")
	if err == nil {
		t.Fatalf("Expected get err but not get err")
	}
}

func Test_GetConfig_Throw_Get_Expected_Configuration_Values(t *testing.T) {
	config, err := configs.GetConfig("config.yaml")
	if err != nil {
		t.Fatalf("Expected succes but get err")
	}

	expected_aof_on := true
	exepected_aof_location := "redis_default.aof"
	expected_aof_appendfsync := "always"

	actual_aof_on := config.Aof.On
	actual_aof_location := config.Aof.Location
	actual_aof_appendfsync := config.Aof.Appendfsync

	if expected_aof_on != actual_aof_on {
		t.Fatalf("expected %v vs actual %v", expected_aof_on, actual_aof_on)
	}

	if exepected_aof_location != actual_aof_location {
		t.Fatalf("expected %s vs actual %s", exepected_aof_location, actual_aof_location)
	}

	if expected_aof_appendfsync != actual_aof_appendfsync {
		t.Fatalf("expected %s vs actual %s", expected_aof_appendfsync, actual_aof_appendfsync)
	}
}
