package main

import (
	"testing"
)

func TestStringRotation(t *testing.T) {

	tables := []struct {
		in1 string // the string to test against
		in2 string // the "rotated" string
		out bool   // the result of the test
	}{
		{
			in1: "erbottlewat",
			in2: "waterbottle",
			out: true,
		},
		{
			in1: "waterbottle",
			in2: "erbottlewat",
			out: true,
		},
		{
			in1: "wawaterbottle",
			in2: "erbottlewat",
			out: false,
		},
		{
			in1: "atwatw",
			in2: "watwat",
			out: true,
		},
		{
			in1: "iln",
			in2: "lin",
			out: false,
		},
	}

	for _, table := range tables {
		result := IsStringRotation(table.in1, table.in2)
		if result != table.out {
			t.Errorf("Result of ('%v:%v') was incorrect, got: '%v', want: '%v'.", table.in1, table.in2, result, table.out)
		}
	}

}
