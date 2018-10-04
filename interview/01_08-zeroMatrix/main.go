package main

import "fmt"

// ZeroMatrix ...
// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0
func ZeroMatrix(x [][]int) [][]int {
	n := len(x)
	// create a slice of boleans that each index will refer
	// to a row or col that we found a zero at
	zr := make([]bool, n)
	zc := make([]bool, len(x[0]))

	for i := 0; i < n; i++ {
		for j := 0; j < len(x[0]); j++ {
			if x[i][j] == 0 {
				// save the row that we found a 0 at in the bool array of rows
				zr[i] = true
				// save the col that we found a 0 at in the bool array of cols
				zc[j] = true
			}
		}
	}
	// loop through again and zero out everything at that row
	for r, v := range zr {
		if v == true {
			zeroRow(x, r)
		}
	}
	// same as above but for cols
	for c, v := range zc {
		if v == true {
			zeroCol(x, c)
		}
	}
	return x
}

func zeroRow(x [][]int, r int) {
	for c := range x[0] {
		x[r][c] = 0
	}
}

func zeroCol(x [][]int, c int) {
	for r := range x {
		x[r][c] = 0
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 2, 2},
		{1, 1, 0, 4},
		{1, 1, 2, 9},
	}
	for i := range matrix {
		fmt.Print("\n")
		for _, v2 := range matrix[i] {

			fmt.Print(v2, " ")
		}
	}
	zeroedMatrix := ZeroMatrix(matrix)
	for i := range zeroedMatrix {
		fmt.Print("\n")
		for _, v2 := range zeroedMatrix[i] {

			fmt.Print(v2, " ")
		}
	}
}
