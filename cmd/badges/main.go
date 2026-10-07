package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Badge struct {
	Repo      string
	Width     float32
	Color     string
	Version   string
	Center    float32
	BadgeFile string
}

func main() {
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists || token == "" {
		log.Fatal("GITHUB_TOKEN not set")
	}
	cfg, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(OUTPUT_DIRECTORY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	generateIndexHtml(cfg)
	generateReleaseBadges(token, cfg)
}
func generateIndexHtml(cfg Config) {
	t, err := template.ParseFiles(INDEX_HTML_TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Index.html Template loaded!")
	GenerateIndexHtml(t, cfg)
}

func generateReleaseBadges(token string, cfg Config) {

	t, err := template.ParseFiles(RELEASE_BADGE_TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Release Badge Template loaded!")

	badges, lookupErr := GetReleaseBadges(token, cfg)

	for _, badge := range badges {
		err = GenerateReleaseBadge(t, badge)
		if err != nil {
			log.Printf("Fail to generate release badge for %s: %v", badge.Repo, err)
		}
	}
	if lookupErr != nil {
		log.Fatal(lookupErr)
	}

}
