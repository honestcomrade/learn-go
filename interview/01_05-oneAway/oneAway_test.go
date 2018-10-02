package main

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	tables := []struct {
		str []string
		out bool
	}{
		{[]string{"a", " "}, true},
		{[]string{"bb", "aa"}, false},
		{[]string{"pale", "ple"}, true},
		{[]string{"pales", "pale"}, true},
		{[]string{"pale", "pole"}, true},
		{[]string{"pale", "bake"}, false},
	}

	for _, table := range tables {
		result := OneAway(table.str)
		if result != table.out {
			t.Errorf("Result of ('%v, %v') was incorrect, got: '%v', want: '%v'.", table.str[0], table.str[1], result, table.out)
		}
	}
}
