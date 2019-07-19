package config

import (
	"github.com/BurntSushi/toml"
)

// BaseConfig holds our basic configuration details
type BaseConfig struct {
	Title   string
	Devices map[string]deviceConfig
}

type deviceConfig struct {
	Type      string
	Something string
}

// ParseConfig takes in a filename and returns a config
func ParseConfig(fileName string) (BaseConfig, error) {
	var config BaseConfig
	if _, err := toml.DecodeFile(fileName, &config); err != nil {
		return config, err
	}

	return config, nil
}
