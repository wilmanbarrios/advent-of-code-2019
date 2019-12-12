package main

import (
	"testing"
)

func TestInts(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output string
	}{
		{"sample 1", []int{1, 0, 0, 0, 99}, "2,0,0,0,99"},
		{"sample 2", []int{2, 3, 0, 3, 99}, "2,3,0,6,99"},
		{"sample 3", []int{2, 4, 4, 5, 99, 0}, "2,4,4,5,99,9801"},
		{"sample 4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, "30,1,1,4,2,5,6,0,99"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := Opcode(tc.input)
			if r != tc.output {
				t.Fatalf("%v should be %v; got %v", tc.name, tc.output, r)
			}
		})
	}
}
