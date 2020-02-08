package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Liene")
	want := "Hello, Liene"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
