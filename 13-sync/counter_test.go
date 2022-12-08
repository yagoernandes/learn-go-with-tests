package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("should count to 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, expected int) {
	t.Helper()
	if counter.Value() != expected {
		t.Errorf("got %v, want %v", counter.Value(), expected)
	}
}
