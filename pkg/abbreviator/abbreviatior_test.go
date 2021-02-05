package abbreviator

import (
	"github.com/juliabiro/expander/pkg/utils"
	"testing"
)

func TestAbbreviate(t *testing.T) {
	var testCases = []struct {
		input         string
		abbreviations []utils.StringPair
		output        string
	}{
		{"apple", []utils.StringPair{utils.StringPair{"apple", "a"}}, "a"},
		{"pear", []utils.StringPair{utils.StringPair{"apple", "a"}}, "pear"},
		{"apple-pear", []utils.StringPair{utils.StringPair{"apple", "a"}, utils.StringPair{"pear", "p"}}, "a-p"},
		{"apple-pear", []utils.StringPair{utils.StringPair{"pear", "p"}, utils.StringPair{"apple", "a"}}, "a-p"},
		{"apple-pear", []utils.StringPair{utils.StringPair{"-", ""}}, "applepear"},
		{"apple-00pear", []utils.StringPair{utils.StringPair{"-", ""}, utils.StringPair{"-0", ""}}, "apple00pear"},
		{"apple-00pear", []utils.StringPair{utils.StringPair{"-0", ""}, utils.StringPair{"-", ""}}, "apple0pear"},
		{"apple", []utils.StringPair{utils.StringPair{"", "a"}}, "apple"},
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
	m := []utils.StringPair{utils.StringPair{"-", ""}, utils.StringPair{"0", ""}}
	output := AbbreviateExpressions(input, m)

	if output != nil {
		t.Fatalf("Expected to abort on repeated abbreviation, but it went forward.")
	}

}

func TestParseConfigFile(t *testing.T) {
	abbreviations := make([]utils.StringPair, 0)
	ParseConfigFile("../../example_mapping", &abbreviations)

	expected := []utils.StringPair{
		utils.StringPair{"apple", "a"},
		utils.StringPair{"pear", "p"},
		utils.StringPair{"domain1.com", "d1"},
		utils.StringPair{"domain2.com", "d2"},
		utils.StringPair{"production", "p"},
		utils.StringPair{"staging", "s"},
		utils.StringPair{"-0", ""},
		utils.StringPair{"-", ""},
	}

	if len(abbreviations) != len(expected) {
		t.Fatalf("%s", abbreviations)
		t.Fatalf("Parsing example_mapping, I didn't get the expected length of data")
		return
	}

	for i, _ := range abbreviations {
		if abbreviations[i] != expected[i] {
			t.Fatalf("Mismatch while reading in the example mapping. Ezpected %s but got %s", expected[i], abbreviations[i])
		}
	}
}
