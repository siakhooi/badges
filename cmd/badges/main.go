package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const TEMPLATE_FILENAME = "./templates/release.svg.template"
const OUTPUT_DIRECTORY = "./docs/releases"
const COLOR_GREEN = "#4c1"
const COLOR_YELLOW = "#dfb317"

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

	data := Badge{"json2table", 50, COLOR_GREEN, "1.1.3", 25, ""}
	data.BadgeFile = fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Repo)

	file, err := os.OpenFile(data.BadgeFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	err = t.Execute(file, data)
	fmt.Printf("%s generated!\n", data.Repo)

	if err != nil {
		log.Print(err)
		return
	}

}
