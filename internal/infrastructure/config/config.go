package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
