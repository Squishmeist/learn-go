package helloworld

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Ainsley", "")
		want := "Hello, Ainsley"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Ainsley", "Spanish")
		want := "Hola, Ainsley"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T){
		got := Hello("Ainsley", "French")
		want := "Bonjour, Ainsley"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}