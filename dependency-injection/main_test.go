package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buff := bytes.Buffer{}
	Greet(&buff, "Yo!")

	got := buff.String()
	want := "Hey, Yo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
