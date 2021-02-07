package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var configfile string

func ParseInput(args []string) ([]string, error) {
	// this will become more complex later
	return args, nil
}

var rootCmd = &cobra.Command{
	Use:   "expander",
	Short: "Expander expands abbreviations to longer strings based on a map.",
	Long:  `Expander can be used to expand abbreviations to longer strings based on a map. The intended use is to produce long argument strings for other applications`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.PersistentFlags().StringVar(&configfile, "config", "", "file containing the abbreviations mapping and the generated and custom configs")
}
