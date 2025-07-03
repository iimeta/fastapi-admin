package util

import (
	"bytes"
	"html/template"
)

func RenderTemplate(name, content string, data map[string]any) (string, error) {

	tmpl, err := template.New(name).Parse(content)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err = tmpl.Execute(&buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
