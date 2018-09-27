package main

import "fmt"

const conversion = 1.09361

func wrapper() func() float64 {
	meters := 26.0
	return func() float64 {
		meters = meters * conversion
		// fmt.Printf(meters)
		return meters
	}
}

func main() {
	increment := wrapper()

	fmt.Println(increment())
}
