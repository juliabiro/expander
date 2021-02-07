package expander

import (
	"github.com/juliabiro/expander/pkg/utils"
)

func ParseConfigData(configfile string) *utils.ExpanderData {
	// I want to unite the 2 maps. Since I am not going to write them back, it is Ok to modify
	data := utils.ReadDataFromFile(configfile)
	// TODO validate config file
	return data
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
