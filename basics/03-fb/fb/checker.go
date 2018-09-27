package fb

import (
	"fmt"
)

func Stringchecker(slice []string, num int) {
	if num%3 == 0 && num%5 > 0 {
		fizz := slice[0:1]
		fmt.Printf("%v\n", fizz[0])
	} else if num%5 == 0 && num%3 > 0 {
		buzz := slice[1:2]
		fmt.Printf("%v\n", buzz[0])
	} else if num%5 == 0 && num%3 == 0 {
		fmt.Printf("%v\n", slice[0]+slice[1])
	} else {
		fmt.Println(num)
	}
}
