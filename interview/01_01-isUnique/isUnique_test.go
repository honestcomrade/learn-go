package main

import "testing"

func TestIsUniqueChars(t *testing.T) {
	tables := []struct {
		x string
		r bool
	}{
		{x: "hello", r: false},
		{x: "normal", r: true},
		{x: "opportunity", r: false},
		{x: "abcdef", r: true},
	}

	for _, table := range tables {
		result := IsUniqueChars(table.x)
		if result != table.r {
			t.Errorf("Result of (%v) was incorrect, got: %v, want: %v.", table.x, table.r, result)
		}
	}

}
