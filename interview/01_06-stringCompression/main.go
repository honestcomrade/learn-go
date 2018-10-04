package main

import "fmt"

// StringCompression ...
// @param input string
// returns string
// given an input string
// write an algorithm which compresses the string
// by converting sequential repeats of characters
// into the count of repeats
// if the "compressed" string would be longer
// than the input string, return the input string
// assume the string has only lower case (a-z)
func StringCompression(x string) string {
	replaceCount := 0
	chars := []rune(x) // format the input string into a slice of runes
	var retStr []rune  //  make an output slice of runes to return (as a string later)
	for i := 0; i < len(chars); i++ {
		// early return for if we ever make a string
		// that would be longer than the input
		if len(retStr) > i {
			return string(chars)
		}
		if i == 0 {
			// chars[0] always goes in
			retStr = append(retStr, chars[i])
			// so you can skip to next iteration
			continue
		}
		// if the char is the same as the one before it
		// that means we need to count the replacement ticker up
		if chars[i-1] == chars[i] {
			replaceCount++
			// special case for last char always need to insert replaceCount
			if i == len(chars)-1 {
				retStr = append(retStr, rune(replaceCount+1)+'0') // convert the replaceCount int to it's ascii equiv rune
			}
		} else {
			// otherwise since we are onto a new char we need to append
			// the old char and the replaceCount
			retStr = append(retStr, rune(replaceCount+1)+'0')
			retStr = append(retStr, chars[i])
			replaceCount = 0
		}
	}
	return string(retStr)
}

func main() {
	fmt.Println(StringCompression("aaaabbaaa"))
}
