package abbreviator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type stringPair struct {
	f string
	s string
}
type Abbreviator struct {
	configFile string
	mapping    []stringPair
}

func NewAbbreviator(file string) *Abbreviator {
	abbreviations := Abbreviator{}
	abbreviations.configFile = file
	abbreviations.mapping = make([]stringPair, 0)

	return &abbreviations

}

//TODO this part should be abstracted away, this is code duplication
func (a *Abbreviator) ParseConfigFile() error {
	data, err := ioutil.ReadFile(a.configFile)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		f, s := "", ""
		switch len(pairs) {
		case 0:
			continue
		case 1:
			f = strings.TrimSpace(pairs[0])
			s = ""
		default:
			f = strings.TrimSpace(pairs[0])
			s = strings.TrimSpace(pairs[1])
		}

		a.mapping = append(a.mapping, stringPair{f, s})
	}
	return nil
}

func (a *Abbreviator) abbreviate(ctx string) string {
	res := strings.Repeat(ctx, 1)
	for _, sp := range a.mapping {
		res = strings.ReplaceAll(res, sp.f, sp.s)
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
