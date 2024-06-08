package channel

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBufferedChannel_01(t *testing.T) {
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

func TestBufferedChannel_02(t *testing.T) {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	inStream := make(chan int, 4)
	go func() {
		defer close(inStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 4; i++ {
			time.Sleep(20 * time.Millisecond)
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			inStream <- i
		}
	}()

	for i := range inStream {
		fmt.Fprintf(&stdoutBuff, "Received: %d\n", i)
	}
}

func TestBufferedChannel_03(t *testing.T) {
}
