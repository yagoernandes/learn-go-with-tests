package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("NameTest", "")
		want := "Hello, NameTest"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World when a empty string is suplied'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Jesus", spanish)
		want := "Hola, Jesus"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Jaques", french)
		want := "Bonjour, Jaques"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("José", portuguese)
		want := "Olá, José"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
