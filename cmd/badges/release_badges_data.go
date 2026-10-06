package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

func GetReleaseBadges() ([]Badge, error) {
	var badges []Badge

	configraw, err := os.ReadFile(BADGES_CONFIG_FILE)
	if err != nil {
		return nil, fmt.Errorf("Fail to read config: %s, %s", BADGES_CONFIG_FILE, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(configraw, &cfg); err != nil {
		return nil, fmt.Errorf("Fail to parse config data, %w", err)
	}
	for _, repo := range cfg.Releases {
		badgeFile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, repo)
		release_version, _, err := GetReleaseVersion(repo)
		if err != nil {
			log.Print(err)
			continue
		}
		color := GetColor(release_version)
		release_version = strings.TrimPrefix(release_version, "v")
		data := Badge{repo, BADGE_WIDTH, color, release_version, BADGE_X, badgeFile}

		badges = append(badges, data)

	}

	return badges, nil

}
