package maps_test

import (
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/maps"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		if got != want {
			t.Errorf("got %q, want %q, given %q", got, want, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Error("expected to get an error")
			t.FailNow()
		}
		nfe, ok := err.(NotFoundError)
		if !ok {
			t.Errorf("recieved error does not implement a NotFound method")
		}
		if nfe.NotFound() != true {
			t.Errorf("got %v, want %v", nfe.NotFound(), true)
		}
		if nfe.Error() != "could not find the word you were looking for" {
			t.Errorf("got %q, want %q", nfe.Error(), "could not find the word you were looking for")
		}
	})
}

// Use extensions of the error interface to provide a safe way to determine
// error handling without having to assert on the string from error.Error()
// Specifically, this is useful when function can return two or more different
// errors.
type NotFoundError interface {
	NotFound() bool
	Error() string
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := maps.Dictionary{}
		word := "test"
		def := "this is just a test"

		_ = dictionary.Add(word, def)

		got, err := dictionary.Search(word)
		if err != nil {
			t.Error("got an error but didn't expect one", err)
		}
		if got != def {
			t.Errorf("got %q, want %q", got, def)
		}
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := maps.Dictionary{word: def}
		err := dictionary.Add(word, "new test")
		got, _ := dictionary.Search(word)

		if err == nil {
			t.Error("didn't get an error but expected one")
		}
		if got != def {
			t.Errorf("got %q, want %q", got, def)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := maps.Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		if err != nil {
			t.Errorf("got an error and didn't expect one")
		}

		got, _ := dictionary.Search(word)

		if got != newDefinition {
			t.Errorf("got %q, want %q", got, newDefinition)
		}
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := maps.Dictionary{}

		err := dictionary.Update(word, definition)
		if err == nil {
			t.Errorf("didn't get an error but expect one")
		}

		got, err := dictionary.Search(word)
		if got != "" && err == nil {
			t.Errorf("update should not create a new record")
		}
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := maps.Dictionary{word: "test def"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err == nil {
		t.Errorf("didn't get an error and expected one")
	}
}
