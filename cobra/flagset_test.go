package cobra

import (
	"bytes"
	"testing"

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
