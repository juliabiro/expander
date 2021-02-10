package cmd

import (
	"errors"
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var configfile string
var configEnvvar string

func ParseInput(args []string) ([]string, error) {
	configEnvVar := os.Getenv("EXPANDER_CONFIG")

	if configEnvVar != "" {
		configfile = configEnvVar
	}

	if configfile == "" {
		return nil, errors.New("No configfile specified.")
	}

	// this may become more complex later
	return args, nil
}

func ParseConfigData(configfile string, validate func(*utils.ExpanderData) bool) *utils.ExpanderData {
	data := utils.ReadDataFromFile(configfile)
	if validate(data) {
		return data
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "expander",
	Short: "Expander expands abbreviations to longer strings based on a map.",
	Long:  `Expander can be used to expand abbreviations to longer strings based on a map. The intended use is to produce long argument strings for other applications`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.PersistentFlags().StringVar(&configfile, "config", "", "file containing the abbreviations mapping and the generated and custom configs")
}
