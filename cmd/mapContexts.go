package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type stringPair struct {
	f string
	s string
}
type Abbreviations struct {
	configFile string
	mapping    []stringPair
}

func NewAbbreviations(file string) *Abbreviations {
	abbreviations := Abbreviations{}
	abbreviations.configFile = file
	abbreviations.mapping = make([]stringPair, 0)

	return &abbreviations

}

//TODO this part should be abstracted away, this is code duplication
func (a *Abbreviations) ParseConfigFile() error {
	data, err := ioutil.ReadFile(a.configFile)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		f, s := "", ""
		switch len(pairs) {
		case 0:
			continue
		case 1:
			f = strings.TrimSpace(pairs[0])
			s = ""
		default:
			f = strings.TrimSpace(pairs[0])
			s = strings.TrimSpace(pairs[1])
		}

		a.mapping = append(a.mapping, stringPair{f, s})
	}
	return nil
}

func abbreviate(ctx string, a Abbreviations) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range a.mapping {
		res = strings.ReplaceAll(res, sp.f, sp.s)
	}
	return res
}

var longExpressions string
var expanderAbbrevations string

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate abbreviations for available kubernetes contexts",
	Long:  "Map available kubernetes contexts, and print the abbreviations. These will be available for expansion in future runs.",
	Run: func(cmd *cobra.Command, args []string) {
		// process pipe content here

		abbreviations := NewAbbreviations(expanderAbbrevations)
		err := abbreviations.ParseConfigFile()

		if err != nil {
			log.Fatal(err)
		}

		ctxMap := make(map[string]string)
		contextLines := strings.Split(string(longExpressions), " ")
		for _, line := range contextLines {
			ctx := strings.Split(strings.Trim(line, " "), " ")[0]
			abbr := abbreviate(ctx, *abbreviations)
			ctxMap[abbr] = ctx
		}
		// TODO: sort and clean
		out := ""
		for k, v := range ctxMap {
			if k == "" {
				continue
			}
			if k == v {
				continue
			}
			out = out + fmt.Sprintf("%s: %s\n", k, v)

		}

		// TODO: write to file
		configfile := os.Getenv("EXPANDER_CONF")
		err = ioutil.WriteFile(configfile, []byte(out), 0444)

		fmt.Println("Generated Abbreviations:")
		fmt.Println(out)
		fmt.Printf("Mapping saved to %s", configfile)
	},
}

func init() {
	mapCmd.PersistentFlags().StringVar(&longExpressions, "expressions", "", "space separated values of long strings that need to be abbreviated")
	mapCmd.PersistentFlags().StringVar(&expanderAbbrevations, "abbreviations", "", "file containing the abbreviations to be applied")

	rootCmd.AddCommand(mapCmd)
}
