package cmd

import (
	"testing"
)

func TestGenNgconf(t *testing.T) {
	var test_server Servers

	test_server.Servers = append(test_server.Servers, Server{
		Name:        "test",
		Listen:      80,
		ListenV6:    80,
		Return:      "test",
		Root_path_s: "/home/d3m0k1d/repo/ngcfg/test",
		Charset:     "utf-8",
		Locations: []Location{
			{
				Name:       "test",
				Alias_path: "/home/d3m0k1d/repo/ngcfg/test",
				Root_path:  "/home/d3m0k1d/repo/ngcfg/test",
			},
		},
	})

	_, err := GenNgconf(test_server)
	if err != nil {
		t.Error(err)
	}
}
