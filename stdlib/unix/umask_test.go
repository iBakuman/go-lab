package unix

import (
	"os"
	"testing"
)

func TestUmask(t *testing.T) {
	os.Create("testfile1")
}
