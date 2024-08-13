package error_test

import (
	"net"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

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

func client(t *testing.T, addr string) {
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

func TestBrokenPipe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	addr := server(t, &wg)
	client(t, addr)
	wg.Wait()
}
