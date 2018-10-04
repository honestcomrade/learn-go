package main

import (
	"fmt"
	"strings"
)

func isSubstring(s1, s2 string) bool {
	if strings.Contains(s1, s2) {
		return true
	}
	return false
}

// IsStringRotation ...
// Assume you have a method isSubstring which checks if one word is a substring of another.
// Given two strings, s1, and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring
// (e.g. "waterbottle" IS a rotation of "erbottlewat")
func IsStringRotation(s1, s2 string) bool {
	sLen := len(s1)
	// if length of the strings is different, false
	// if length is 0, false
	if sLen == 0 && sLen != len(s2) {
		return false
	}
	// concatenate s1 with itself
	// because then it will have to contain
	// all of s2 for isSubString to return true
	s1s1 := s1 + s1
	return isSubstring(s1s1, s2)
}

func main() {
	fmt.Println(IsStringRotation("erbottlewat", "waterbottle"))
}
