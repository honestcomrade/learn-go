package main

import (
	"fmt"
)

func main() {
	for out := range fb(100) {
		fmt.Println(out)
	}
}

func fb(j int) <-chan string {
	out := make(chan string, j)

	go func() {
		for i := 1; i <= j; i++ {
			js := ""
			if i%3 == 0 {
				js += "Fizz"
			}
			if i%5 == 0 {
				js += "Buzz"
			}
			if js == "" {
				js = fmt.Sprintf("%v", i)
			}
			out <- js
		}
		close(out)
	}()
	return out
}
