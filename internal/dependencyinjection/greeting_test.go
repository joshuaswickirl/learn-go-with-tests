package dependencyinjection_test

import (
	"bytes"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/dependencyinjection"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	dependencyinjection.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris.\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
