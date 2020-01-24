package main

import "testing"

var values = []struct {
	scenario string
	in       []int
	out      int
}{
	{"Scenario 1 want 3", []int{1, 2}, 3},
	{"Scenario 2 want 10", []int{1, 2, 3, 4}, 10},
}

func TestAdd(t *testing.T) {
	for _, v := range values {
		got := Add(v.in...)

		if got != v.out {
			t.Errorf("Failed %v, got %v want %v", v.scenario, got, v.out)
		}
	}
}
