package local_addr_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLocalAddrByUDP(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, conn.Close())
	})
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	t.Logf("Local address: %s\n", localAddr.String())
}

func TestGetLocalAddrByLooping(t *testing.T) {
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	require.NoError(t, err)
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	for _, ip := range ips {
		t.Logf("Local address: %s\n", ip.String())
	}
}
