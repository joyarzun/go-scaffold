package scaffold

import (
	"bytes"
	"text/template"
)

func ApplyTemplate(templateContent string, config interface{}) (string, error) {
	template, err := template.New("").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = template.Execute(&result, config)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
