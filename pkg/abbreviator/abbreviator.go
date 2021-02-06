package abbreviator

import (
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

func AbbreviateExpressions(expressions []string, abbreviation_mapping []map[string]string) map[string]string {
	data := make(map[string]string)
	for i, word := range expressions {
		abbr := abbreviate(word, abbreviation_mapping)
		for _, word2 := range expressions[:i] {
			if abbr == abbreviate(word2, abbreviation_mapping) {
				fmt.Printf("Both %s and %s abbreviate to %s. Aborting. Check you abbreviation mapping!", word, word2, abbr)
				return nil
			}
		}
		data[abbr] = word
	}

	return data
}
