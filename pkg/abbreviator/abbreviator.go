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
		fmt.Println("Couldn't read configfile %s", configfile)
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

func (a *Abbreviator) GenerateMapping(expressions string) string {
	ctxMap := make(map[string]string)
	contextLines := strings.Split(string(expressions), " ")
	for _, line := range contextLines {
		ctx := strings.Split(strings.Trim(line, " "), " ")[0]
		abbr := a.abbreviate(ctx)
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

	return out
}

func (a *Abbreviator) IsEmptyMap() bool {
	return len(a.mapping) == 0
}
