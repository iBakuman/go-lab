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

func TestTwoJob(t *testing.T) {
	cronJobs := cron.New()
	cronJobs.AddFunc("@every 2s", func() {
		t.Log("job 1 start, current time: ", time.Now())
		time.Sleep(3 * time.Second)
		t.Log("job 1 done, current time: ", time.Now())
	})
	cronJobs.AddFunc("@every 4s", func() {
		t.Log("job 2 start, current time: ", time.Now())
		time.Sleep(3 * time.Second)
		t.Log("job 2 done, current time: ", time.Now())
	})
	cronJobs.Start()
	t.Log("cron started")
	time.Sleep(5 * time.Second)
	endTime := time.Now()
	t.Logf("notify jobs to stop, current time: %v", endTime)
	<-cronJobs.Stop().Done()
	t.Logf("all jobs stopped, current time: %v", time.Now())
	t.Logf("total elapsed time: %v", time.Since(endTime))
}