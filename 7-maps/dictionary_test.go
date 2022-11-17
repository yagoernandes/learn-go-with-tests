package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown_word")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err_add := dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertError(t, err_add, nil)
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "existed"
		definition := "test for word that already exists"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "this is just a test")
		got, _ := dictionary.Search(word)

		assertError(t, err, ErrWordExists)
		assertStrings(t, got, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		old_definition := "old definition"
		new_definition := "new definition"
		dictionary := Dictionary{word: old_definition}

		err := dictionary.Update(word, new_definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, new_definition)
	})
	t.Run("unexisting word", func(t *testing.T) {
		word := "word"
		new_definition := "new definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, new_definition)
		assertError(t, err, ErrWordDoesNotExists)

	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		dictionary := Dictionary{word: "test"}

		err := dictionary.Delete(word)

		assertError(t, err, nil)
		assertEmptyDefinition(t, dictionary, word)
	})

	t.Run("unexisting word", func(t *testing.T) {
		word := "word"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExists)
		assertEmptyDefinition(t, dictionary, word)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertStrings(t, got, want)
	assertError(t, err, nil)
}

func assertEmptyDefinition(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertStrings(t, got, "")
	assertError(t, err, ErrNotFound)
}
