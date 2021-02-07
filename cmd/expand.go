package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func parseExArguments(args []string) []string {

	configEnvVar := os.Getenv("EXPANDER_CONFIG")

	if configEnvVar != "" {
		configfile = configEnvVar
	}

	input, err := ParseInput(args)

	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		return nil
	}

	return input

}

var expandCmd = &cobra.Command{
	Use:   "ex",
	Short: "Expand known abbreviations",
	Long:  "Expand Abbreviation",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// get parameters
		expressions := parseExArguments(args)

		// perform logic
		data := expander.ParseConfigData(configfile)
		if data == nil {
			return
		}

		out := expander.ExpandExpressions(expressions, data)

		// print output
		fmt.Println(strings.Join(out, " "))

	},
}

func init() {
	rootCmd.AddCommand(expandCmd)
}
