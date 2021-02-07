package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ExpanderData struct {
	AbbreviationRules []map[string]string `json:"abbreviation_rules"`
	GeneratedConfig   map[string]string   `json:"generated_config"`
	CustomConfig      map[string]string   `json:"custom_config"`
}

func comparemaps(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		val, ok := m2[k]
		if !ok {
			return false
		}
		if v != val {
			return false
		}
	}
	return true
}
func (e1 *ExpanderData) IsIdentical(e2 *ExpanderData) bool {
	ret := true
	ret = ret && comparemaps(e1.GeneratedConfig, e2.GeneratedConfig)
	ret = ret && comparemaps(e1.CustomConfig, e2.CustomConfig)

	for i, _ := range e1.AbbreviationRules {
		ret = ret && comparemaps(e1.AbbreviationRules[i], e2.AbbreviationRules[i])
	}

	return ret

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
