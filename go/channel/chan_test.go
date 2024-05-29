package channel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)
	require.Equal(t, 0, len(ch))
	ch <- 1
	require.Equal(t, 3, cap(ch))
	require.Equal(t, 1, len(ch))
	ch <- 2
	require.Equal(t, 2, len(ch))
	ch <- 3
	require.Equal(t, 3, len(ch))
	i := <-ch
	require.Equal(t, 1, i)
	require.Equal(t, 2, len(ch))
	close(ch)
	for v := range ch {
		t.Log(v)
	}
}
