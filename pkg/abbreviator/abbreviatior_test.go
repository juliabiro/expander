package abbreviator

import (
	"github.com/juliabiro/expander/pkg/utils"
	"testing"
)

func TestAbbreviate(t *testing.T) {
	var testCases = []struct {
		input         string
		abbreviations []map[string]string
		output        string
	}{
		{"apple", []map[string]string{{"apple": "a"}}, "a"},
		{"pear", []map[string]string{{"apple": "a"}}, "pear"},
		{"apple-pear", []map[string]string{{"apple": "a"}, {"pear": "p"}}, "a-p"},
		{"apple-pear", []map[string]string{{"pear": "p"}, {"apple": "a"}}, "a-p"},
		{"apple-pear", []map[string]string{{"-": ""}}, "applepear"},
		{"apple-00pear", []map[string]string{{"-": ""}, {"-0": ""}}, "apple00pear"},
		{"apple-00pear", []map[string]string{{"-0": ""}, {"-": ""}}, "apple0pear"},
		{"apple", []map[string]string{{"": "a"}}, "apple"},
	}

	for _, tc := range testCases {
		o := abbreviate(tc.input, tc.abbreviations)
		if o != tc.output {
			t.Fatalf("Abbreviation mismatch. Input: %s, abbreviations: %s, expected %s but got %s", tc.input, tc.abbreviations, tc.output, o)
		}
	}
}

func TestAbbreviateExpressionsNoRepeats(t *testing.T) {
	input := []string{"apple-pear", "apple0pear"}
	m := []map[string]string{{"-": ""}, {"0": ""}}
	ed := utils.ExpanderData{m, map[string]string{}, map[string]string{}}
	err := AbbreviateExpressions(input, &ed)

	if err == nil {
		t.Fatalf("Expected to abort on repeated abbreviation, but it went forward.")
	}

}
