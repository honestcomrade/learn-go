package main

import (
	"fmt"
)

// RotateMatrix ...
// Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degress.
// Can you do it in place?
func RotateMatrix(x [][]int) [][]int {
	n := len(x)
	// i represents a layer
	// only need to crawl through half of the layers
	// because updates will be done in place
	for i := 0; i < n/2; i++ {
		// save the first index in the layer
		begin := i
		// save the last index in the layer
		end := n - 1 - i
		// loop through each 'pixel' at this layer
		for j := begin; j < end; j++ {
			// save an offset which is the distance
			// from each j to the beginning of the layer
			os := j - begin
			// save the first index at this layer in some temp,
			// because it will be replaced first
			temp := x[begin][j]
			// put beginning of the last layer into the beginning of first layer
			x[begin][j] = x[end-os][begin]
			// put the end of the last layer into the beginning of the last layer
			x[end-os][begin] = x[end][end-os]
			// put the end of the first layer into the end of the last layer
			x[end][end-os] = x[j][end]
			// put the begin of first layer into end of first layer
			x[j][end] = temp
		}
	}
	return x
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotated := RotateMatrix(matrix)
	for _, v := range rotated {
		fmt.Print("\n")
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
	}
}
