package main

import "fmt"

// ZeroMatrix ...
// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0
func ZeroMatrix(x [][]int) [][]int {
	n := len(x)
	zeroRow := 0
	zeroCol := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if x[i][j] == 0 {
				// save the location that we found the 0 at index
				zeroRow = i
				zeroCol = j
				fmt.Print("Found a zero at:", zeroRow, ":", zeroCol, "\n")
			}
		}
	}
	// loop through again and zero out everything at that row or col
	for i := range x {
		for j := range x[i] {
			x[zeroRow][j] = 0
			x[i][zeroCol] = 0
		}
	}
	return x
}

func main() {
	matrix := [][]int{
		{1, 1, 2, 2},
		{1, 1, 0, 4},
		{1, 1, 2, 9},
		{1, 1, 2, 9},
	}
	zeroedMatrix := ZeroMatrix(matrix)
	for i := range zeroedMatrix {
		fmt.Print("\n")
		for _, v2 := range zeroedMatrix[i] {

			fmt.Print(v2, " ")
		}
	}
	fmt.Println(ZeroMatrix(matrix))
}
