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
	gateway.UnimplementedTestGatewayServiceServer
}

func (s *testImpl) Echo(ctx context.Context, req *gateway.EchoRequest) (*gateway.EchoResponse, error) {
	return &gateway.EchoResponse{
		Msg: fmt.Sprintf("Echo: %s", req.Msg),
	}, nil
}

func (s *testImpl) EchoPathParams(ctx context.Context, req *gateway.EchoRequest) (*gateway.EchoResponse, error) {
	return &gateway.EchoResponse{
		Msg: fmt.Sprintf("EchoPathParams: %s", req.Msg),
	}, nil
}

func (s *testImpl) EchoQueryParams(ctx context.Context, req *gateway.EchoRequest) (*gateway.EchoResponse, error) {
	return &gateway.EchoResponse{
		Msg: fmt.Sprintf("EchoQueryParams: %s", req.Msg),
	}, nil
}
func setupGRPCGatewayTestEnv(t *testing.T,
	grpcRegister func(*grpc.Server),
	gatewayRegister func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error,
) (*grpc.ClientConn, *httptest.Server) {
	// 10240 * 1024 = 10 MB
	lis := bufconn.Listen(10240 * 1024)
	grpcServer := grpc.NewServer()
	grpcRegister(grpcServer)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			require.ErrorContainsf(t, err, "closed", "unexpected error: %v occurred", err)
		}
	}()
	t.Cleanup(func() {
		grpcServer.GracefulStop()
		require.NoError(t, lis.Close())
	})
	// the first argument passed to grpc.NewClient must be "bufnet"
	conn, err := grpc.NewClient("bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	mux := runtime.NewServeMux()
	require.NoError(t, gatewayRegister(context.Background(), mux, conn))
	httpServer := httptest.NewServer(mux)
	httpServer.URL = strings.Replace(httpServer.URL, "127.0.0.1", "localhost", -1)
	t.Cleanup(httpServer.Close)
	return conn, httpServer
}

func TestGateway(t *testing.T) {
	_, server := setupGRPCGatewayTestEnv(t, func(server *grpc.Server) {
		gateway.RegisterTestGatewayServiceServer(server, &testImpl{})
	}, gateway.RegisterTestGatewayServiceHandler)

	t.Run("request body", func(t *testing.T) {
		payload := gateway.EchoRequest{
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
	})

	t.Run("path params", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, server.URL+"/v1/test/hello_from_client", nil)
		require.NoError(t, err)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		all, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(all), "EchoPathParams: hello_from_client")
		t.Log(string(all))
	})

	t.Run("query params", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, server.URL+"/v1/test?msg=hello_from_client", nil)
		require.NoError(t, err)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		all, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(all), "EchoQueryParams: hello_from_client")
		t.Log(string(all))
	})
}
