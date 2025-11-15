package internal

import (
	"fmt"

	"os"

	"strings"
	"text/template"
)

func GenNgconf(config HttpConfig) (string, error) {
	if len(config.Http.Servers) == 0 {
		return "", fmt.Errorf("no servers in http config")
	}

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl, err := template.New("nginx").Funcs(funcMap).Parse(FullTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder

	err = tmpl.Execute(&buf, config.Http)
	if err != nil {
		return "", err
	}

	fmt.Println("Written to nginx.conf")
	os.Create("nginx.conf")
	os.WriteFile("nginx.conf", []byte(buf.String()), 0644)
	return buf.String(), nil
}
