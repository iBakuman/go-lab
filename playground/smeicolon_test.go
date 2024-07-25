package playground

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func FunA(a int) (int, int) {
	return 2 * a, 4 * a
}

func FunB(a int) (err error) {
	err, a = errors.New("hello"), 2
	return
}

func FunC(t *testing.T, a int) (err error) {
	errStr := "hello"
	defer func() {
		if err == nil {
			t.Fatal("err is nil")
		} else {
			if err.Error() != errStr {
				t.Fatal("err is not hello")
			}
		}
		err = nil
	}()
	err, b := errors.New(errStr), 2
	t.Logf("a: %d, b: %d", a, b)
	return
}

func TestSemicolonFuncA(t *testing.T) {
	a := 4
	t.Logf("&a: %p", &a)
	a, b := FunA(a)
	t.Logf("&a: %p", &a)
	t.Logf("a: %d, b: %d", a, b)
}

func TestSemicolonFuncC(t *testing.T) {
	require.NoError(t, FunC(t, 2))
}
