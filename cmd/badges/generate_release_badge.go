package main

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateReleaseBadge(t *template.Template, data Badge) error {
	file, err := os.OpenFile(data.BadgeFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return fmt.Errorf("Fail to open file: %s", data.BadgeFile)
	}
	defer file.Close()

	err = t.Execute(file, data)
	fmt.Printf("Generated: %s %s %s\n", data.BadgeFile, data.Repo, data.Version)

	if err != nil {
		return fmt.Errorf("Fail to generate file: %s", data.BadgeFile)
	}
	return nil
}
