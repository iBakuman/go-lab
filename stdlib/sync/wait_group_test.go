package sync_test

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Wait()
	})
}