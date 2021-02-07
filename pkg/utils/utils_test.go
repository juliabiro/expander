package utils

import (
	"testing"
)

// TODO test comapremaps and IsIdentical

func TestReadDataFromFile(t *testing.T) {
	data := *ReadDataFromFile("../../example_conf.json")
	expected := ExpanderData{
		[]map[string]string{
			{"apple": "a"},
			{"pear": "p"},
			{"domain1.com": "d1"},
			{"domain2.com": "d2"},
			{"production": "p"},
			{"staging": "s"},
			{"-0": ""},
			{"-": ""},
		},
		map[string]string{"a23z": "apple-23-z"},
		map[string]string{},
	}

	if !expected.IsIdentical(&data) {
		t.Fatalf("Mismatch while reading in the example mapping. Epected %s but got %s", expected, data)
	}

	if ReadDataFromFile("") != nil {
		t.Fatalf("Shouldn't try to read empty filename, bit it does")
	}
}
