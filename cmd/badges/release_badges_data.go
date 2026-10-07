package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func GetReleaseBadges(token string, cfg Config) ([]Badge, error) {
	var badges []Badge
	var lookupErr error

	const center = BADGE_WIDTH / 2
	for _, repo := range cfg.Releases {
		badgeFile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, repo)
		release_version, err := GetReleaseVersion(token, repo)
		color := COLOR_GREY
		if err != nil {
			err = fmt.Errorf("%s: %w", repo, err)
			log.Print(err)
			release_version = "---"
			if !errors.Is(err, errNoRelease) {
				lookupErr = errors.Join(lookupErr, err)
			}
		} else {
			color = GetColor(release_version)
			release_version = strings.TrimPrefix(release_version, "v")
		}
		data := Badge{repo, BADGE_WIDTH, color, release_version, center, badgeFile}
		badges = append(badges, data)

	}

	return badges, lookupErr

}
