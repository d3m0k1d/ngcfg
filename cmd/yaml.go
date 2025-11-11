package cmd

import (
	"fmt"
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
	Name        string `yaml:"name" validate:"required"`
	Listen      int    `yaml:"listen" validate:"required,min=1,max=65535"`
	ListenV6    int    `yaml:"listen_v6" validate:"min=1,max=65535"`
	Return      string `yaml:"return"`
	Root_path_s string `yaml:"root_path_s"`
	Charset     string `yaml:"charset"`
	Index       string `yaml:"index"`

	SSL       bool     `yaml:"ssl"`
	SSL_Cert  string   `yaml:"ssl_cert"`
	SSL_key   string   `yaml:"ssl_key"`
	SSL_proto []string `yaml:"ssl_protocols"`

	Locations []Location `yaml:"locations"`
}

type Location struct {
	Name       string `yaml:"name"`
	Root_path  string `yaml:"root_path"`
	Alias_path string `yaml:"alias_path"`
}

var ssl_protocols = []string{"TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3"}

func ParseServersFromYaml(file string) (Servers, error) {
	var servers Servers

	data, err := os.ReadFile(file)
	if err != nil {
		return servers, err
	}

	validate := validator.New()

	if err := yaml.Unmarshal(data, &servers); err != nil {
		return servers, err
	}
	if err := validate.Struct(servers); err != nil {
		return servers, err
	}

	return servers, nil
}

func GenNgconf(servers Servers) (string, error) {
	if len(servers.Servers) == 0 {
		return "", fmt.Errorf("no servers in yaml file")
	}

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
