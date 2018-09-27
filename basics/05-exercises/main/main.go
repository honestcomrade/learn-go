package main

import (
	"fmt"

	"github.com/honestcomrade/go-basics/05-exercises/natural"
)

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if natural.NatChecker(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
