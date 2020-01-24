package main

import "testing"

func TestLongFunc(t *testing.T) {
	// go test -short
	if testing.Short() {
		t.Skip("Skipping long test.")
	}

	// ...
}

func TestAwsesome(t *testing.T) {
	// go test -v
	t.Skip("Not implement yet.")
}
