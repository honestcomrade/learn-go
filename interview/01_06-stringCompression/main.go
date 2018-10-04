package main

import (
	"fmt"
)

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
func StringCompression(x string) []byte {
	replaceCount := 0
	chars := []byte(x) // format the input string into a slice of bytes
	var retStr []byte  //  make an output slice of bytes to return (as a string later)
	// retStr = append(retStr, chars[0])
	for i := 0; i < len(chars); i++ {
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
			// special case for last char, need to append replacecount right away
			if i == len(chars)-1 {
				retStr = append(retStr, byte(replaceCount+1))
			}
		} else {
			// otherwise since we are onto a new char we need to append
			// the old char and the replacecount
			retStr = append(retStr, byte(replaceCount+1))
			retStr = append(retStr, chars[i])
			replaceCount = 0
		}
	}
	return retStr
}

func main() {
	str := StringCompression("aaaabbaaa")
	for _, char := range str {
		fmt.Printf("%c", char)
	}
	// fmt.Println(str)
}
