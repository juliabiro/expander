package cmd

import (
	"errors"
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

func abbreviate(expressions []string) (*utils.ExpanderData, error) {
	//abbreviations := make([]utils.StringPair, 0)
	//abbreviator.ParseConfigFile(expanderAbbrevations, &abbreviations)

	data := abbreviator.ParseDataFile(expanderAbbrevations)

	if len(data.AbbreviationRules) == 0 {
		return data, errors.New("No abbreviations rule found")
	}

	// This is where the magic happens
	err := abbreviator.AbbreviateExpressions(expressions, data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func printOutput(data *utils.ExpanderData, targetfile string) {

	// format output
	out := utils.MakeSortedString(data.GeneratedConfig)

	// print output
	fmt.Println("Generated Abbreviations:")
	fmt.Println(out)

	//utils.WriteToFile(out, targetfile)
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
		data, err := abbreviate(expressions)

		if err != nil {
			fmt.Println(err)
			return
		}

		//print output
		printOutput(data, generatedConfigFile)

	},
}

func init() {
	mapCmd.PersistentFlags().StringVar(&expanderAbbrevations, "abbreviations", "", "file containing the abbreviations to be applied")
	mapCmd.PersistentFlags().StringVar(&expanderGeneratedConf, "generated-config", "", "file to which generated conf should be written. Default is $EXPANDER_GENERATED_CONF")

	rootCmd.AddCommand(mapCmd)
}
