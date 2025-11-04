package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ngcfg",
	Short: "ngcfg",
	Long:  `ngcfg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Errorf("Please use the subcommands for help use ngcfg --help")
	},
}

var yamltestCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Work with YAML configurations",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		if file == "" {
			return fmt.Errorf("флаг -f/--file обязателен")
		}

		_, err = UnmarshalYAML(file)
		return err
	},
}

func init() {
	yamltestCmd.Flags().StringP("file", "f", "", "Путь к YAML файлу")
	yamltestCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(yamltestCmd)
}

func Init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
