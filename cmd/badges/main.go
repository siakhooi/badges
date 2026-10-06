package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const TEMPLATE_FILENAME = "./templates/release.svg.template"
const OUTPUT_DIRECTORY = "./docs/releases"

type Release struct {
	Name    string
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
	err = os.MkdirAll(OUTPUT_DIRECTORY, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

	data := Release{"json2table", 100, "red", "1.1.3", 50}
	outfile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Name)

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
