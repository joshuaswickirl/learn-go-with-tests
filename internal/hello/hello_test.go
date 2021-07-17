package hello_test

import (
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/hello"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello.Hello("Josh")
		want := "Hello, Josh"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello.Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
