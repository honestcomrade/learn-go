package main

import "fmt"

// ZeroMatrix ...
// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0
func ZeroMatrix(x [][]int) [][]int {
	return x
}

func main() {
	matrix := [][]int{
		{1, 1, 2, 2},
		{1, 1, 0, 4},
		{1, 1, 2, 9},
		{1, 1, 2, 9},
	}

	fmt.Println(ZeroMatrix(matrix))
}
