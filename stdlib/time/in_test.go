package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIn(t *testing.T) {
	now := time.Now()
	t.Logf("Now: %v", now)
	t.Logf("Now in UTC: %v", now.In(time.UTC))
	t.Logf("Now in Local: %v", now.In(time.Local))
	t.Logf("Now in UTC in Local: %v", now.In(time.UTC).In(time.Local))
	t.Logf("Now in Local in UTC: %v", now.In(time.Local).In(time.UTC))
	tokyoTZ, err := time.LoadLocation("Asia/Tokyo")
	require.NoError(t, err)
	t.Logf("Now in Asia/Tokyo: %v", now.In(tokyoTZ))
}
