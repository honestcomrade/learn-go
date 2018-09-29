package main

import "fmt"

// URLEscape ...
// @param x string - the string to escape
// @param escapedLen int - assume it is provided
// @return string
// Will escape all spaces in an input string
// with URL friendly %20
// and remove trailing whitespace if exists
func URLEscape(x string, providedLen int) string {
	chars := []rune(x)
	run := len(x) - 1
	for i := providedLen - 1; i >= 0; i-- { // start at provided end ("true length") of str
		// find the ascii code of each rune
		if chars[i] == 32 {
			// if it's i points to 32, then it is whitespace
			// replace the char at that index and it's two preceding neighbors with
			// the right sequence to become the %20
			chars[run] = '0'
			chars[run-1] = '2'
			chars[run-2] = '%'
			// then move run pointer ahead 3 places
			// to move to the next original char
			run -= 3
		} else {
			// otherwise if it is not a whitespace,
			fmt.Printf("'%c' is being replaced by '%c'\n", chars[run], chars[i])
			// set the char at run pointer to that which was at i pointer
			// we need to push the chars to the end to make room for the '%20' insertions
			chars[run] = chars[i]
			fmt.Printf("position %v is now '%c'\n", run, chars[run])
			// move run pointer to next char
			run--
		}
	}
	return string(chars)
}

func main() {
	str := "Mr John Smith    "
	fmt.Println(str)
	fmt.Printf("'%v'", URLEscape(str, 13))
}
