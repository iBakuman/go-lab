package _defer

import "testing"

func TestDefer(t *testing.T) {
	fn := func() {
		return
		defer func() {
			t.Fatal("should not reach here")
		}()
	}
	fn()
}
