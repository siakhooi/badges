package main

import (
	"strings"

	"golang.org/x/mod/semver"
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
				return COLOR_STABLE
			} else {
				return COLOR_STABLE_BETA
			}
		} else {
			return COLOR_PRE_STABLE
		}

	}
	return COLOR_UNKNOWN
}
