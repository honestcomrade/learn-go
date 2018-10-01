package main

import "testing"

func TestIsPermPalindrome(t *testing.T) {

	tables := []struct {
		in  string // the string to test against
		out bool   // the result if it is palindrome or not
	}{
		{in: "tact Coa", out: true},
		{in: "race Car", out: true},
		{in: "race Carl", out: false},
	}

	for _, table := range tables {
		result := IsPermPalindrome(table.in)
		if result != table.out {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
