package main

import (
	"golang.org/x/mod/semver"
	"strings"
)

func normalizeVersion(v string) string {
	if !strings.HasPrefix(v, "v") {
		return "v" + v
	}
	return v
}

func GetColor(version string) string {
	v := normalizeVersion(version)

	if semver.IsValid(v) {
		if semver.Compare(v, "v1.0.0") >= 0 {
			if semver.Prerelease(v) == "" {
				return COLOR_GREEN
			} else {
				return COLOR_ORANGE
			}
		} else {
			return COLOR_YELLOW
		}

	}
	return COLOR_GREY
}
