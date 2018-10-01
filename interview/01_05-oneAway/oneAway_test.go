package main

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	tables := []struct {
		in1 string
		in2 string
		out bool
	}{
		{in1: "pale", in2: "ple", out: true},
		{in1: "pales", in2: "pale", out: true},
		{in1: "pale", in2: "pole", out: true},
		{in1: "pale", in2: "bake", out: false},
	}

	for _, table := range tables {
		result := OneAway(table.in1, table.in2)
		if result != table.out {
			t.Errorf("Result of ('%v, %v') was incorrect, got: '%v', want: '%v'.", table.in1, table.in2, result, table.out)
		}
	}
}
