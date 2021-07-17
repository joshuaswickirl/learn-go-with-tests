package hello_test

import (
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello("Josh")
	want := "Hello, Josh"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
