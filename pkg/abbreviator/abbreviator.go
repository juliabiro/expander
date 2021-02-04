package abbreviator

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
	"strings"
)

type Abbreviator struct {
	mapping []utils.StringPair
}

func NewAbbreviator() *Abbreviator {
	abbreviations := Abbreviator{}
	return &abbreviations
}

func (a *Abbreviator) ParseConfigFile(configfile string) {
	if configfile == "" {
		return
	}
	pairs, err := utils.ReadPairsFromFile(configfile)
	if err != nil {
		fmt.Printf("Couldn't read configfile %s\n", configfile)
		return
	}
	a.mapping = *pairs
}

func (a *Abbreviator) abbreviate(ctx string) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range a.mapping {
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

func (a *Abbreviator) GenerateMappingString(expressions []string) string {
	data := make(map[string]string)
	for _, word := range expressions {
		abbr := a.abbreviate(word)
		data[abbr] = word
	}

	return makeSortedString(data)
}

func (a *Abbreviator) IsEmptyMap() bool {
	return len(a.mapping) == 0
}
