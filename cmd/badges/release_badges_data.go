package main

import (
	"fmt"
	"log"
	"strings"
)

func GetReleaseBadges(cfg Config) ([]Badge, error) {
	var badges []Badge

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
