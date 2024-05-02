package config

import (
	"encoding/json"
	"os"
)

type AppConfiguration struct {
	Database string             `json:"database"`
	Redis    RedisConfiguration `json:"redis"`
	Http     HttpConfiguration  `json:"http"`
}

type RedisConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type HttpConfiguration struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"readTimeout"`
	WriteTimeout int    `mapstructure:"writeTimeout"`
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
