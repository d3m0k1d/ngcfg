package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
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

		data, err := UnmarshalYAML(file)
		if err != nil {
			return err
		}

		cfg, err := GenNgconf(data)
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

func Init() {
	rootCmd.AddCommand(yamlCmd)
	yamlCmd.Flags().StringP("file", "f", "", "YAML file path")
	yamlCmd.MarkFlagRequired("file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
