package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	config := BaseConfig{}

	config.Title = "Test"
	config.Devices = map[string]deviceConfig{
		"sampleDevice": deviceConfig{Type: "ReturnPrinter"},
	}
}

func TestParseConfig(t *testing.T) {
	config, _ := ParseConfig("sample.cfg")
	want := "My Cool Sample Config"

	if config.Title != want {
		t.Errorf("got %q want %q", config.Title, want)
	}

	device := config.Devices["return"]

	if device.Type != "ReturnPrinter" {
		t.Errorf("Device type incorrect")
	}
}

func TestGenerateDevices(t *testing.T) {
	t.Run("good config", func(t *testing.T) {
		config, _ := ParseConfig("sample.cfg")

		devices, err := config.GenerateDevices()

		if len(devices) != 2 {
			t.Error("unwanted number of devices")
		}

		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}
	})

	t.Run("invalid device type", func(t *testing.T) {
		config, _ := ParseConfig("invalid.cfg")

		_, err := config.GenerateDevices()

		if err == nil {
			t.Errorf("expected an error but did not receive one")
		}
	})
}

func TestMissingConfig(t *testing.T) {
	_, err := ParseConfig("beepboop.cfg")

	if err == nil {
		t.Error("Expected error but did not get anything")
	}
}
