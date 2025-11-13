package cmd

import (
	"strings"
	"testing"
)

func TestGenNgconf(t *testing.T) {
	config := HttpConfig{
		Http: Http{
			ClientMaxBodySize: "100m",
			KeepaliveTimeout:  65,
			SendTimeout:       30,
			Servers: []Server{
				{
					Name:        "test",
					Listen:      80,
					ListenV6:    80,
					Return:      "301 https://example.com",
					Root_path_s: "/home/d3m0k1d/repo/ngcfg/test",
					Charset:     "utf-8",
					Index:       "index.html",
					SSL:         false,
					Locations: []Location{
						{
							Name:       "test",
							Alias_path: "/home/d3m0k1d/repo/ngcfg/test",
							Root_path:  "/home/d3m0k1d/repo/ngcfg/test",
						},
					},
				},
			},
		},
	}

	result, err := GenNgconf(config)
	if err != nil {
		t.Fatalf("GenNgconf failed: %v", err)
	}

	expectedStrings := []string{
		"http {",
		"client_max_body_size 100m;",
		"keepalive_timeout 65;",
		"send_timeout 30;",
		"server_name test;",
		"listen 80;",
		"listen [::]:80;",
		"charset utf-8;",
		"root /home/d3m0k1d/repo/ngcfg/test;",
		"return 301 https://example.com;",
		"location /test {",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected output to contain %q, but it doesn't. Got: %s", expected, result)
		}
	}
}

