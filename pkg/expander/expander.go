package expander

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
)

func ParseConfigFile(configfile string, mapping map[string]string) {
	if configfile == "" {
		return
	}
	pairs, err := utils.ReadPairsFromFile(configfile)
	if err != nil {
		fmt.Printf("Couldn't read configfile %s\n", configfile)
		return
	}

	for _, p := range *pairs {
		mapping[p.Key] = p.Value
	}
}

func Expand(s string, mapping map[string]string) string {
	return mapping[s]
}
