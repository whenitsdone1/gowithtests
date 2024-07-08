package Dictionary

import (
	"errors"
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	dict := Dictionary{
		"MKDIR": "Create a directory on unix-line system",
	}
	t.Run("Testing normal search..", func(t *testing.T) {

		got := dict.Search("MKDIR")
		want := "Create a directory on unix-line system"

		assertStrings(t, got, want)
	})
	t.Run("Testing failed search", func(t *testing.T) {
		got := dict.Search("aaaaaa")
		want := "Key not found"
		assertStrings(t, got, want)

	})

}

func Test_Add(t *testing.T) {
	t.Run("testing adding...", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("Wow", "Surprise")
		got := dict.Search("Wow")
		want := "Surprise"
		assertStrings(t, got, want)
	})
	t.Run("testing adding with word that already exists...", func(t *testing.T) {
		dict := Dictionary{"Wow": "Surprise"}
		got := dict.Add("Wow", "Surprise")
		if !errors.Is(got, alreadyPresentError) {
			t.Fatal("Did not find expected error")
		}
	})
}

func Test_Update(t *testing.T) {
	t.Run("Base case....", func(t *testing.T) {
		dict := Dictionary{"Cat": "Feline"}
		dict.Update("Cat", "Domesticated Feline")
		got := dict["Cat"]
		want := "Domesticated Feline"
		assertStrings(t, got, want)
	})
	t.Run("Testing not in dict....", func(t *testing.T) {
		dict := Dictionary{}
		got := dict.Update("Cat", "Domesticated Feline")
		if !errors.Is(got, notPresentError) {
			t.Fatal("FATAL: Correct error not thrown")
		}
		fmt.Println(notPresentError.Error())
		_, exists := dict["Cat"]
		if !exists {
			t.Fatal("Word was not added as expected")
		}
	})

}

func Test_Delete(t *testing.T) {
	t.Run("Base case...", func(t *testing.T) {
		dict := Dictionary{"Cat": "Hat"}
		got := dict.Delete("Cat")
		if got != nil {
			t.Fatal("FATAL: Encountered error where none was expected")
		}
		_, exists := dict["Cat"]
		if exists {
			t.Fatal("FATAL: Word could not be removed from map")
		}
	})
	t.Run("Not in dictionary case....", func(t *testing.T) {
		dict := Dictionary{}
		got := dict.Delete("Sam")
		if !errors.Is(got, notPresentError) {
			t.Fatal("FATAL: Expected error not encountered...")
		}
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Failure: got %s want %s", got, want)
	}
}
