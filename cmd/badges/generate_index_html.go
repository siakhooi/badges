package main

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateIndexHtml(t *template.Template, cfg Config) error {
	file, err := os.OpenFile(INDEX_HTML_OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return fmt.Errorf("Fail to open file: %s", INDEX_HTML_OUTPUT)
	}
	defer file.Close()

	err = t.Execute(file, cfg)
	fmt.Printf("Generated: %s\n", INDEX_HTML_OUTPUT)

	if err != nil {
		return fmt.Errorf("Fail to generate file: %s", INDEX_HTML_OUTPUT)
	}
	return nil
}
