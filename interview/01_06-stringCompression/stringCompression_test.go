package main

import "testing"

func TestStringCompression(t *testing.T) {

	tables := []struct {
		in  string // the string to test against
		out string // the result of the compression
	}{
		{"aabbaaacxxxxx", "a2b2a3c1x5"},
		{"a", "a"},
		{"mmmmnzz", "m4n1z2"},
	}

	for _, table := range tables {
		result := StringCompression(table.in)
		if result != table.out {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
