package error_test

import (
	"net"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Quoted from https://gosamples.dev/connection-reset-by-peer/
// Both connection reset by peer and broken pipe errors occur when a peer (the other end)
// unexpectedly closes the underlying connection. However, there is a subtle difference
// between them. Usually, you get the connection reset by peer when you read from the connection
// after the server sends the RST packet, and when you write to the connection after the RST instead,
// you get the broken pipe error.

func server(t *testing.T, wg *sync.WaitGroup) string {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		require.NoError(t, listener.Close())
	})

	go func() {
		defer wg.Done()
		conn, err := listener.Accept()
		require.NoError(t, err)
		data := make([]byte, 1)
		if _, err := conn.Read(data); err != nil {
			t.Error(err)
			return
		}
		require.NoError(t, conn.Close())
	}()
	return listener.Addr().String()
}

func clientA(t *testing.T, addr string) {
	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)

	_, err = conn.Write([]byte("a"))
	require.NoError(t, err)

	time.Sleep(time.Second)

	_, err = conn.Write([]byte("b"))
	require.NoError(t, err)

	time.Sleep(time.Second)

	_, err = conn.Write([]byte("c"))
	t.Logf("client error: %+v\n", err)
	require.ErrorIs(t, err, syscall.EPIPE)
}

// The client generate `connection reset by peer` error
// If the server closes the connection with the remaining bytes in the socket’s receive buffer,
// then an RST packet is sent to the client. When the client tries to read from such a closed connection,
// it will get the `connection reset by peer error`.
func clientB(t *testing.T, addr string) {
	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)

	// The server will close the connection after reading the first byte
	// we send two bytes to the server, this will cause the server to send an RST packet.
	_, err = conn.Write([]byte("ab"))
	require.NoError(t, err)

	time.Sleep(time.Second)
	data := make([]byte, 1)
	_, err = conn.Read(data)
	t.Logf("client error: %+v\n", err)
	require.ErrorIs(t, err, syscall.ECONNRESET)
}

func TestBrokenPipe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	addr := server(t, &wg)
	clientA(t, addr)
	wg.Wait()
}

func TestResetByPeer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	addr := server(t, &wg)
	time.Sleep(2 * time.Second)
	clientB(t, addr)
	wg.Wait()
}
