package cmd

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/abbreviator"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

func writeToFile(out string, filename string) {
	err := ioutil.WriteFile(filename, []byte(out), 0444)
	if err != nil {
		log.Fatal(err)
	}
}

var longExpressions string
var expanderAbbrevations string

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate abbreviations for available kubernetes contexts",
	Long:  "Map available kubernetes contexts, and print the abbreviations. These will be available for expansion in future runs.",
	Run: func(cmd *cobra.Command, args []string) {
		// process pipe content here

		abbreviations := abbreviator.NewAbbreviator(expanderAbbrevations)
		err := abbreviations.ParseConfigFile()

		if err != nil {
			log.Fatal(err)
		}

		out := abbreviations.GenerateMapping(longExpressions)

		// TODO: write to file
		configfile := os.Getenv("EXPANDER_CONF")

		fmt.Println("Generated Abbreviations:")
		fmt.Println(out)

		if configfile == "" {
			fmt.Println("Mapping not saved. To save, set the EXPANDER_CONF env var")

		} else {
			writeToFile(out, configfile)
			fmt.Printf("Mapping saved to %s", configfile)
		}
	},
}

func init() {
	mapCmd.PersistentFlags().StringVar(&longExpressions, "expressions", "", "space separated values of long strings that need to be abbreviated")
	mapCmd.PersistentFlags().StringVar(&expanderAbbrevations, "abbreviations", "", "file containing the abbreviations to be applied")

	rootCmd.AddCommand(mapCmd)
}
