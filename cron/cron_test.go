//go:build long

package cron_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/require"
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
	stopCron(t, cronJobs)
}

func TestSkipIfStillRunning(t *testing.T) {
	runTest := func(t *testing.T, skipIfStillRunning bool) {
		var cronJobs *cron.Cron
		if skipIfStillRunning {
			cronJobs = cron.New(cron.WithChain(cron.SkipIfStillRunning(NewCronLogger())))
		} else {
			cronJobs = cron.New()
		}
		var job1Count, job2Count atomic.Int32
		cronJobs.AddFunc("@every 3s", func() {
			job1Count.Add(1)
			runJob(t, 1, 4*time.Second)
		})
		cronJobs.AddFunc("@every 3s", func() {
			job2Count.Add(1)
			runJob(t, 2, 2*time.Second)
		})
		cronJobs.Start()
		time.Sleep(10 * time.Second)
		stopCron(t, cronJobs)
		t.Logf("job 1 count: %d, job 2 count: %d", job1Count.Load(), job2Count.Load())
	}

	t.Run("skip if still running", func(t *testing.T) {
		runTest(t, true)
	})

	t.Run("do not skip if still running", func(t *testing.T) {
		runTest(t, false)
	})
}

func TestDelayIfStillRunning(t *testing.T) {
	runTest := func(t *testing.T, delayIfStillRunning bool) {
		var cronJobs *cron.Cron
		if delayIfStillRunning {
			cronJobs = cron.New(cron.WithChain(cron.DelayIfStillRunning(NewCronLogger())))
		} else {
			cronJobs = cron.New()
		}
		var job1Count, job2Count atomic.Int32
		cronJobs.AddFunc("@every 3s", func() {
			job1Count.Add(1)
			runJob(t, 1, 4*time.Second)
		})
		cronJobs.AddFunc("@every 3s", func() {
			job2Count.Add(1)
			runJob(t, 2, 2*time.Second)
		})
		cronJobs.Start()
		time.Sleep(10 * time.Second)
		stopCron(t, cronJobs)
		t.Logf("job 1 count: %d, job 2 count: %d", job1Count.Load(), job2Count.Load())
		require.Equal(t, int32(3), job1Count.Load())
	}

	t.Run("delay if still running", func(t *testing.T) {
		runTest(t, true)
	})

	t.Run("do not delay if still running", func(t *testing.T) {
		runTest(t, false)
	})
}
