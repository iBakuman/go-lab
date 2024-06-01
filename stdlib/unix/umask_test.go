package unix

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestUmask(t *testing.T) {
	temp1, err := os.CreateTemp("", "temp1")
	require.NoError(t, err)
	defer os.Remove(temp1.Name())
	t.Logf("temp1.Name(): %s\n", temp1.Name())
	tmp1Info, err := os.Stat(temp1.Name())
	require.NoError(t, err)
	t.Logf("tmp1Info.Mode(): %s\n", tmp1Info.Mode())
	t.Logf("tmp1Info.Mode().Perm(): %o\n", tmp1Info.Mode().Perm())
	initialMask := unix.Umask(0)
	fmt.Printf("%o\n", initialMask)
}
