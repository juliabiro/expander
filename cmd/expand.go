package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var configfile string

func parseExArguments(args []string) (string, []string) {

	generatedConfigFile := os.Getenv("EXPANDER_CONFIG")

	if configfile != "" {
		generatedConfigFile = configfile
	}

	input, err := ParseInput(args)

	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		return "", nil
	}

	return generatedConfigFile, input

}

func expand(generatedConfigFile string, expressions []string) []string {
	//abbreviations := make(map[string]string)

	//expander.ParseConfigFile(generatedConfigFile, abbreviations)

	abbreviations2 := expander.ParseConfigData(generatedConfigFile)

	if len(abbreviations2) == 0 {
		fmt.Println("No mapping found, exiting")
		return nil
	}

	return expander.ExpandExpressions(expressions, abbreviations2)
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
		generatedConfigFile, expressions := parseExArguments(args)

		// perform logic
		out := expand(generatedConfigFile, expressions)

		// print output
		fmt.Println(strings.Join(out, " "))

	},
}

func init() {
	expandCmd.PersistentFlags().StringVar(&configfile, "config", "", "path to the generated config file to use for expandsion. Default is the EXPANDER_GENERATED_CONF env var. ")
	rootCmd.AddCommand(expandCmd)
}
