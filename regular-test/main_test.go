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
}

func TestAdd(t *testing.T) {
	// with scenarios
	t.Run("return 0 with no params", func(t *testing.T) {
		if got := Add(); got != 0 {
			t.Errorf("got %q want 0", got)
		}
	})

	if got := Add(1, 2, 3, 4); got != 10 {
		t.Errorf("got %q want 10", got)
	}
}
