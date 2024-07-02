package cobra

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarkShorthandDeprecated(t *testing.T) {
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	buf := new(bytes.Buffer)
	flagSet.SetOutput(buf)
	deprecatedMsg := "please use --verbose"
	flagSet.BoolP("verbose", "v", false, "enable verbose mode")
	err := flagSet.MarkShorthandDeprecated("verbose", deprecatedMsg)
	require.NoError(t, err)
	err = flagSet.Parse([]string{"-v"})
	assert.Contains(t, buf.String(), deprecatedMsg)
	require.NoError(t, err)
	v, err := flagSet.GetBool("verbose")
	require.NoError(t, err)
	require.True(t, v)
}

func TestChanged(t *testing.T) {
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	buf := new(bytes.Buffer)
	flagSet.SetOutput(buf)
	flagSet.BoolP("verbose", "v", false, "enable verbose mode")
	flagSet.BoolP("debug", "d", false, "enable debug mode")
	flagSet.BoolP("trace", "t", true, "enable trace mode")
	flagSet.BoolP("tls", "s", false, "enable tls mode")
	require.Panics(t, func() {
		flagSet.BoolP("disabled", "d", true, "enable disabled mode")
	})
	err := flagSet.Parse([]string{"-v", "-t", "-d=false"})
	require.NoError(t, err)
	changed := flagSet.Changed("verbose")
	require.True(t, changed)
	changed = flagSet.Changed("debug")
	require.True(t, changed)
	changed = flagSet.Changed("trace")
	require.True(t, changed)
	changed = flagSet.Changed("tls")
	require.False(t, changed)
}

func TestVersion(t *testing.T) {
	cmd := &cobra.Command{
		Short: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("args: %v\n", args)
			return nil
		},
		Version: "v1.0.0",
	}
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	os.Args = []string{"test", "--version"}
	require.NoError(t, cmd.Execute())
	require.Contains(t, buf.String(), "v1.0.0")
}
