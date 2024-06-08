package channel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Summary:
// Read
// - Nil: blocking
// - Closed: zero value
// - Open and not empty: value
// - Open and empty: blocking
// - WriteOnly: Compile error
//
// Write
// - Nil: blocking
// - Closed: panic
// - Open and Not Full: Write value
// - Open and Full: blocking
// - ReadOnly: Compile error
//
// Close
// - nil: panic
// - Open and Not Empty: Closes Channel, read succeed until channel is drained, the reads produce zero value.
// - Open and Empty: Closes Channel, reads produce zero value.
// - Closed: panic
// - ReadOnly: Compile error
func TestOpOnChannel(t *testing.T) {
	t.Run("nil channel", func(t *testing.T) {
		testRWOpOnNilChannel := func(t *testing.T, isReadOp bool) {
			chanOwner := func() <-chan struct{} {
				done := make(chan struct{})
				go func() {
					defer close(done)
					var ch chan int
					require.Nil(t, ch)
					if isReadOp {
						<-ch
					} else {
						ch <- 1
					}
				}()
				return done
			}
			done := chanOwner()
			select {
			case <-done:
				t.Fatal("should not reach here")
			case <-time.After(300 * time.Millisecond):
			}
		}
		// blocking send on nil channel
		t.Run("blocking send on nil channel", func(t *testing.T) {
			testRWOpOnNilChannel(t, false)
		})

		t.Run("blocking receive on nil channel", func(t *testing.T) {
			testRWOpOnNilChannel(t, true)
		})

		t.Run("close nil channel", func(t *testing.T) {
			// panic: close of nil channel
			require.Panics(t, func() {
				var ch chan int
				close(ch)
			})
		})
	})

	t.Run("closed channel", func(t *testing.T) {
		t.Run("panic on sending to closed channel", func(t *testing.T) {
			ch := make(chan int)
			close(ch)
			require.Panics(t, func() {
				ch <- 1
			})
		})
		t.Run("zero value on receiving from closed channel", func(t *testing.T) {
			ch := make(chan int)
			close(ch)
			require.Zero(t, <-ch)
			v, ok := <-ch
			require.Zero(t, v)
			require.False(t, ok)
			for range ch {
				t.Fatal("should not reach here")
			}
		})
		t.Run("panic on closing closed channel", func(t *testing.T) {
			ch := make(chan int)
			close(ch)
			require.Panics(t, func() {
				close(ch)
			})
		})
		t.Run("read closed buffered channel", func(t *testing.T) {
			ch := make(chan int, 3)
			for i := 0; i < 3; i++ {
				ch <- i
			}
			close(ch)
			var res []int
			for v := range ch {
				res = append(res, v)
			}
			require.Equal(t, []int{0, 1, 2}, res)
		})
	})
}
