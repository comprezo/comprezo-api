package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port int `json:"port"`
}

func Load(path string) (*Config, error) {
	cfg, err := readFromFile(path)
	if err != nil {
		return nil, err
	}

	if err := validate(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func validate(cfg *Config) error {
	if cfg.Port <= 0 {
		return fmt.Errorf("port is not valid: %d", cfg.Port)
	}

	return nil
}

func readFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &cfg, nil
}
