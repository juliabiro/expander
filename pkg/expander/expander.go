package expander

import (
	"fmt"
	"github.com/juliabiro/expander/pkg/utils"
)

func ValidateData(data *utils.ExpanderData) bool {
	if !data.HasConfig() {
		fmt.Println("No config found in file.")
		return false
	}
	if data.IsConsistent() {
		return true
	} else {
		fmt.Printf("Config file contains inconsistent data, there are colliding keys in the generated and the Custom config parts. Check configfile.")
		return false
	}

}

func expand(s string, data *utils.ExpanderData) string {

	val, ok := data.GeneratedConfig[s]
	if ok {
		return val
	}

	val, ok = data.CustomConfig[s]
	if ok {
		return val
	}
	return ""
}

func ExpandExpressions(expressions []string, data *utils.ExpanderData) []string {
	ret := make([]string, 0)
	for _, c := range expressions {
		ret = append(ret, expand(c, data))
	}
	return ret
}
