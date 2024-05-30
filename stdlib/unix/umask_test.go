package unix

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/sys/unix"
)

func TestUmask(t *testing.T) {
	os.Create("testfile1")
	initialMask := unix.Umask(0)
	fmt.Printf("%o\n", initialMask)
}
