package cron_test

import (
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	kitlog "github.com/theplant/appkit/log"
)

func stopCron(t *testing.T, c *cron.Cron) {
	endTime := time.Now()
	t.Logf("notify jobs to stop, current time: %v", endTime)
	<-c.Stop().Done()
	t.Logf("all jobs stopped, current time: %v", time.Now())
	t.Logf("total elapsed time: %v", time.Since(endTime))
}

func runJob(t *testing.T, id int, sleep time.Duration) {
	startAt := time.Now()

	t.Logf("job %d start, current time: %v", id, time.Now())
	time.Sleep(sleep)
	t.Logf("job %d done, current time: %v", id, time.Now())
	t.Logf("job %d elapsed time: %v", id, time.Since(startAt))
}

type CronLogger struct {
	Log kitlog.Logger
}

func (cr *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	cr.Log.With("msg", msg).Info().Log(keysAndValues...)
}

func (cr *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	cr.Log.With("msg", msg, "err", err).Error().Log(keysAndValues...)
}

func NewCronLogger() *CronLogger {
	return &CronLogger{Log: kitlog.Default()}
}

