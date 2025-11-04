package cmd

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Servers struct {
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Name      string     `yaml:"name"`
	Listen    int        `yaml:"listen"`
	SSL       bool       `yaml:"ssl"`
	Locations []Location `yaml:"locations"`
}

type Location struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func UnmarshalYAML(file string) (Servers, error) {
	var servers Servers

	data, err := os.ReadFile(file)
	if err != nil {
		return servers, err
	}
	if err := yaml.Unmarshal(data, &servers); err != nil {
		return servers, err
	}
	return servers, nil
}
