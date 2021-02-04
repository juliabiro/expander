package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/spf13/cobra"
	"os"
)

var generatedConfig string
var customConfig string
var expandCmd = &cobra.Command{
	Use:   "ex",
	Short: "Expand known abbreviations",
	Long:  "Expand Abbreviation",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		generatedConfigFile := generatedConfig
		customConfigFile := customConfig

		if generatedConfigFile == "" {
			generatedConfigFile = os.Getenv("EXPANDER_GENERATED_CONF")
		}
		if customConfigFile == "" {
			customConfigFile = os.Getenv("EXPANDER_CUSTOM_CONF")
		}

		expander := expander.NewExpander()
		expander.ParseConfigFile(generatedConfigFile)
		expander.ParseConfigFile(customConfigFile)

		if expander.IsEmptyMap() {
			fmt.Println("No mapping found, exiting")
			return
		}

		input, err := ParseInput(args)

		if err != nil {
			fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		}
		// This is where the magic happens
		for _, c := range input {
			fmt.Print(expander.Expand(c))
		}

	},
}

func init() {
	expandCmd.PersistentFlags().StringVar(&generatedConfig, "generated-config", "", "path to the generated config file to use for expandsion. Default is the EXPANDER_GENERATED_CONF env var. ")
	expandCmd.PersistentFlags().StringVar(&customConfig, "custom-config", "", "path to the custom config file to use for expandsion. Default is the EXPANDER_CUSTOM_CONF env var. ")
	rootCmd.AddCommand(expandCmd)
}
