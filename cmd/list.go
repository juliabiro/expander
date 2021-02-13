package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/expander"
	"github.com/juliabiro/expander/pkg/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured abbreviations.",
	Long:  "List configured abbreviations, both generated and custom.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, err := ParseInput(args)

		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		// perform logic
		data := ParseConfigData(configfile, expander.ValidateData)
		if data == nil {
			return
		}
		generated := utils.MakeSortedString(data.GeneratedConfig)
		fmt.Println("Generated Abbreviations:")
		fmt.Println(generated)
		custom := utils.MakeSortedString(data.CustomConfig)
		fmt.Println("Custom Abbreviations:")
		fmt.Println(custom)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
