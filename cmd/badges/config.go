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
	configRaw, err := os.ReadFile(BADGES_CONFIG_FILE)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read config %s: %w", BADGES_CONFIG_FILE, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(configRaw, &cfg); err != nil {
		return Config{}, fmt.Errorf("fail to parse config data %s: %w", BADGES_CONFIG_FILE, err)
	}
	return cfg, nil
}
