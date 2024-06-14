package time

import (
	"testing"
	"time"
)

func TestTruncate(t *testing.T) {
	fakeTime := time.Date(2024, 6, 14, 12, 34, 56, 123456789, time.UTC)
	t.Logf("Now: %v", fakeTime)
	truncatedHour := fakeTime.Truncate(time.Hour)
	t.Logf("Truncated Hour: %v", truncatedHour)
	truncatedMinute := fakeTime.Truncate(time.Minute)
	t.Logf("Truncated Minute: %v", truncatedMinute)
	truncatedSecond := fakeTime.Truncate(time.Second)
	t.Logf("Truncated Second: %v", truncatedSecond)
	truncatedMillisecond := fakeTime.Truncate(time.Millisecond)
	t.Logf("Truncated Millisecond: %v", truncatedMillisecond)
	truncatedMicrosecond := fakeTime.Truncate(time.Microsecond)
	t.Logf("Truncated Microsecond: %v", truncatedMicrosecond)
	truncatedNanosecond := fakeTime.Truncate(time.Nanosecond)
	t.Logf("Truncated Nanosecond: %v", truncatedNanosecond)
}
