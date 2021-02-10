package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/spf13/cobra"
	"strings"
)

var expandCmd = &cobra.Command{
	Use:   "ex",
	Short: "Expand known abbreviations",
	Long:  "Expand Abbreviation",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// get parameters
		expressions, err := ParseInput(args)

		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		// perform logic
		data := ParseConfigData(configfile, expander.ValidateData)
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
