package comb

import (
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
