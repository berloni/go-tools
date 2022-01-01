package tools

import (
	"html/template"
	"os"
)

// GetHTMLTemplate returns a specific html template from a selected directory
func GetHTMLTemplate(directory string, templateName string) (*template.Template, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	t, err := template.ParseFiles(wd + "/" + directory + "/" + templateName)
	if err != nil {
		return nil, err
	}

	return t, nil
}
