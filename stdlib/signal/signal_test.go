package signal_test

import (
	"os"
	"os/signal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterrupt(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		process, err := os.FindProcess(os.Getpid())
		require.NoError(t, err)
		require.NoError(t, process.Signal(os.Interrupt))
	}()
	<-c
	t.Logf("Signal received, shutting down...")
}
