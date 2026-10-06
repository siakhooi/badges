package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type GithubRelease struct {
	TagName    string `json:"tag_name"`
	Prerelease bool   `json:"prerelease"`
}

func GetReleaseVersion(repo string) (string, bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", GITHUB_OWNER, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", true, err
	}
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		return "", true, errors.New("GITHUB_TOKEN not set")
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", true, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", true, fmt.Errorf("Github API returned %s", resp.Status)
	}
	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", true, err

	}
	fmt.Printf("Github Release: %s %s %v\n", repo, release.TagName, release.Prerelease)

	return release.TagName, release.Prerelease, nil
}
