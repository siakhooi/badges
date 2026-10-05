package main

import (
	"fmt"
	"log"
	"text/template"
	"os"
)

const TEMPLATE_FILENAME = "./templates/badge.svg.template"
type Release struct {
	Width uint
	Color string
	Version string
	Center uint 
}
func main(){
  fmt.Println("Hello, World!")
	t, err := template.ParseFiles(TEMPLATE_FILENAME)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("Template loaded!")

	data := Release{100, "red", "1.1.3", 50}

	err = t.Execute(os.Stdout, data)

	if err != nil {
		log.Print(err)
		return
	}

}
