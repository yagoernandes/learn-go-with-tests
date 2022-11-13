package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("NameTest")
	want := "Olá, NameTest!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
