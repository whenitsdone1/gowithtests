package sum

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_sum(t *testing.T) {
	t.Run("Testing sum function...", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5, 6})
		want := 21

		if got != want {
			t.Errorf("Failure: Got '%d', expected '%d'", got, want)
		}
	})

}

// func Test_all(t *testing.T) {
// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	want := []int{3, 9}

// 	if !slices.Equal(got, want) {
// 		t.Errorf("Failure: Got '%v', expected '%v'", got, want)
// 	}
// }

func Test_sumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("Failure: Got '%v', expected '%v'", got, want)
		}
	}
	t.Run("Testing tail summation..", func(t *testing.T) {
		got := SumTails([]int{1, 2, 1}, []int{0, 9, 3})
		want := []int{3, 12}

		checkSums(t, got, want)
	})

	t.Run("Sum empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
