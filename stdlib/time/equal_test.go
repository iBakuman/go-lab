package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	t1 := time.Date(2021, 1, 0, 0, 0, 0, 9999999999, time.UTC)
	t2 := time.Date(2021, 1, 0, 0, 0, 9, 999999999, time.UTC)
	require.True(t, t1.Equal(t2))
}
