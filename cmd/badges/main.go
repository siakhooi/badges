package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Releases []string `yaml:"releases"`
}
type Badge struct {
	Repo      string
	Width     float32
	Color     string
	Version   string
	Center    float32
	BadgeFile string
}

func main() {
	t, err := template.ParseFiles(TEMPLATE_FILENAME)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("Template loaded!")
	err = os.MkdirAll(OUTPUT_DIRECTORY, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

	badges, err1 := GetBadges()
	if err1 != nil {
		log.Print(err1)
		return
	}

	for _, badge := range badges {
		err = Generate(t, badge)
		if err != nil {
			log.Print(err)
		}
	}
}
func GetColor(version string) string {
	return COLOR_GREEN
}

func GetBadges() ([]Badge, error) {
	var badges []Badge

	configraw, err := os.ReadFile(BADGES_CONFIG_FILE)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Fail to read config: %s", BADGES_CONFIG_FILE)
	}
	var cfg Config
	if err := yaml.Unmarshal(configraw, &cfg); err != nil {
		log.Fatalf("error %v", err)
	}
	for _, repo := range cfg.Releases {
		badgeFile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, repo)
		release_version, _, err := GetReleaseVersion(repo)
		if err != nil {
			log.Print(err)
			continue
		}
		color := GetColor(release_version)
		data := Badge{repo, BADGE_WIDTH, color, release_version, BADGE_X, badgeFile}

		badges = append(badges, data)

	}

	return badges, nil

}
