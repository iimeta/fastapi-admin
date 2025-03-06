package util

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"html/template"
)

func RenderTemplate(data any, name string) (string, error) {

	tmpl, err := template.New("tmpl").Parse(gfile.GetContents(fmt.Sprintf("./resource/template/email/%s.tmpl", name)))
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
