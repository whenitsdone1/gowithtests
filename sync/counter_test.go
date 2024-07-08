package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times should have it be at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
		t.Run("Our counter should be thread-safe...", func(t *testing.T) {
			expected := 1000
			counter := Counter{}

			var wg sync.WaitGroup
			wg.Add(expected)

			for i := 0; i < expected; i++ {
				go func() {
					counter.Inc()
					wg.Done()
				}()
			}
			wg.Wait()

			assertCounter(t, &counter, expected)
		})
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("FATAL: got %d, expected %d", got.Value(), want)
	}
}
