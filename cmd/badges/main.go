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
	cfg, err := GetConfig()
	if err != nil {
		log.Print(err)
		return
	}
	err = os.MkdirAll(OUTPUT_DIRECTORY, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

	generateIndexHtml(cfg)
	generateReleaseBadges(cfg)
}
func generateIndexHtml(cfg Config) {
	t, err := template.ParseFiles(INDEX_HTML_TEMPLATE)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("Index.html Template loaded!")
	GenerateIndexHtml(t, cfg)
}

func generateReleaseBadges(cfg Config) {

	t, err := template.ParseFiles(RELEASE_BADGE_TEMPLATE)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("Release Badge Template loaded!")

	badges, err1 := GetReleaseBadges(cfg)
	if err1 != nil {
		log.Print(err1)
		return
	}

	for _, badge := range badges {
		err = GenerateReleaseBadge(t, badge)
		if err != nil {
			log.Print(err)
		}
	}
}
