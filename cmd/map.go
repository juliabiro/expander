package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/abbreviator"
	"github.com/juliabiro/expander/pkg/utils"
	"github.com/spf13/cobra"
)

var dryRun bool
var clear bool

func parseMapArguments(args []string) (input []string) {

	input, err := ParseInput(args)
	if err != nil {
		fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		return nil
	}
	return input
}

func printOutput(data *utils.ExpanderData, configfile string) {

	// format output
	out := utils.MakeSortedString(data.GeneratedConfig)

	// print output
	fmt.Println("Generated Abbreviations:")
	fmt.Println(out)

	if dryRun {
		fmt.Println("Mapping not saved. To save, use the --dry-run=false flag.")
	} else {
		utils.WriteToFile(data, configfile)
	}
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
		expressions := parseMapArguments(args)

		// get config
		data := ParseConfigData(configfile, abbreviator.ValidateData)
		if data == nil {
			return
		}

		//perform logic
		if clear == true {
			data.GeneratedConfig = make(map[string]string)
		}
		err := abbreviator.AbbreviateExpressions(expressions, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		//print output
		printOutput(data, configfile)

	},
}

func init() {
	mapCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", true, "toggles dry-run mode. When False, the generated abbreviations are saved to the config file. Default is true")
	mapCmd.PersistentFlags().BoolVar(&clear, "clear-existing-conf", false, "When true, the mapping replaces the exisiting generated conf. When false, it just adds to it (also overwrites it). Default is false")

	rootCmd.AddCommand(mapCmd)
}
