package ints

import (
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	tests := []struct {
		a, min, max int
		want        int
	}{
		{5, 7, 9, 8},
		{5, 6, 8, 8},
		{5, 5, 7, 5},
		{5, 4, 6, 5},
		{5, 3, 5, 5},
		{5, 2, 4, 2},
		{5, 1, 3, 2},

		{0, 2, 4, 3},
		{0, 1, 3, 3},
		{0, 0, 2, 0},
		{0, -1, 1, 0},
		{0, -2, 0, 0},
		{0, -3, -1, -3},
		{0, -4, -2, -3},

		{-5, -3, -1, -2},
		{-5, -4, -2, -2},
		{-5, -5, -3, -5},
		{-5, -6, -4, -5},
		{-5, -7, -5, -5},
		{-5, -8, -6, -8},
		{-5, -9, -7, -8},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d/%d/%d", test.a, test.min, test.max), func(t *testing.T) {
			if have, want := Wrap(test.a, test.min, test.max), test.want; have != want {
				t.Fatalf("Wrap(%d, %d, %d), have %d, want %d", test.a, test.min, test.max, have, want)
			}
		})
	}
}
