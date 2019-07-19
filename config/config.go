package config

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/barbourkd/docdeliver/devices"
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

// GenerateDevices returns a map of devices from the config
func (b *BaseConfig) GenerateDevices() (map[string]devices.Device, error) {
	deviceMap := map[string]devices.Device{}
	var err error

	for name, device := range b.Devices {
		switch device.Type {
		case "ReturnPrinter":
			deviceMap[name] = devices.NewReturnPrinter(name)
		default:
			err = errors.New("Invalid device type")
			break
		}

	}
	return deviceMap, err
}
