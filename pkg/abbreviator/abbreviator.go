package abbreviator

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
	"strings"
)

func ParseConfigFile(configfile string, abbreviations *[]utils.StringPair) {
	*abbreviations = append(*abbreviations, *(utils.ReadPairsFromFile(configfile))...)
}

func abbreviate(ctx string, abbreviation_mapping []utils.StringPair) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range abbreviation_mapping {
		if sp.Key == "" {
			continue
		}
		res = strings.ReplaceAll(res, sp.Key, sp.Value)
	}
	return res
}

func AbbreviateExpressions(expressions []string, abbreviation_mapping []utils.StringPair) map[string]string {
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
