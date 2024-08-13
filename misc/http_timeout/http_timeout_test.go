package http_timeout_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func helloHandler(_ http.ResponseWriter, _ *http.Request) {
	time.Sleep(2 * time.Second)
}

func TestHTTPClientTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(helloHandler))
	t.Cleanup(server.Close)
	client := http.Client{
		Timeout: time.Second,
	}
	_, err := client.Get(server.URL)
	require.ErrorContains(t, err, "context deadline exceeded")
	t.Logf("error: %+v", err)
}

// set timeout per request
func TestHTTPRequestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(helloHandler))
	t.Cleanup(server.Close)
	var client http.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	require.NoError(t, err)
	_, err = client.Do(req)
	require.ErrorContains(t, err, "context deadline exceeded")
	t.Logf("error: %+v\n", err)
}
