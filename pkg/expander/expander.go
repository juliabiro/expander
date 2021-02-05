package expander

import (
	"github.com/juliabiro/expander/pkg/utils"
)

func ParseConfigFile(configfile string, mapping map[string]string) {
	pairs := utils.ReadPairsFromFile(configfile)

	if pairs != nil {
		for _, p := range *pairs {
			mapping[p.Key] = p.Value
		}
	}

}

func expand(s string, mapping map[string]string) string {
	return mapping[s]
}

func ExpandExpressions(expressions []string, abbreviations map[string]string) []string {
	ret := make([]string, 0)
	for _, c := range expressions {
		ret = append(ret, expand(c, abbreviations))
	}
	return ret
}
