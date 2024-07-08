package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Testing repeating 5 times...", func(t *testing.T) {
		got := Repeat('a', 7)
		want := "aaaaaaa"

		if got != want {
			t.Errorf("Failure: Got, '%v' expected '%v' ", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat('v', 8)
	fmt.Println(repeated)
	// Output: vvvvvvvv
}
