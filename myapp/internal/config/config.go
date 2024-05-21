package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"server_port"`
	LogFilePath string `mapstructure:"log_file_path"`
	// Add other configurations here
}

func LoadConfig() (*Config, error) {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" {
		profile = "dev" // default to dev profile
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Construct the path to the config directory
	configPath := filepath.Join(cwd, "configs")

	viper.SetConfigName(fmt.Sprintf("config.%s", profile))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
