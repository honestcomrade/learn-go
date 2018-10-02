package main

import (
	"fmt"
	"sort"
)

// byLength is a custom type that we can add a sort function to
type byLength []string

// need a Len function on the byLength type, for sorting
func (s byLength) Len() int {
	return len(s)
}

// need a Swap function on the byLength type
// for swapping positions in the sort
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// need a Comparator function on the byLength
//type to determine which positions to Swap in the sort
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// OneAway ...
// @param s []string
// returns bool
// given two input strings (in a slice, for sorting purposes), write an algorithm
// to determine if the two strings
// differ by 1 or 0 edits
func OneAway(s []string) bool {
	sort.Sort(byLength(s))
	// we need to sort so that we don't range over the shorter string
	// which would cuase out of index error
	// x is shorter
	x := s[0]
	y := s[1]
	// if the difference in char length between x and y is greater than 1
	// no way they can be one char away, return false early(short circuit)
	if len(x)-len(y) < -1 || len(x)-len(y) > 1 {
		return false
	}

	// make a map of the counts of each char in each string
	xChars := make(map[rune]int)
	yChars := make(map[rune]int)

	diffCount := 0
	for _, v := range x {
		xChars[v]++
	}

	for _, v := range y {
		yChars[v]++
		// if a char in y is not present in x, tick up the counter
		if xChars[v] == 0 {
			diffCount++
		}
	}

	// if that counter got higher than 1, return false
	if diffCount > 1 {
		return false
	}
	return true
}

func main() {
	strings := byLength{"pale", "palle"}
	fmt.Println(OneAway(strings))
}
