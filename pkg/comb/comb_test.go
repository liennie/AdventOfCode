package comb

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestComb(t *testing.T) {
	c := slices.Collect(Comb([]int{1, 2, 3}))
	want := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}

	if !reflect.DeepEqual(c, want) {
		t.Fatalf("want: %v, have: %v", want, c)
	}
}

func TestChoose(t *testing.T) {
	tests := []struct {
		n, k int
		want int
	}{
		{6, 0, 1},
		{6, 1, 6},
		{6, 2, 15},
		{6, 3, 20},
		{6, 4, 15},
		{6, 5, 6},
		{6, 6, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d choose %d", test.n, test.k), func(t *testing.T) {
			have := Choose(test.n, test.k)
			if have != test.want {
				t.Errorf("have: %d, want: %d", have, test.want)
			}
		})
	}
}
