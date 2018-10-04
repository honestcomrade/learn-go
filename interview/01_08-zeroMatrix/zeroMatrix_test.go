package main

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {

	tables := []struct {
		in  [][]int // the string to test against
		out [][]int // the result of the compression
	}{
		{
			in: [][]int{
				{1, 1, 2, 1},
				{1, 1, 0, 4},
				{1, 1, 2, 9},
				{1, 1, 2, 9},
			},
			out: [][]int{
				{1, 1, 0, 1},
				{0, 0, 0, 0},
				{1, 1, 0, 9},
				{1, 1, 0, 9},
			},
		},
		{
			in: [][]int{
				{1, 1, 2, 1},
				{1, 1, 0, 0},
				{1, 1, 2, 9},
				{1, 1, 2, 9},
			},
			out: [][]int{
				{1, 1, 0, 0},
				{0, 0, 0, 0},
				{1, 1, 0, 0},
				{1, 1, 0, 0},
			},
		},
	}

	for _, table := range tables {
		result := ZeroMatrix(table.in)
		if !reflect.DeepEqual(result, table.out) {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
