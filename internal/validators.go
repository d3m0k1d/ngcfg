package internal

import (
	"fmt"
	"regexp"
	"strconv"
)

// Validate blocks
func ValidateHttp(config HttpConfig) error {
	if len(config.Http.Servers) == 0 {
		return fmt.Errorf("no servers defined")
	}

	if config.Http.ClientMaxBodySize != "" {
		if err := ValidateSizeStr(config.Http.ClientMaxBodySize); err != nil {
			return fmt.Errorf("invalid client_max_body_size: %w", err)
		}
	}

	if config.Http.KeepaliveTimeout < 1 {
		return fmt.Errorf("invalid keepalive_timeout: %d (must be >= 1)", config.Http.KeepaliveTimeout)
	}

	if config.Http.SendTimeout < 1 {
		return fmt.Errorf("invalid send_timeout: %d (must be >= 1)", config.Http.SendTimeout)
	}

	if config.Http.Worker_processes != "" {
		if config.Http.Worker_processes != "auto" {
			c, err := strconv.Atoi(config.Http.Worker_processes)
			if err != nil {
				return fmt.Errorf("invalid worker_processes: %s (must be 'auto' or a number)", config.Http.Worker_processes)
			}
			if c < 1 {
				return fmt.Errorf("invalid worker_processes: %s (must be >= 1)", config.Http.Worker_processes)
			}
		}
	}

	return nil
}

func ValidateServer(server Server, index int) error {
	if server.Return != "" {
		if err := ValidateReturn(server.Return); err != nil {
			return fmt.Errorf("server[%d] '%s': invalid return directive: %w", index, server.Name, err)
		}
	}

	return nil
}

func ValidateLocation(loc Location, serverIdx, locIdx int, serverName string) error {
	if loc.Proxy_buffer_size != "" {
		if err := ValidateSizeStr(loc.Proxy_buffer_size); err != nil {
			return fmt.Errorf("server[%d] '%s': location[%d] '%s': invalid proxy_buffer_size: %w",
				serverIdx, serverName, locIdx, loc.Name, err)
		}
	}
	return nil
}

func ValidateSSLProtocols(protocols []string, serverIdx int, serverName string) error {
	validProtocols := map[string]bool{
		"TLSv1":   true,
		"TLSv1.1": true,
		"TLSv1.2": true,
		"TLSv1.3": true,
	}

	for j, protocol := range protocols {
		if !validProtocols[protocol] {
			return fmt.Errorf("server[%d] '%s': invalid ssl_protocol[%d]: %s (allowed: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3)",
				serverIdx, serverName, j, protocol)
		}
	}

	return nil
}

// Validators for specific fields
func ValidateSizeStr(s string) error {
	if s == "" {
		return nil
	}

	re := regexp.MustCompile(`^\d+[mkg]?$`)
	if !re.MatchString(s) {
		return fmt.Errorf("%s (expected format: 100m, 1g, 4k)", s)
	}

	return nil
}

func ValidateReturn(s string) error {
	re := regexp.MustCompile(`^[1-5]\d{2}\s+https?://[a-zA-Z0-9./_-]+$`)
	if !re.MatchString(s) {
		return fmt.Errorf("%s (expected format: 301 https://example.com)", s)
	}

	return nil
}

func ValidateURL(s string) error {
	re := regexp.MustCompile(`^https?://[a-zA-Z0-9./_-]+$`)
	if !re.MatchString(s) {
		return fmt.Errorf("%s (expected format: http://example.com)", s)
	}

	return nil
}

