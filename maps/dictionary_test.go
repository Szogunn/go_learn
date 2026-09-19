package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a text"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a text"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, DictionaryErr(ErrNotFound))
	})
}

func TestAdd(t *testing.T) {
	var key = "test"
	var value = "This is just a text"

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing key", func(t *testing.T) {	
		dictionary := Dictionary{key: value}
		
		err := dictionary.Add(key, "some new value")
		assertError(t, err, DictionaryErr(ErrWordExists))
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating existing key", func(t *testing.T) {
		dictionary := Dictionary{"test":"some text"}
		dictionary.Update("test", "new text")

		got, err := dictionary.Search("test")
		want := "new text"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("updating new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "some text")
	
		assertError(t, err, DictionaryErr(ErrKeyDoesNotExist))
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleting existing key", func(t *testing.T) {
		dictionary := Dictionary{"test":"some text"}
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		
		assertError(t, err, DictionaryErr(ErrNotFound))
	})

	t.Run("deleting non existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
	
		assertError(t, err, DictionaryErr(ErrKeyDoesNotExist))
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}
