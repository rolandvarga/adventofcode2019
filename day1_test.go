package main

import (
	"testing"
)

func TestFuelCounterUpper(t *testing.T) {
	var cases = []struct {
		mass int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		got := FuelCounterUpper(c.mass)
		if got != c.want {
			t.Errorf("FuelCounterUpper returned unexepcted value for mass '%d'\ngot: '%d'\nwant: '%d'", c.mass, got, c.want)
		}
	}
}
