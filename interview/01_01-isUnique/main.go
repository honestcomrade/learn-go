package main

// IsUniqueChars ....the actual function
// first ask if the length of the string can be assumed to only include ascii characters
// if it is longer than the whole ascii chart, then it can not be unique
func IsUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}
	// create a slice of bool
	// and map to them the result of if they have been seen
	seen := make(map[rune]bool)
	// loop over the string
	for _, char := range str {
		if seen[char] {
			// if the seen array at this chars index is true
			return false
		}
		// mark the seen array at this index as seen
		seen[char] = true
	}
	// otherwise weve made it through entire array
	return true
}

func main() {}
