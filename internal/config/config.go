package config

import (
	"encoding/json"
	"os"
)

type AppConfiguration struct {
	Database string `json:"database"`
}

func Initialize(configPath string) (*AppConfiguration, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := AppConfiguration{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
