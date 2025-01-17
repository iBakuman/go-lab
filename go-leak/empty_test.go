package go_leak

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGoroutineA(t *testing.T) {
	go func() {
		time.Sleep(200*time.Minute)
		fmt.Println("hello world!")
	}()
}

