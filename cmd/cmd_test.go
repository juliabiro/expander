package cmd

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	var testcases = []struct {
		in  []string
		out []string
	}{
		{[]string{"a"}, []string{"a"}},
		{[]string{"ab"}, []string{"ab"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{}, []string{}},
	}

	for _, c := range testcases {
		o, err := ParseInput(c.in)
		if err != nil {
			panic(err)
		}

		if len(o) != len(c.out) {
			t.Fatalf("InputParsing failed for %s. Output length doesn't match", c)
		}
		for i, _ := range o {
			if o[i] != c.out[i] {
				t.Fatalf("InputParsing failed for %s. Element %d doesn't match: expected: %s, got %s.", c, i, c.out[i], o[i])
			}
		}
	}
}
