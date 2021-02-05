package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var generatedConfig string
var customConfig string

func parseExArguments(args []string) (string, string, []string) {

	generatedConfigFile := os.Getenv("EXPANDER_GENERATED_CONF")
	customConfigFile := os.Getenv("EXPANDER_CUSTOM_CONF")

	if customConfig != "" {
		customConfigFile = customConfig
	}
	if generatedConfig != "" {
		generatedConfigFile = generatedConfig
	}

	input, err := ParseInput(args)

	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		return "", "", nil
	}

	return generatedConfigFile, customConfigFile, input

}

func expand(generatedConfigFile string, customConfigFile string, expressions []string) []string {
	abbreviations := make(map[string]string)

	expander.ParseConfigFile(generatedConfigFile, abbreviations)
	expander.ParseConfigFile(customConfigFile, abbreviations)

	if len(abbreviations) == 0 {
		fmt.Println("No mapping found, exiting")
		return nil
	}

	return expander.ExpandExpressions(expressions, abbreviations)
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
		generatedConfigFile, customConfigFile, expressions := parseExArguments(args)

		// perform logic
		out := expand(generatedConfigFile, customConfigFile, expressions)

		// print output
		fmt.Println(strings.Join(out, " "))

	},
}

func init() {
	expandCmd.PersistentFlags().StringVar(&generatedConfig, "generated-config", "", "path to the generated config file to use for expandsion. Default is the EXPANDER_GENERATED_CONF env var. ")
	expandCmd.PersistentFlags().StringVar(&customConfig, "custom-config", "", "path to the custom config file to use for expandsion. Default is the EXPANDER_CUSTOM_CONF env var. ")
	rootCmd.AddCommand(expandCmd)
}
