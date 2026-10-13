package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
)

func GetReleaseBadges(token string, cfg Config) ([]Badge, error) {
	badges := make([]Badge, len(cfg.Releases))
	var lookupErr error
	var wg sync.WaitGroup
	var mu sync.Mutex

	const center = BADGE_WIDTH / 2
	for i, repo := range cfg.Releases {
		wg.Add(1)

		go func(i int, repo string) {
			defer wg.Done()

			badgeFile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, repo)
			release_version, err := GetReleaseVersion(token, repo)
			color := COLOR_UNKNOWN
			if err != nil {
				err = fmt.Errorf("fail to get latest version %s: %w", repo, err)
				log.Print(err)
				release_version = "---"
				if !errors.Is(err, errNoRelease) {
					mu.Lock()
					lookupErr = errors.Join(lookupErr, err)
					mu.Unlock()
				}
			} else {
				color = GetColor(release_version)
				release_version = strings.TrimPrefix(release_version, "v")
			}
			badges[i] = Badge{repo, BADGE_WIDTH, color, release_version, center, badgeFile}
		}(i, repo)
	}
	wg.Wait()

	return badges, lookupErr

}
