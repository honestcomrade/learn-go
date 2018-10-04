package main

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {

	tables := []struct {
		in  [][]int // the string to test against
		out [][]int // the result of the compression
	}{
		{
			in: [][]int{
				{1, 1, 1, 2},
				{1, 1, 1, 3},
				{1, 1, 1, 4},
				{1, 1, 1, 5},
			},
			out: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{5, 4, 3, 2},
			},
		},
	}

	for _, table := range tables {
		result := RotateMatrix(table.in)
		if !reflect.DeepEqual(result, table.out) {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
