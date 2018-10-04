package main

import "fmt"

// ZeroMatrix ...
// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0
func ZeroMatrix(x [][]byte) [][]byte {
	return x
}

func main() {
	matrix := [][]uint8{
		{1, 1, 2},
		{1, 1, 0},
		{1, 1, 2},
	}

	fmt.Println(ZeroMatrix(matrix))
}
