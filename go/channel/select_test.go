package channel

import "testing"

func TestSelect(t *testing.T) {
	c1 := make(chan int)
	close(c1)
	c2 := make(chan int)
	close(c2)
	var c1Count, c2Count int
	for i := 0; i < 10000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	t.Logf("c1Count: %d, c2Count: %d", c1Count, c2Count)
}
