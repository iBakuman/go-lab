package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBefore(t *testing.T) {
	tokyoLoc, err := time.LoadLocation("Asia/Tokyo")
	require.NoError(t, err)
	shanghaiLoc, err := time.LoadLocation("Asia/Shanghai")
	require.NoError(t, err)
	tokyoTime := time.Date(2021, 1, 0, 1, 30, 0, 0, tokyoLoc)
	shanghaiTime := time.Date(2021, 1, 0, 0, 30, 0, 0, shanghaiLoc)
	require.True(t, tokyoTime.Equal(shanghaiTime))
	tokyoTimeStr := "2020-12-31T01:30:00+09:00"
	// shanghaiTimeStr := "2020-12-31T00:30:00+08:00"
	tokyoTime2, err := time.ParseInLocation(time.RFC3339, tokyoTimeStr, shanghaiLoc)
	require.NoError(t, err)
	require.True(t, tokyoTime2.Equal(shanghaiTime))

	t.Logf("Tokyo Time: %v", tokyoTime)
}
