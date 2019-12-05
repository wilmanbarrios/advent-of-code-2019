package main

import (
	"testing"
)

func TestInts(t *testing.T) {
	tt := []struct {
		name   string
		mass   int
		result int
	}{
		{"mass of 14", 14, 2},
		{"mass of 1969", 1969, 966},
		{"mass of 100756", 100756, 50346},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := ProcessMass(tc.mass)
			if r != tc.result {
				t.Fatalf("%v needs %v of fuel; got %v", tc.name, tc.result, r)
			}
		})
	}
}
