package expander

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Printf("Failed to open configfile %s, error is %s.", configfile, err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		if len(pairs) < 2 {
			continue
		}
		e.mapping[pairs[0]] = strings.TrimSpace(pairs[1])
	}
}

func (e *Expander) Expand(s string) string {
	return e.mapping[s]
}

func (e *Expander) IsEmptyMap() bool {
	return len(e.mapping) == 0
}
