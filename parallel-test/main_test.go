package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	got := Add(1, 2, 3, 4)
	want := 10

	time.Sleep(3 * time.Second)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	got := Sub(10, 1, 2, 3)
	want := 4

	time.Sleep(3 * time.Second)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
