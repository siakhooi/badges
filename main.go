package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const TEMPLATE_FILENAME = "./templates/badge.svg.template"

type Release struct {
	Width   uint
	Color   string
	Version string
	Center  uint
}

func main() {
	t, err := template.ParseFiles(TEMPLATE_FILENAME)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("Template loaded!")
	err = os.MkdirAll("docs/releases", os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

	data := Release{100, "red", "1.1.3", 50}
	outfile := "docs/releases/json2table.svg"

	file, err := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	err = t.Execute(file, data)

	if err != nil {
		log.Print(err)
		return
	}

}
