package http

// Dependencies for tests:
// - github.com/stretchr/testify/assert
// - github.com/stretchr/testify/require

import (
	"crypto/tls"
	"io"
	nethttp "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// helloHandler responds with a static "hello" payload.
func helloHandler() nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		_, _ = w.Write([]byte("hello"))
	})
}

// logTLSState emits key TLS handshake details for study.
func logTLSState(t testing.TB, who string, cs *tls.ConnectionState) {
	if cs == nil {
		t.Logf("%s: no TLS (plain HTTP)", who)
		return
	}

	t.Logf(
		"%s: handshakeComplete=%v version=0x%04x cipherSuite=0x%04x alpn=%q serverName=%q resumed=%v peerCerts=%d",
		who,
		cs.HandshakeComplete,
		cs.Version,
		cs.CipherSuite,
		cs.NegotiatedProtocol,
		cs.ServerName,
		cs.DidResume,
		len(cs.PeerCertificates),
	)
}

func TestHTTPServer_Hello(t *testing.T) {
	srv := httptest.NewServer(helloHandler())
	defer srv.Close()

	client := srv.Client()
	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	t.Logf(srv.URL)
	require.NotNil(t, resp)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, nethttp.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello", string(body))

	assert.Nil(t, resp.TLS)
	logTLSState(t, "client", resp.TLS)
}

func TestHTTPSServer_HelloAndHandshake(t *testing.T) {
	// Wrap the shared handler to also log server-side TLS details.
	h := nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		logTLSState(t, "server", r.TLS)
		_, _ = w.Write([]byte("hello"))
	})

	srv := httptest.NewTLSServer(h)
	defer srv.Close()

	client := srv.Client() // trusts the server's self-signed cert
	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	require.NotNil(t, resp)
	defer resp.Body.Close()
	t.Logf(srv.URL)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, nethttp.StatusOK, resp.StatusCode)
	require.NotNil(t, resp.TLS)
	logTLSState(t, "client", resp.TLS)
	assert.Equal(t, "hello", string(body))
	for {
		time.Sleep(20 * time.Second)
	}
}
