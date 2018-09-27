package main

import "testing"

// TestIsPermutation
// @param t, a pointer to a testing.T
func TestIsPermutation(t *testing.T) {
	// create a slice of structs
	// with each param and the result
	tables := []struct {
		x string
		y string
		r bool
	}{
		// instantiate copies of the struct with test data and known return type
		{x: "hello", y: "olleh", r: true},
		{x: "ooooo", y: "ooooo", r: true},
		{x: "opportunity", y: "golang", r: false},
		{x: "hello", y: "goodb", r: false},
		{x: "hello", y: "helll", r: false},
	}

	// loop over each struct and test it
	for _, table := range tables {
		result := IsPermutation(table.x, table.y)
		// check the result against the intended result
		if result != table.r {
			// if failing, print a testing Error
			t.Errorf("Result of (%v - %v) was incorrect, got: %v, want: %v.", table.x, table.y, table.r, result)
		}
	}
}
