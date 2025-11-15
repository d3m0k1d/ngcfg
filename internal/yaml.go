package internal

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
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
	if ValidateSizeStr(config.Http.ClientMaxBodySize) != true {
		return config, fmt.Errorf("invalid client_max_body_size: %s", config.Http.ClientMaxBodySize)
	}

	if config.Http.KeepaliveTimeout < 1 {
		return config, fmt.Errorf("invalid keepalive_timeout: %d (must be >= 1)", config.Http.KeepaliveTimeout)
	}

	if config.Http.SendTimeout < 1 {
		return config, fmt.Errorf("invalid send_timeout: %d (must be >= 1)", config.Http.SendTimeout)
	}

	if config.Http.Worker_processes != "" {
		if config.Http.Worker_processes != "auto" {
			c, err := strconv.Atoi(config.Http.Worker_processes)
			if err != nil {
				return config, fmt.Errorf("invalid worker_processes: %s (must be 'auto' or a number)", config.Http.Worker_processes)
			}
			if c < 1 {
				return config, fmt.Errorf("invalid worker_processes: %s (must be >= 1)", config.Http.Worker_processes)
			}
		}
	}

	for i, server := range config.Http.Servers {
		// Validate return
		if server.Return != "" {
			if ValidateReturn(server.Return) != true {
				return config, fmt.Errorf("server[%d] '%s': invalid return directive: %s", i, server.Name, server.Return)
			}
		}

		// Validate ssl protocol
		if len(server.SSL_proto) != 0 {
			for j, protocol := range server.SSL_proto {
				if protocol != "TLSv1" && protocol != "TLSv1.1" && protocol != "TLSv1.2" && protocol != "TLSv1.3" {
					return config, fmt.Errorf("server[%d] '%s': invalid ssl_protocol[%d]: %s (allowed: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3)",
						i, server.Name, j, protocol)
				}
			}
		}

		// Validate ssl_buffer_size
		if server.SSL_buffer_size != "" {
			if ValidateSizeStr(server.SSL_buffer_size) != true {
				return config, fmt.Errorf("server[%d] '%s': invalid ssl_buffer_size: %s", i, server.Name, server.SSL_buffer_size)
			}
		}

		for k, loc := range server.Locations {
			// Validate proxy_buffer_size
			if loc.Proxy_buffer_size != "" {
				if ValidateSizeStr(loc.Proxy_buffer_size) != true {
					return config, fmt.Errorf("server[%d] '%s': location[%d] '%s': invalid proxy_buffer_size: %s",
						i, server.Name, k, loc.Name, loc.Proxy_buffer_size)
				}
			}
		}
	}

	return config, nil
}
