package cmd

import (
	"gopkg.in/yaml.v3"
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
