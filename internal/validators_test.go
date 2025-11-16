package internal

import (
	"testing"
)

func TestValidateSizeStr(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "Valid size with k", input: "1k", wantErr: false},
		{name: "Valid size with m", input: "100m", wantErr: false},
		{name: "Valid size with g", input: "2g", wantErr: false},
		{name: "Valid size without unit", input: "512", wantErr: false},
		{name: "Empty string", input: "", wantErr: false},
		{name: "Invalid letter", input: "1x", wantErr: true},
		{name: "Invalid format negative", input: "-1", wantErr: true},
		{name: "Invalid format with space", input: "100 m", wantErr: true},
		{name: "Multiple units", input: "100mk", wantErr: true},
		{name: "Only letter", input: "m", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSizeStr(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSizeStr(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateReturn(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "Valid 301 redirect", input: "301 https://example.com", wantErr: false},
		{name: "Valid 302 redirect", input: "302 http://test.org/path", wantErr: false},
		{name: "Valid 404 error", input: "404 https://error.page", wantErr: false},
		{name: "Invalid status code", input: "600 https://example.com", wantErr: true},
		{name: "Missing protocol", input: "301 example.com", wantErr: true},
		{name: "No space separator", input: "301https://example.com", wantErr: true},
		{name: "Two digit code", input: "30 https://example.com", wantErr: true},
		{name: "Empty string", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateReturn(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateReturn(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "Valid HTTP URL", input: "http://example.com", wantErr: false},
		{name: "Valid HTTPS URL", input: "https://test.org/path", wantErr: false},
		{name: "Valid with subdomain", input: "https://api.example.com", wantErr: false},
		{name: "Valid with path and dash", input: "https://example.com/test-path", wantErr: false},
		{name: "Invalid without protocol", input: "example.com", wantErr: true},
		{name: "Invalid protocol", input: "ftp://example.com", wantErr: true},
		{name: "Empty string", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateURL(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateHttp(t *testing.T) {
	tests := []struct {
		name    string
		config  HttpConfig
		wantErr bool
	}{
		{
			name: "Valid config",
			config: HttpConfig{
				Http: Http{
					ClientMaxBodySize: "100m",
					KeepaliveTimeout:  65,
					SendTimeout:       30,
					Worker_processes:  "auto",
					Servers: []Server{
						{Name: "test", Listen: 80},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "No servers",
			config: HttpConfig{
				Http: Http{
					KeepaliveTimeout: 65,
					SendTimeout:      30,
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid keepalive timeout",
			config: HttpConfig{
				Http: Http{
					KeepaliveTimeout: 0,
					SendTimeout:      30,
					Servers: []Server{
						{Name: "test", Listen: 80},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid client_max_body_size",
			config: HttpConfig{
				Http: Http{
					ClientMaxBodySize: "invalid",
					KeepaliveTimeout:  65,
					SendTimeout:       30,
					Servers: []Server{
						{Name: "test", Listen: 80},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid worker_processes",
			config: HttpConfig{
				Http: Http{
					KeepaliveTimeout: 65,
					SendTimeout:      30,
					Worker_processes: "-1",
					Servers: []Server{
						{Name: "test", Listen: 80},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateHttp(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateHttp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSSLProtocols(t *testing.T) {
	tests := []struct {
		name       string
		protocols  []string
		serverIdx  int
		serverName string
		wantErr    bool
	}{
		{name: "Valid TLSv1.2 and TLSv1.3", protocols: []string{"TLSv1.2", "TLSv1.3"}, serverIdx: 0, serverName: "test", wantErr: false},
		{name: "Valid all versions", protocols: []string{"TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3"}, serverIdx: 0, serverName: "test", wantErr: false},
		{name: "Invalid protocol", protocols: []string{"TLSv1.4"}, serverIdx: 0, serverName: "test", wantErr: true},
		{name: "Invalid protocol SSLv3", protocols: []string{"SSLv3"}, serverIdx: 0, serverName: "test", wantErr: true},
		{name: "Mixed valid and invalid", protocols: []string{"TLSv1.2", "TLSv2.0"}, serverIdx: 0, serverName: "test", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSSLProtocols(tt.protocols, tt.serverIdx, tt.serverName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSSLProtocols() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateFilePath(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "Valid file path", input: "/etc/nginx", wantErr: false},
		{name: "Invalid file path", input: "invalid/path", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFilePath(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFilePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
