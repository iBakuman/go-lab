//go:build long

package cron

import (
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	cronJobs := cron.New()
	cronJobs.AddFunc("@every 2s", func() {
		t.Log("start")
		time.Sleep(10 * time.Second)
		t.Log("done")
	})
	cronJobs.Start()
	time.Sleep(3 * time.Second)
	<-cronJobs.Stop().Done()
}
