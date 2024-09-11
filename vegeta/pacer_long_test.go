package vegeta

import (
	"testing"
	"time"
)

func TestPacer(t *testing.T) {
	pacer := vegeta.ConstantPacer{Per: time.Minute, Freq: 2}
	startTime := time.Now()
	var count uint64
	for {
		nextStart, exit := pacer.Pace(time.Since(startTime), count)
		if exit {
			break
		}
		t.Logf("nextStart: %v", nextStart)
		count++
		time.Sleep(nextStart)
	}
}
