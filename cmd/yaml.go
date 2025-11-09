package cmd

import (
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"text/template"
)

type Servers struct {
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Name      string     `yaml:"name" validate:"required"`
	Listen    int        `yaml:"listen" validate:"required, min=1, max=65535"`
	Return    string     `yaml:"return"`
	SSL       bool       `yaml:"ssl"`
	Locations []Location `yaml:"locations"`
}

type Location struct {
	Name       string `yaml:"name"`
	Root_path  string `yaml:"root_path"`
	Alias_path string `yaml:"alias_path"`
}

func UnmarshalYAML(file string) (Servers, error) {
	var servers Servers

	data, err := os.ReadFile(file)
	if err != nil {
		return servers, err
	}

	validate := validator.New()
	if err := validate.Struct(servers); err != nil {
		return servers, err
	}
	if err := yaml.Unmarshal(data, &servers); err != nil {
		return servers, err
	}

	return servers, nil
}

func GenNgconf(servers Servers) (string, error) {
	tmpl, err := template.New("ngconf").Parse(ServerBlockTemplate)
	if err != nil {
		return "", err
	}

	var buf strings.Builder

	for _, server := range servers.Servers {
		err = tmpl.Execute(&buf, server)
		if err != nil {
			return "", err
		}
		buf.WriteString("\n\n")
	}

	return buf.String(), nil
}
