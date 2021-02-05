package abbreviator

import (
	"github.com/juliabiro/expander/pkg/utils"
	"strings"
)

func ParseConfigFile(configfile string, abbreviations []utils.StringPair) {
	abbreviations = *utils.ReadPairsFromFile(configfile)
}

func abbreviate(ctx string, abbreviation_mapping []utils.StringPair) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range abbreviation_mapping {
		res = strings.ReplaceAll(res, sp.Key, sp.Value)
	}
	return res
}

func AbbreviateExpressions(expressions []string, abbreviation_mapping []utils.StringPair) map[string]string {
	data := make(map[string]string)
	for _, word := range expressions {
		abbr := abbreviate(word, abbreviation_mapping)
		data[abbr] = word
	}

	return data
}
