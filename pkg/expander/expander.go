package expander

import (
	"io/ioutil"
	"strings"
)

var mapping map[string]string

type Expander struct {
	configFile string
	mapping    map[string]string
}

func NewExpander(file string) *Expander {
	expander := Expander{}
	expander.configFile = file
	expander.mapping = make(map[string]string)

	return &expander

}

func (e *Expander) ParseConfigFile() error {
	data, err := ioutil.ReadFile(e.configFile)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		e.mapping[pairs[0]] = strings.TrimSpace(pairs[1])
	}
	return nil
}

func (e *Expander) Expand(s string) string {
	return e.mapping[s]
}
