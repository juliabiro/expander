package utils

import (
	"testing"
)

func TestReadPairsfromFile(t *testing.T) {
	abbreviations := *ReadPairsFromFile("../../example_mapping")
	expected := []StringPair{
		StringPair{"apple", "a"},
		StringPair{"pear", "p"},
		StringPair{"domain1.com", "d1"},
		StringPair{"domain2.com", "d2"},
		StringPair{"production", "p"},
		StringPair{"staging", "s"},
		StringPair{"-0", ""},
		StringPair{"-", ""},
	}

	if len(abbreviations) != len(expected) {
		t.Fatalf("reading in example_mapping, I didn't get the expected length of data")
		return
	}

	for i, _ := range abbreviations {
		if abbreviations[i] != expected[i] {
			t.Fatalf("Mismatch while reading in the example mapping. Ezpected %s but got %s", expected[i], abbreviations[i])
		}
	}
}

func TestEquals(t *testing.T) {
	var testCases = []struct {
		s1     StringPair
		s2     StringPair
		result bool
	}{
		{StringPair{"a", "b"}, StringPair{"a", "b"}, true},
		{StringPair{"c", "b"}, StringPair{"a", "b"}, false},
		{StringPair{"a", "b"}, StringPair{"c", "b"}, false},
		{StringPair{"a", "b"}, StringPair{"c", "d"}, false},
	}

	for _, tc := range testCases {
		if tc.s1.Equals(&tc.s2) != tc.result {
			t.Fatalf("%s == %s should be %t, but it isn't", tc.s1, tc.s2, tc.result)
		}
	}

	if (&StringPair{"a", "b"}).Equals(nil) != false {
		t.Fatalf("Comparing with nil should be false but it isnt")
	}

}

func TestProcessPair(t *testing.T) {

	var testCases = []struct {
		input  []string
		output *StringPair
	}{
		{[]string{"a", "b"}, &StringPair{"a", "b"}},
		{[]string{"a", ""}, &StringPair{"a", ""}},
		{[]string{"", "b"}, nil},
		{[]string{" a ", "b"}, &StringPair{"a", "b"}},
		{[]string{"a", " b "}, &StringPair{"a", "b"}},
		{[]string{"a", " b ", "c"}, &StringPair{"a", "b"}},
	}

	for _, tc := range testCases {
		o := processPair(tc.input)
		if o == nil && tc.output != nil {
			t.Fatalf("Error processing input %s, expected %s but got %s", tc.input, tc.output, o)
		}
		if o != nil && !o.Equals(tc.output) {
			t.Fatalf("Error processing input %s, expected %s but got %s", tc.input, tc.output, o)
		}
	}
}
