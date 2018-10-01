package main

import (
	"fmt"
	"strings"
)

// IsPermPalindrome ...
// @param s string - the input to test
// @returns bool - the result of the test
// determines if a string is a permutation of a palindrome
// a palindrome is a string that is the same backwards as it is forwards
// a permutation of it must contain all the same letters but in any order
func IsPermPalindrome(s string) bool {
	charCounts := make(map[rune]int)
	lower := strings.ToLower(s)
	for _, char := range lower {
		fmt.Printf("%c\n", char)
		// ignore whitespace characters for counts
		if char != 32 {
			charCounts[char]++
		}
	}

	countOfOdds := 0
	for _, count := range charCounts {
		if count != 32 {
			if !determineEvenInt(count) {
				countOfOdds++
			}
			// if there are more than one character with an odd count
			// false
			if countOfOdds > 1 {
				return false
			}
		}
	}
	return true
}

func determineEvenInt(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func main() {
	str := "tact Coa"
	fmt.Println(IsPermPalindrome(str))
}
