package cmd

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"text/template"
)

type HttpConfig struct {
	Http Http `yaml:"http"`
}

type Http struct {
	Servers []Server `yaml:"servers"`

	ClientMaxBodySize string `yaml:"client_max_body_size"`
	KeepaliveTimeout  int    `yaml:"keepalive_timeout"`
	SendTimeout       int    `yaml:"send_timeout"`
	Gzip              bool   `yaml:"gzip"`
}

type Server struct {
	Name        string `yaml:"name" validate:"required"`
	Listen      int    `yaml:"listen" validate:"required,min=1,max=65535"`
	ListenV6    int    `yaml:"listen_v6" validate:"min=1,max=65535"`
	Return      string `yaml:"return"`
	Root_path_s string `yaml:"root_path_s"`
	Charset     string `yaml:"charset"`
	Index       string `yaml:"index"`

	SSL             bool   `yaml:"ssl"`
	SSL_buffer_size string `yaml:"ssl_buffer_size"`

	SSL_Cert  string   `yaml:"ssl_cert"`
	SSL_key   string   `yaml:"ssl_key"`
	SSL_proto []string `yaml:"ssl_protocols"`

	Locations []Location `yaml:"locations"`
}

type Location struct {
	Name              string `yaml:"name"`
	Root_path         string `yaml:"root_path"`
	Alias_path        string `yaml:"alias_path"`
	Proxy_pass        string `yaml:"proxy_pass"`
	Proxy_buffer_size string `yaml:"proxy_buffer_size"`
	Proxy_set_header  string `yaml:"proxy_set_header"`
}

func ParseServersFromYaml(file string) (HttpConfig, error) {
	var config HttpConfig

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	validate := validator.New()

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, err
	}

	if err := validate.Struct(config); err != nil {
		return config, err
	}
	// Validate http block
	if ValidateSizeStr(config.Http.ClientMaxBodySize) != true {
		panic(fmt.Sprintf("invalid client_max_body_size %s", config.Http.ClientMaxBodySize))
	}

	for _, server := range config.Http.Servers {
		// Validate return
		if server.Return != "" {
			if ValidateReturn(server.Return) != true {
				panic(fmt.Sprintf("invalid return %s", server.Return))
			}
		}
		// Validate ssl protocol
		if len(server.SSL_proto) != 0 {
			for _, protocol := range server.SSL_proto {
				if protocol != "TLSv1" && protocol != "TLSv1.1" && protocol != "TLSv1.2" && protocol != "TLSv1.3" {
					panic(fmt.Sprintf("invalid ssl protocol %s", protocol))
				}
			}
		}
		// Validate ssl_buffer_size
		if server.SSL_buffer_size != "" {
			if ValidateSizeStr(server.SSL_buffer_size) != true {
				panic(fmt.Sprintf("invalid ssl_buffer_size %s", server.SSL_buffer_size))
			}
		}
		for _, loc := range server.Locations {
			// Validate proxy_buffer_size
			if loc.Proxy_buffer_size != "" {
				if ValidateSizeStr(loc.Proxy_buffer_size) != true {
					panic(fmt.Sprintf("invalid proxy_buffer_size %s", loc.Proxy_buffer_size))
				}
			}

		}
	}

	return config, nil
}

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
