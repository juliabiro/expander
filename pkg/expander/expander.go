package expander

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
)

var mapping map[string]string

type Expander struct {
	mapping map[string]string
}

func NewExpander() *Expander {
	expander := Expander{}
	expander.mapping = make(map[string]string)

	return &expander

}

func (e *Expander) ParseConfigFile(configfile string) {
	if configfile == "" {
		return
	}
	pairs, err := utils.ReadPairsFromFile(configfile)
	if err != nil {
		fmt.Printf("Couldn't read configfile %s\n", configfile)
		return
	}

	for _, p := range *pairs {
		e.mapping[p.Key] = p.Value
	}
}

func (e *Expander) Expand(s string) string {
	return e.mapping[s]
}

func (e *Expander) IsEmptyMap() bool {
	return len(e.mapping) == 0
}
