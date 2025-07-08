package util

import (
	"bytes"
	"html/template"
	"regexp"
)

var templateVarRegex = regexp.MustCompile(`{{\s*\.\s*([a-zA-Z_][a-zA-Z0-9_]*(?:\s*\.\s*[a-zA-Z_][a-zA-Z0-9_]*)*)\s*}}`)

func GetTemplateVariables(title, content string) []string {

	var (
		titleMatches   = templateVarRegex.FindAllStringSubmatch(title, -1)
		contentMatches = templateVarRegex.FindAllStringSubmatch(content, -1)
		seen           = make(map[string]bool)
		vars           []string
	)

	for _, m := range titleMatches {
		if len(m) > 1 && !seen[m[1]] {
			seen[m[1]] = true
			vars = append(vars, m[1])
		}
	}

	for _, m := range contentMatches {
		if len(m) > 1 && !seen[m[1]] {
			seen[m[1]] = true
			vars = append(vars, m[1])
		}
	}

	return vars
}

func RenderTemplate(title, content string, data map[string]any) (string, string, error) {

	titleTmpl, err := template.New("title").Parse(title)
	if err != nil {
		return "", "", err
	}

	contentTmpl, err := template.New("content").Parse(content)
	if err != nil {
		return "", "", err
	}

	var titleBuffer bytes.Buffer
	if err = titleTmpl.Execute(&titleBuffer, data); err != nil {
		return "", "", err
	}

	var contentBuffer bytes.Buffer
	if err = contentTmpl.Execute(&contentBuffer, data); err != nil {
		return "", "", err
	}

	return titleBuffer.String(), contentBuffer.String(), nil
}
