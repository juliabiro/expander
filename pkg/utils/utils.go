package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type StringPair struct {
	Key   string
	Value string
}

func (s *StringPair) Equals(s2 *StringPair) bool {
	if s2 == nil {
		return false
	}
	return s.Key == s2.Key && s.Value == s2.Value
}

func processPair(pairs []string) *StringPair {
	f, s := "", ""
	switch len(pairs) {
	case 0:
		return nil
	case 1:
		f = strings.TrimSpace(pairs[0])
		s = ""
	default:
		f = strings.TrimSpace(pairs[0])
		s = strings.TrimSpace(pairs[1])
	}

	if f == "" {
		return nil
	}
	return &StringPair{f, s}
}

type ExpanderData struct {
	AbbreviationRules []map[string]string `json:"abbreviation_rules"`
	GeneratedConfig   map[string]string   `json:"generated_config"`
	CustomConfig      map[string]string   `json:"custom_config"`
}

func ReadDataFromFile(file string) *ExpanderData {
	if file == "" {
		return nil
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to open configfile %s, error is %s.", file, err)
		return nil
	}

	ed := ExpanderData{}
	json.Unmarshal(data, &ed)
	return &ed
}
func ReadPairsFromFile(file string) *[]StringPair {
	if file == "" {
		return nil
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to open configfile %s, error is %s.", file, err)
		return nil
	}

	mapping := make([]StringPair, 0)
	for _, line := range strings.Split(string(data), "\n") {
		pairs := strings.Split(line, ":")
		p := processPair(pairs)

		if p != nil {
			mapping = append(mapping, *p)
		}
	}
	return &mapping

}

func WriteToFile(out string, filename string) {
	if filename == "" {
		fmt.Println("Mapping not saved. To save, use the --generated-config flag or set the EXPANDER_GENERATED_CONF env var.")

	} else {
		err := ioutil.WriteFile(filename, []byte(out), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Mapping saved to %s", filename)
	}
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
