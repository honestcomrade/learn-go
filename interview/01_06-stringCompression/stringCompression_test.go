package main

import "testing"

func TestStringCompression(t *testing.T) {

	tables := []struct {
		in  string // the string to test against
		out string // the result of the compression
	}{}

	for _, table := range tables {
		result := StringCompression(table.in)
		if result != table.out {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
