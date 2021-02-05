package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type StringPair struct {
	Key   string
	Value string
}

func ReadPairsFromFile(file string) *[]StringPair {
	mapping := make([]StringPair, 0)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to open configfile %s, error is %s.", file, err)
		return nil
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

		mapping = append(mapping, StringPair{f, s})
	}
	return &mapping

}
func MakeSortedString(m map[string]string) string {
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
