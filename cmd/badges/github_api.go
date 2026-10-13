package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
}

var errNoRelease = errors.New("no release")
var githubClient = &http.Client{Timeout: 30 * time.Second}

const GITHUB_API = "https://api.github.com"
const GITHUB_OWNER = "siakhooi"

func GetReleaseVersion(token, repo string) (string, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", GITHUB_API, GITHUB_OWNER, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	resp, err := githubClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", errNoRelease
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("fail to get latest release, API returned %s", resp.Status)
	}
	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err

	}
	fmt.Printf("Github Release Retrieved: %s %s\n", repo, release.TagName)

	return release.TagName, nil
}
