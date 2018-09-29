package main

import "testing"

// TestIsPermutation
// @param t, a pointer to a testing.T
func TestURLEscape(t *testing.T) {
	// create a slice of structs
	// with each param and the result
	tables := []struct {
		in  string // the provided string to escape
		len int    // the given "true length"
		out string // the desired return
	}{
		// instantiate copies of the struct with test data and known return type
		{in: "Mr John Smith    ", len: 13, out: "Mr%20John%20Smith"},
		{in: "Dr Kamilan  ", len: 10, out: "Dr%20Kamilan"},
		{in: "A longer string to test        ", len: 23, out: "A%20longer%20string%20to%20test"},
	}

	// loop over each struct and test it
	for _, table := range tables {
		result := URLEscape(table.in, table.len)
		// check the result against the intended result
		if result != table.out {
			// if failing, print a testing Error
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}
}
