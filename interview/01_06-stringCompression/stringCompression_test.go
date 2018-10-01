package main

import "testing"

func TestStringCompression(t *testing.T) {

	tables := []struct {
		in  string // the string to test against
		out string // the result if it is palindrome or not
	}{}

	for _, table := range tables {
		result := StringCompression(table.in)
		if result != table.out {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
