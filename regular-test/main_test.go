package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
		// or
		// t.Log("want", want)
		// t.Fail()
	}

	// or
	// if got := Hello(); got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
}
