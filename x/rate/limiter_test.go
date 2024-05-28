package rate

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestLimiterAllow(t *testing.T) {
	limiter := rate.NewLimiter(1, 1)
	require.True(t, limiter.Allow())
	require.False(t, limiter.Allow())
	time.Sleep(1 * time.Second)
	require.True(t, limiter.Allow())
}

func TestDrainToken(t *testing.T) {
	cases := []struct {
		name    string
		r       rate.Limit
		b       int
		wait    time.Duration
		allowed bool
	}{
		{
			name:    "r equals b",
			r:       5,
			b:       5,
			wait:    200 * time.Millisecond,
			allowed: true,
		},
		{
			name:    "r less than b",
			r:       5,
			b:       10,
			wait:    100 * time.Millisecond,
			allowed: false,
		},
		{
			name:    "r greater than b",
			r:       10,
			b:       5,
			wait:    100 * time.Millisecond,
			allowed: true,
		},
		{
			name:    "r equals 0",
			r:       0,
			b:       5,
			wait:    0 * time.Millisecond,
			allowed: false,
		},
		{
			name:    "b equals 0",
			r:       5,
			b:       0,
			wait:    0 * time.Millisecond,
			allowed: false,
		},
		{
			name:    "wait is equal to 1000ms/r",
			r:       100,
			b:       10,
			wait:    10 * time.Millisecond,
			allowed: true,
		},
		{
			name:    "wait is less than 1000ms/r",
			r:       100,
			b:       10,
			wait:    5 * time.Millisecond,
			allowed: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			limiter := rate.NewLimiter(c.r, c.b)
			for i := 0; i < c.b; i++ {
				require.True(t, limiter.Allow())
			}
			time.Sleep(c.wait)
			require.Equal(t, c.allowed, limiter.Allow())
		})
	}
}

func TestLimiterReserve(t *testing.T) {
	limiter := rate.NewLimiter(1, 1)
	r1 := limiter.Reserve()
	require.True(t, r1.OK())
	require.EqualValues(t, 0, r1.Delay())
	r2 := limiter.Reserve()
	require.True(t, r2.OK())
	require.Greater(t, r2.Delay(), int64(0))
}

func TestLimiterWait(t *testing.T) {
	limiter := rate.NewLimiter(1, 1)
	ctx := context.Background()
	require.NoError(t, limiter.Wait(ctx))
	done := make(chan struct{})
	go func() {
		defer close(done)
		require.NoError(t, limiter.Wait(ctx))
	}()
	select {
	case <-done:
		t.Fatal("expected Wait() to block")
	case <-time.After(500 * time.Millisecond):
		// This is expected. as the limiter should block for approximately 1 second.
	}
	time.Sleep(1 * time.Second)
	select {
	case <-done:
	// This is expected. as the limiter should now have a token available.
	case <-time.After(500 * time.Millisecond):
		t.Error("expected Wait() to unblock after token is available")
	}
}
