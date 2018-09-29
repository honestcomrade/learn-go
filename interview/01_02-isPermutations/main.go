package main

// IsPermutation ...
// @param s1 type string
// @param s2 type string
// @returns bool
func IsPermutation(s1 string, s2 string) bool {
	// if the lengths are different, they can't be permutations
	if len(s1) != len(s2) {
		return false
	}
	// make a new map of rune:int pairs
	counts := make(map[rune]int)
	// loop over string 1
	// and tick up the count at each unique char found
	for _, char := range s1 {
		counts[char]++
	}
	// loop over string 2 and tick down the count at each rune found
	// we have seen it again!
	// this way when the second loop is over, each char in the first string and the second string
	// should have a count of 0
	// if it's higher or lower, that means we saw it too many or too few times
	// in the second string compared to the first
	for _, char := range s2 {
		counts[char]--
	}
	// loop over the whole map
	// if any of the counts at any rune are < 0 || > 0
	// then it is not a permutation
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {}
