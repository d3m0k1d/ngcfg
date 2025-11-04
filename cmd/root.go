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

func Init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
