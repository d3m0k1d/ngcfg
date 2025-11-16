package internal

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseServersFromYaml(file string) (HttpConfig, error) {
	var config HttpConfig

	data, err := os.ReadFile(file)
	if err != nil {
		return config, fmt.Errorf("failed to read file: %w", err)
	}

	validate := validator.New()

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := validate.Struct(config); err != nil {
		return config, fmt.Errorf("validation failed: %w", err)
	}

	// Validate http block
	if err := ValidateHttp(config); err != nil {
		return config, err
	}

	// Validate servers
	for i, server := range config.Http.Servers {
		// Validate server basics
		if err := ValidateServer(server, i); err != nil {
			return config, err
		}

		// Validate SSL protocols
		if len(server.SSL_proto) != 0 {
			if err := ValidateSSLProtocols(server.SSL_proto, i, server.Name); err != nil {
				return config, err
			}
		}

		// Validate ssl_buffer_size
		if server.SSL_buffer_size != "" {
			if err := ValidateSizeStr(server.SSL_buffer_size); err != nil {
				return config, fmt.Errorf("server[%d] '%s': invalid ssl_buffer_size: %w",
					i, server.Name, err)
			}
		}

		// Validate locations
		for k, loc := range server.Locations {
			if err := ValidateLocation(loc, i, k, server.Name); err != nil {
				return config, err
			}
		}
	}

	return config, nil
}
