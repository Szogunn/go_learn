package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying, hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() //t.Helper() is needed to tell the test suite that this method is a helper. 
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
