package main

import (
	"errors"
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

	badges := GetBadges()

	for _, badge := range badges {
		err = Generate(t, badge)
		if err != nil {
			log.Print(err)
		}
	}
}
func GetBadges() []Badge {
	var badges []Badge

	data := Badge{"json2table", 50, COLOR_GREEN, "1.1.3", 25, ""}
	data.BadgeFile = fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Repo)

	badges = append(badges, data)

	data = Badge{"fibo-planner", 50, COLOR_YELLOW, "0.1.3", 25, ""}
	data.BadgeFile = fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Repo)
	badges = append(badges, data)

	return badges

}
func Generate(t *template.Template, data Badge) error {
	file, err := os.OpenFile(data.BadgeFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return errors.New(fmt.Sprintf("Fail to open file: %s", data.BadgeFile))
	}
	defer file.Close()

	err = t.Execute(file, data)
	fmt.Printf("%s generated!\n", data.Repo)

	if err != nil {
		return errors.New(fmt.Sprintf("Fail to generate file: %s", data.BadgeFile))
	}
	return nil
}
