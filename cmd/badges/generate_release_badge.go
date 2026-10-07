package main

import (
	"fmt"
	"html/template"
	"os"
)

func GenerateReleaseBadge(t *template.Template, data Badge) error {
	file, err := os.OpenFile(data.BadgeFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return fmt.Errorf("Fail to open file %s, %w", data.BadgeFile, err)
	}
	defer file.Close()

	err = t.Execute(file, data)

	if err != nil {
		return fmt.Errorf("Fail to generate file %s, %w", data.BadgeFile, err)
	}
	fmt.Printf("Generated: %s %s %s\n", data.BadgeFile, data.Repo, data.Version)
	return nil
}
