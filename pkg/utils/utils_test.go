package utils

import (
	"testing"
)

var mm1 = map[string]string{"a": "apple", "p": "pear"}
var mm2 = map[string]string{"a": "apple"}
var mm3 = map[string]string{"b": "bear"}
var mm4 = map[string]string{"b": "bear", "p": "pear"}
var mm5 = map[string]string{"a": "bear", "p": "pear"}

func testComparemaps(t *testing.T) {

	var testcases = []struct {
		m1  map[string]string
		m2  map[string]string
		out bool
	}{{mm1, mm1, true},
		{mm1, mm2, false},
		{mm1, mm3, false},
		{mm2, mm3, false},
		{mm1, mm4, false},
	}
	for _, tc := range testcases {
		if comparemaps(tc.m1, tc.m2) != tc.out {
			t.Fatalf("comparing the equality of maps %s and %s should be %t but it isn't", tc.m1, tc.m2, tc.out)
		}
	}
}

func TestIsIdentical(t *testing.T) {
	base := ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		mm2,
	}
	differentAR1 := ExpanderData{
		[]map[string]string{mm1},
		mm1,
		mm2,
	}
	differentAR2 := ExpanderData{
		[]map[string]string{mm2, mm1},
		mm1,
		mm2,
	}
	differentGC := ExpanderData{
		[]map[string]string{mm2, mm1},
		mm3,
		mm2,
	}
	differentCC := ExpanderData{
		[]map[string]string{mm2, mm1},
		mm1,
		mm3,
	}

	if base.IsIdentical(&base) != true {
		t.Fatalf("Expander data is not identical to itself")
	}

	if base.IsIdentical(&differentAR1) {
		t.Fatalf("ExpanderData deemed identical, even though the AbbreviationRules lists have different length.")
	}

	if base.IsIdentical(&differentAR2) {
		t.Fatalf("ExpanderData deemed identical, even though the AbbreviationRules lists have different order.")
	}

	if base.IsIdentical(&differentGC) {
		t.Fatalf("ExpanderData deemed identical, even though the GeneratedConfig is a different map.")
	}

	if base.IsIdentical(&differentCC) {
		t.Fatalf("ExpanderData deemed identical, even though the CustomConfig is a different map.")
	}
}

func TestHasConfig(t *testing.T) {
	var e1 = ExpanderData{
		[]map[string]string{mm1, mm2},
		mm2,
		mm1,
	}
	var e2 = ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		mm1,
	}
	var e3 = ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		map[string]string{},
	}
	var e4 = ExpanderData{
		[]map[string]string{},
		mm1,
		mm2,
	}
	var e5 = ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		map[string]string{},
	}

	var testcases = []struct {
		e1  ExpanderData
		out bool
	}{{e1, true},
		{e2, true},
		{e3, true},
		{e4, true},
		{e5, false},
	}

	for _, tc := range testcases {
		if tc.e1.HasConfig() != tc.out {
			t.Fatalf("HasConfig for %s should be %t, nut it isn't.", tc.e1, tc.out)
		}
	}
}

func TestIsConsistent(t *testing.T) {
	var e1 = ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		map[string]string{},
	}
	var e2 = ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		mm1,
	}
	var e3 = ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		map[string]string{},
	}
	var e4 = ExpanderData{
		[]map[string]string{},
		mm1,
		mm2,
	}
	var e5 = ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		mm5,
	}

	var testcases = []struct {
		e1  ExpanderData
		out bool
	}{
		{e1, true},
		{e2, true},
		{e3, true},
		{e4, true},
		{e5, false},
	}

	for _, tc := range testcases {
		if tc.e1.IsConsistent() != tc.out {
			t.Fatalf("Consistency check failed for %s, should be %t but it isn't", tc.e1, tc.out)
		}
	}
}

func TestHasAbbreviationRules(t *testing.T) {
	var e1 = ExpanderData{
		[]map[string]string{mm1, mm2},
		mm1,
		mm5,
	}
	var e2 = ExpanderData{
		[]map[string]string{},
		mm1,
		mm5,
	}
	var e3 = ExpanderData{
		[]map[string]string{mm1, mm2},
		map[string]string{},
		map[string]string{},
	}
	var testcases = []struct {
		e1  ExpanderData
		out bool
	}{
		{e1, true},
		{e2, false},
		{e3, true},
	}
	for _, tc := range testcases {
		if tc.e1.HasAbbreviationRules() != tc.out {
			t.Fatalf("Abbreviation rules check failed for %s, should be %t but it isn't", tc.e1, tc.out)
		}
	}
}

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
