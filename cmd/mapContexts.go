package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate abbreviations for available kubernetes contexts",
	Long:  "Map available kubernetes contexts, and print the abbreviations. These will be available for expansion in future runs.",
	Run: func(cmd *cobra.Command, args []string) {
		// do something
		fmt.Println("this is where I will map the contexts")
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)
}
