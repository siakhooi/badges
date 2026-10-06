package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

const BADGES_CONFIG_FILE = "./badges.yaml"
const TEMPLATE_FILENAME = "./templates/release.svg.template"
const OUTPUT_DIRECTORY = "./docs/releases"
const COLOR_GREEN = "#4c1"
const COLOR_YELLOW = "#dfb317"

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
func GetBadges() ([]Badge, error) {
	var badges []Badge

	configraw, err := os.ReadFile(BADGES_CONFIG_FILE)
	if err != nil {
		log.Print(err)
		return nil, errors.New(fmt.Sprintf("Fail to read config: %s", BADGES_CONFIG_FILE))
	}
	var cfg Config
	if err := yaml.Unmarshal(configraw, &cfg); err != nil {
		log.Fatalf("erro %v", err)
	}
	for _, release := range cfg.Releases {
		badgeFile := fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, release)
		data := Badge{release, 50, COLOR_GREEN, "1.1.3", 25, badgeFile}

		badges = append(badges, data)

	}

	//data := Badge{"json2table", 50, COLOR_GREEN, "1.1.3", 25, ""}
	//data.BadgeFile = fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Repo)

	//badges = append(badges, data)

	//data = Badge{"fibo-planner", 50, COLOR_YELLOW, "0.1.3", 25, ""}
	//data.BadgeFile = fmt.Sprintf("%s/%s.svg", OUTPUT_DIRECTORY, data.Repo)
	//badges = append(badges, data)

	return badges, nil

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
