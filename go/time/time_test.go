package time

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	dateStr := "2024-12-10"
	loc, err := time.LoadLocation("Asia/Tokyo")
	require.NoError(t, err)
	pTime, err := time.ParseInLocation(time.DateOnly, dateStr, loc)
	require.NoError(t, err)
	fmt.Println(pTime)
	fmt.Println(pTime.UTC())
}
