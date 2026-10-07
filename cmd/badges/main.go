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

func generateIndexHtml() {
	fmt.Println("TODO")
}
func main() {
	generateReleaseBadges()
	generateIndexHtml()
}
func generateReleaseBadges() {

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

	badges, err1 := GetReleaseBadges()
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
