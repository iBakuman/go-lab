package backoff_test

import (
	"context"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/require"
)

func TestContextBackoff(t *testing.T) {
	baf := backoff.NewConstantBackOff(10 * time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	contextBackoff := backoff.WithContext(baf, ctx)
	nextWaitTime := contextBackoff.NextBackOff()
	require.Equal(t, 10*time.Millisecond, nextWaitTime)
	cancel()
	require.Equal(t, backoff.Stop, contextBackoff.NextBackOff())
}
