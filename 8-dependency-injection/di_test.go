package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Yago")

	got := buffer.String()
	want := "Hello, Yago"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
