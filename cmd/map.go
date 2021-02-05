package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/abbreviator"
	"github.com/juliabiro/expander/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var expanderAbbrevations string
var expanderGeneratedConf string

func parseMapArguments(args []string) (generatedConfigFile string, input []string) {
	input, err := ParseInput(args)
	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		return "", nil
	}
	configfile := expanderGeneratedConf
	if configfile == "" {
		configfile = os.Getenv("EXPANDER_GENERATED_CONF")
	}
	return configfile, input
}

func abbreviate(expressions []string) map[string]string {
	abbreviations := make([]utils.StringPair, 0)
	abbreviator.ParseConfigFile(expanderAbbrevations, &abbreviations)

	if len(abbreviations) == 0 {
		fmt.Println("No mapping found.")
		return nil
	}

	// This is where the magic happens
	return abbreviator.AbbreviateExpressions(expressions, abbreviations)
}

func printOutput(data map[string]string, targetfile string) {
	if len(data) == 0 {
		fmt.Println("No abbreviations made. Not saving anything.")
		return
	}

	// format output
	out := utils.MakeSortedString(data)

	// print output
	fmt.Println("Generated Abbreviations:")
	fmt.Println(out)

	utils.WriteToFile(out, targetfile)
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate abbreviations for available kubernetes contexts",
	Long:  "Map available kubernetes contexts, and print the abbreviations. These will be available for expansion in future runs.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// get parameters
		generatedConfigFile, expressions := parseMapArguments(args)

		//perform logic
		data := abbreviate(expressions)

		//print output
		printOutput(data, generatedConfigFile)

	},
}

func init() {
	mapCmd.PersistentFlags().StringVar(&expanderAbbrevations, "abbreviations", "", "file containing the abbreviations to be applied")
	mapCmd.PersistentFlags().StringVar(&expanderGeneratedConf, "generated-config", "", "file to which generated conf should be written. Default is $EXPANDER_GENERATED_CONF")

	rootCmd.AddCommand(mapCmd)
}
