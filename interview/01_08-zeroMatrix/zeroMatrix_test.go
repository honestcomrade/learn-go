package main

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {

	tables := []struct {
		in  [][]uint8 // the string to test against
		out [][]uint8 // the result of the compression
	}{
		{
			in: [][]uint8{
				{1, 1, 2},
				{1, 1, 0},
				{1, 1, 2}},
			out: [][]uint8{
				{2, 2, 0},
				{0, 0, 0},
				{1, 1, 0}},
		},
	}

	for _, table := range tables {
		result := ZeroMatrix(table.in)
		if !reflect.DeepEqual(result, table.out) {
			t.Errorf("Result of ('%v') was incorrect, got: '%v', want: '%v'.", table.in, result, table.out)
		}
	}

}
