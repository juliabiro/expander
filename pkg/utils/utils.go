package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

type StringPair struct {
	Key   string
	Value string
}

func ReadPairsFromFile(file string) (*[]StringPair, error) {
	mapping := make([]StringPair, 0)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to open configfile %s, error is %s.", file, err)
		return nil, err
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
	return &mapping, nil

}
