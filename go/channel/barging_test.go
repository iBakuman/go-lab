package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestBarging(t *testing.T) {
	type token struct{}
	sem := make(chan token, 1)

	// Goroutine 1: 等待发送 token
	go func() {
		time.Sleep(1 * time.Second) // 确保它稍后开始
		sem <- token{}
		fmt.Println("Goroutine 1: sent token")
	}()

	// Goroutine 2: 尝试发送 token
	go func() {
		time.Sleep(1 * time.Second) // 确保它稍后开始
		select {
		case sem <- token{}:
			fmt.Println("Goroutine 2: sent token")
		default:
			fmt.Println("Goroutine 2: could not send token")
		}
	}()

	time.Sleep(3 * time.Second)
}
