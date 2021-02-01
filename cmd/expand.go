package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/expander"
	"github.com/spf13/cobra"
	"os"
)

var expandCmd = &cobra.Command{
	Use:   "ex",
	Short: "Expand known abbreviations",
	Long:  "Expand Abbreviation",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		configfile := os.Getenv("EXPANDER_CONF")
		expander := expander.NewExpander(configfile)
		err := expander.ParseConfigFile()
		if err != nil {
			fmt.Printf("Failed to parse configfile %s, error is %s.", configfile, err)
		}

		input, err := ParseInput(args)

		if err != nil {
			fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		}
		for _, c := range input {
			fmt.Print(expander.Expand(c))
		}
		// do something

	},
}

func init() {
	rootCmd.AddCommand(expandCmd)
}
