package main

import "testing"

var values = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestAdd(t *testing.T) {
	for _, v := range values {
		got := Add(v.in...)

		if got != v.out {
			t.Errorf("got %v want %v", got, v.out)
		}
	}
}
