package main

import "github.com/honestcomrade/go-basics/03-fb/fb"

func main() {
	h := "Fizz"
	w := "Buzz"
	newstring := fb.Setstrings(h, w)
	var slice []string = newstring[0:2]
	for i := 1; i <= 100; i++ {
		fb.Stringchecker(slice, i)
	}
}
