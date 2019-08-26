package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("hasan")
	want := "Hello, hasan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
