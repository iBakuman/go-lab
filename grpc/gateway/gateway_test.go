package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ibakuman/go-lab/grpc/gen/gateway"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type testImpl struct {
	gateway.UnimplementedTestGatewayServer
}

func (s *testImpl) Test(ctx context.Context, req *gateway.TestRequest) (*gateway.TestResponse, error) {
	return &gateway.TestResponse{
		Msg: fmt.Sprintf("msg from request: %s", req.Msg),
	}, nil
}

func startGRPCServer(t *testing.T) *bufconn.Listener {
	// 101024 * 1024 = 10 MB
	listener := bufconn.Listen(101024 * 1024)
	server := grpc.NewServer()
	gateway.RegisterTestGatewayServer(server, &testImpl{})
	go func() {
		err := server.Serve(listener)
		if err != nil {
			require.ErrorContains(t, err, "closed")
		}
	}()
	t.Cleanup(func() {
		server.GracefulStop()
		require.NoError(t, listener.Close())
	})
	return listener
}

var httpServerAddr = "localhost:19282"

func startHTTPServer(t *testing.T, listener *bufconn.Listener) *httptest.Server {
	// target must be 'bufnet'
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return listener.DialContext(ctx)
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	mux := runtime.NewServeMux()
	require.NoError(t, gateway.RegisterTestGatewayHandler(context.Background(), mux, conn))
	server := httptest.NewServer(mux)
	server.URL = strings.Replace(server.URL, "127.0.0.1", "localhost", -1)
	t.Cleanup(server.Close)
	return server
}

func TestGateway(t *testing.T) {
	lis := startGRPCServer(t)
	server := startHTTPServer(t, lis)

	type RequestPayload struct {
		Msg string `json:"msg"`
	}
	payload := &RequestPayload{
		Msg: "hello",
	}
	data, err := json.Marshal(payload)
	require.NoError(t, err)
	t.Log(string(data))
	req, err := http.NewRequest(http.MethodPost, server.URL+"/v1/test", bytes.NewBuffer(data))
	require.NoError(t, err)
	resp, err := server.Client().Do(req)
	require.NoError(t, err)
	all, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Log(string(all))
}
