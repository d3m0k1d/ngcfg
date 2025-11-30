package cmd

import (
	"fmt"
	"github.com/d3m0k1d/ngcfg/internal"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strconv"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "ngcfg",
	Short: "ngcfg",
	Long:  `ngcfg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Errorf("Please use the subcommands for help use ngcfg --help")
	},
}

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Work with YAML files",
	Long:  `Parse and process YAML configuration files`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		if file == "" {
			return fmt.Errorf("flag -f/--file is required")
		}

		ext := path.Ext(file)
		if ext != ".yaml" && ext != ".yml" {
			return fmt.Errorf("only .yaml and .yml files are supported, got: %s", ext)
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", file)
		}

		data, err := internal.ParseServersFromYaml(file)
		if err != nil {
			return err
		}

		cfg, err := internal.GenNgconf(data)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(os.Stdout, cfg)
		if err != nil {
			return err
		}

		return nil
	},
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate NGINX configuration from flags",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		listenStr, _ := cmd.Flags().GetString("listen")
		root, _ := cmd.Flags().GetString("root")
		proxyPass, _ := cmd.Flags().GetString("proxy-pass")
		ssl, _ := cmd.Flags().GetBool("ssl")
		sslCert, _ := cmd.Flags().GetString("cert")
		sslKey, _ := cmd.Flags().GetString("key")
		gzip, _ := cmd.Flags().GetBool("gzip")
		output, _ := cmd.Flags().GetString("output")
		charset, _ := cmd.Flags().GetString("charset")
		index, _ := cmd.Flags().GetString("index")

		if name == "" {
			return fmt.Errorf("flag --name (-n) is required")
		}
		if listenStr == "" {
			return fmt.Errorf("flag --listen (-l) is required")
		}

		listen, err := strconv.Atoi(strings.Split(listenStr+":0", ":")[strings.Count(listenStr, ":")])
		if strings.Contains(listenStr, ":") {
			parts := strings.Split(listenStr, ":")
			listen, err = strconv.Atoi(parts[len(parts)-1])
		} else {
			listen, err = strconv.Atoi(listenStr)
		}
		if err != nil || listen < 1 || listen > 65535 {
			return fmt.Errorf("invalid port: must be between 1 and 65535")
		}

		if ssl && (sslCert == "" || sslKey == "") {
			return fmt.Errorf("SSL enabled but certificate or key path not provided")
		}

		server := internal.Server{
			Name:        name,
			Listen:      listen,
			Root_path_s: root,
			Charset:     charset,
			Index:       index,
			SSL:         ssl,
			SSL_Cert:    sslCert,
			SSL_key:     sslKey,
		}

		if proxyPass != "" {
			server.Locations = []internal.Location{
				{
					Name:       "/",
					Proxy_pass: proxyPass,
					Proxy_set_header: []string{
						"Host $host",
						"X-Real-IP $remote_addr",
						"X-Forwarded-For $proxy_add_x_forwarded_for",
						"X-Forwarded-Proto $scheme",
					},
				},
			}
		}

		http := internal.Http{
			Servers: []internal.Server{server},
			Gzip:    gzip,
		}

		config := internal.HttpConfig{Http: http}

		nginxConfig, err := internal.GenNgconf(config)
		if err != nil {
			return fmt.Errorf("failed to generate NGINX config: %w", err)
		}

		if output != "" {
			if err := os.WriteFile(output, []byte(nginxConfig), 0644); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}
			fmt.Printf("✓ Config saved to: %s\n", output)
		} else {
			fmt.Println(nginxConfig)
		}

		return nil
	},
}

func Init() {
	rootCmd.AddCommand(yamlCmd)
	rootCmd.AddCommand(genCmd)

	yamlCmd.Flags().StringP("file", "f", "", "Path to YAML configuration file")
	yamlCmd.MarkFlagRequired("file")

	genCmd.Flags().StringP("name", "n", "", "Server name (required)")
	genCmd.MarkFlagRequired("name")

	genCmd.Flags().StringP("listen", "l", "", "Listen port or address:port (required)")
	genCmd.MarkFlagRequired("listen")

	genCmd.Flags().StringP("root", "r", "", "Server root directory")
	genCmd.Flags().StringP("proxy-pass", "", "", "Upstream proxy address")
	genCmd.Flags().StringP("charset", "c", "", "Default charset")
	genCmd.Flags().StringP("index", "i", "", "Default index files")

	genCmd.Flags().BoolP("ssl", "s", false, "Enable SSL/TLS")
	genCmd.Flags().String("cert", "", "Path to SSL certificate")
	genCmd.Flags().String("key", "", "Path to SSL private key")

	genCmd.Flags().BoolP("gzip", "g", false, "Enable gzip compression")

	genCmd.Flags().StringP("output", "o", "", "Output file path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
