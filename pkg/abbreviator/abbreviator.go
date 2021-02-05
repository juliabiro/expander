package abbreviator

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
	"strings"
)

func ParseConfigFile(configfile string) []utils.StringPair {
	if configfile == "" {
		return nil
	}
	pairs, err := utils.ReadPairsFromFile(configfile)
	if err != nil {
		fmt.Printf("Couldn't read configfile %s\n", configfile)
		return nil
	}
	return *pairs
}

func abbreviate(ctx string, abbreviation_mapping []utils.StringPair) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range abbreviation_mapping {
		res = strings.ReplaceAll(res, sp.Key, sp.Value)
	}
	return res
}

func makeSortedString(m map[string]string) string {
	out := ""
	for k, v := range m {
		if k == "" {
			continue
		}
		if k == v {
			continue
		}
		out = out + fmt.Sprintf("%s: %s\n", k, v)

	}

	return out
}

func AbbreviateExpressions(expressions []string, abbreviation_mapping []utils.StringPair) map[string]string {
	data := make(map[string]string)
	for _, word := range expressions {
		abbr := abbreviate(word, abbreviation_mapping)
		data[abbr] = word
	}

	return data
}
func GenerateMappingString(expressions []string, abbreviation_mapping []utils.StringPair) string {

	data := AbbreviateExpressions(expressions, abbreviation_mapping)
	return makeSortedString(data)
}
