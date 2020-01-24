package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	got := Add(1, 2, 3, 4)
	want := 10

	// go test -timeout 2s
	// panic: test timed out after 2s
	time.Sleep(5 * time.Second)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
