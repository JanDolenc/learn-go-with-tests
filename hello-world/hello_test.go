package main

import "testing"

// Multiple (sub)tests grouped around a thing - func
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Jan", "")
		want := "Hello, Jan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Pedro", "Spanish")
		want := "Hola, Pedro!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pauline", "French")
		want := "Bonjour, Pauline!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Slovenian", func(t *testing.T) {
		got := Hello("Jan", "Slovenian")
		want := "Živjo, Jan!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tell test suite that assertCorrectMessage is a helper and it helps when tracking down the problems when test is failing - the line number reported is in our function call not in this helper func.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
