package main

import "fmt"

func doStuff() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we Done yet?")
	fmt.Println("I'm going to do some stuff...")
}

func main() {
	doStuff()
}
