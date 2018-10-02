package main

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	tables := []struct {
		str []string
		out bool
	}{
		{str: []string{"pale", "ple"}, out: true},
		{str: []string{"pales", "pale"}, out: true},
		{str: []string{"pale", "pole"}, out: true},
		{str: []string{"pale", "bake"}, out: false},
	}

	for _, table := range tables {
		result := OneAway(table.str)
		if result != table.out {
			t.Errorf("Result of ('%v, %v') was incorrect, got: '%v', want: '%v'.", table.str[0], table.str[1], result, table.out)
		}
	}
}
