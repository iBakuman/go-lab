package playground

import (
	"testing"
)

func FunA(a int) (int, int) {
	return 2 * a, 4 * a
}

func TestSemicolon(t *testing.T) {
	a := 4
	t.Logf("&a: %p", &a)
	a, b := FunA(a)
	t.Logf("&a: %p", &a)
	t.Logf("a: %d, b: %d", a, b)
}
