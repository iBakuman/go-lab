package cobra_test

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

var osArgs = []string{"testcobra", "testcobra"}

func TestCommand(t *testing.T) {
	cmd := cobra.Command{
		Use: "testcobra",
		Run: func(cmd *cobra.Command, args []string) {
			require.ElementsMatch(t, osArgs[1:], args)
		},
	}
	os.Args = osArgs
	require.Nil(t, cmd.Context())
	cmd.Execute()
	require.NotNil(t, cmd.Context())
}

func TestExecuteFuncInCommand(t *testing.T) {
	cmd := cobra.Command{
		Use: "testcobra",
		Run: func(cmd *cobra.Command, args []string) {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
			<-ch
			t.Log("signal received")
		},
	}
	os.Args = osArgs
	go func() {
		require.NoError(t, cmd.Execute())
	}()
	p, err := os.FindProcess(os.Getpid())
	require.NoError(t, err)
	err = p.Signal(os.Interrupt)
	require.NoError(t, err)
}
