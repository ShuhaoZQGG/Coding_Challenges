package configs

import (
	"os"

	"github.com/redis-mock/constants"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Aof AOFConfig `yaml:"aof"`
}

type AOFConfig struct {
	On          bool   `yaml:"on"`
	Location    string `yaml:"location"`
	Appendfsync string `yaml:"appendfsync"`
}

func GetConfig(location string) (config *Config, err error) {
	// TODO: Add .json config support
	file, err := os.ReadFile(location)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return &Config{}, err
	}

	config.Aof.Appendfsync = enforceAppendfsyncConfig(config.Aof.Appendfsync)

	return config, nil
}

func enforceAppendfsyncConfig(appendfsync string) string {
	if appendfsync != constants.APPENDSYNC_DEFAULT &&
		appendfsync != constants.APPENDSYNC_ALWAYS &&
		appendfsync != constants.APPENDSYNC_NO {
		appendfsync = constants.APPENDSYNC_DEFAULT
	}

	return appendfsync
}
