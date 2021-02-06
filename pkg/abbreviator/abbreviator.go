package abbreviator

import (
	"errors"
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
	"strings"
)

func ParseConfigFile(configfile string, abbreviations *[]utils.StringPair) {
	pairs := utils.ReadPairsFromFile(configfile)
	if pairs != nil {
		*abbreviations = append(*abbreviations, (*pairs)...)
	}
}

func ParseDataFile(configfile string) *utils.ExpanderData {
	return utils.ReadDataFromFile(configfile)

}

func abbreviate(ctx string, abbreviation_mapping []map[string]string) string {
	res := strings.Repeat(ctx, 1)
	for _, m := range abbreviation_mapping {
		for k, v := range m {
			if k == "" {
				continue
			}
			res = strings.ReplaceAll(res, k, v)

		}
	}
	return res
}

func AbbreviateExpressions(expressions []string, data *utils.ExpanderData) error {
	for i, word := range expressions {
		abbr := abbreviate(word, data.AbbreviationRules)
		for _, word2 := range expressions[:i] {
			if abbr == abbreviate(word2, data.AbbreviationRules) {
				fmt.Printf("Both %s and %s abbreviate to %s. Aborting. Check you abbreviation mapping!", word, word2, abbr)
				return errors.New("Abbreviation collision")
			}
		}
		data.GeneratedConfig[abbr] = word
	}

	return nil
}
