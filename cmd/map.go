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
	err := ioutil.WriteFile(filename, []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

var expanderAbbrevations string
var expanderGeneratedConf string

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate abbreviations for available kubernetes contexts",
	Long:  "Map available kubernetes contexts, and print the abbreviations. These will be available for expansion in future runs.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// process pipe content here

		abbreviations := abbreviator.NewAbbreviator()
		abbreviations.ParseConfigFile(expanderAbbrevations)

		if abbreviations.IsEmptyMap() {
			fmt.Println("No mapping found, exiting")
			return
		}

		input, err := ParseInput(args)
		if err != nil {
			fmt.Printf("Invalid input, %s. Error is %s.", args, err)
		}

		// This is where the magic happens
		out := abbreviations.GenerateMappingString(input)

		configfile := expanderGeneratedConf
		if configfile == "" {
			configfile = os.Getenv("EXPANDER_GENERATED_CONF")
		}

		fmt.Println("Generated Abbreviations:")
		fmt.Println(out)

		if configfile == "" {
			fmt.Println("Mapping not saved. To save, use the --generated-config flag or set the EXPANDER_GENERATED_CONF env var.")

		} else {
			writeToFile(out, configfile)
			fmt.Printf("Mapping saved to %s", configfile)
		}
	},
}

func init() {
	mapCmd.PersistentFlags().StringVar(&expanderAbbrevations, "abbreviations", "", "file containing the abbreviations to be applied")
	mapCmd.PersistentFlags().StringVar(&expanderGeneratedConf, "generated-config", "", "file to which generated conf should be written. Default is $EXPANDER_GENERATED_CONF")

	rootCmd.AddCommand(mapCmd)
}
