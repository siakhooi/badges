package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Releases []string `yaml:"releases"`
}

func GetConfig() (Config, error) {
	configraw, err := os.ReadFile(BADGES_CONFIG_FILE)
	if err != nil {
		return Config{}, fmt.Errorf("Fail to read config: %s, %s", BADGES_CONFIG_FILE, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(configraw, &cfg); err != nil {
		return Config{}, fmt.Errorf("Fail to parse config data, %w", err)
	}
	return cfg, nil
}
