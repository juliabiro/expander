package expander

import (
	"github.com/juliabiro/expander/pkg/utils"
	"testing"
)

var mm1 = map[string]string{"a": "apple", "p": "pear"}
var mm2 = map[string]string{"a": "apple"}
var mm3 = map[string]string{"b": "bear"}
var mm4 = map[string]string{"b": "bear", "p": "pear"}
var mm5 = map[string]string{"a": "bear", "p": "pear"}

func TestValidateData(t *testing.T) {
	var e1 = utils.ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		map[string]string{},
	}
	var e2 = utils.ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		mm1,
	}
	var e3 = utils.ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		map[string]string{},
	}
	var e4 = utils.ExpanderData{
		[]map[string]string{},
		mm1,
		mm2,
	}
	var e5 = utils.ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		mm5,
	}

	var testcases = []struct {
		e1  utils.ExpanderData
		out bool
	}{
		{e1, true},
		{e2, true},
		{e3, false},
		{e4, true},
		{e5, false},
	}

	for _, tc := range testcases {
		if ValidateData(&tc.e1) != tc.out {
			t.Fatalf("Validation failed for exapnding: %s validity should be %t but it isn't", tc.e1, tc.out)
		}
	}

}
