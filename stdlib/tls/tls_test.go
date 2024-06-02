package tls

import (
	"crypto/tls"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTLSHandshake(t *testing.T) {
	cert, err := tls.LoadX509KeyPair(TlsCertFile, TlsKeyFile)
	require.NoError(t, err)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
	}
	listener, err := tls.Listen("tcp", "localhost:0", tlsConfig)
	require.NoError(t, err)
	t.Logf("TLS listener address: %s", listener.Addr().String())

	go func() {
		for {
			conn, err := listener.Accept()
			require.NoError(t, err)
			go func(conn net.Conn) {
				defer conn.Close()
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				require.NoError(t, err)
				t.Logf("Received message: %s", string(buf[:n]))
			}(conn)
		}
	}()

	clientConfig
}
