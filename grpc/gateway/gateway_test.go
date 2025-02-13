package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ibakuman/go-lab/grpc/gen/gateway"
	"github.com/ibakuman/go-lab/grpc/testkit"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
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
func TestGateway(t *testing.T) {
	_, server := testkit.SetupGRPCGatewayTestEnv(t, func(server *grpc.Server) {
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
