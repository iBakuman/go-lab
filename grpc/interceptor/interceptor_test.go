package interceptor

import (
	"context"
	"log"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ibakuman/go-lab/grpc/testkit"
	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("UnaryInterceptor: %v", info.FullMethod)
	return handler(ctx, req)
}

type PingRequest struct {
	Message string
}

type PingResponse struct {
	Message string
}

type PingService interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func (p *pingServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := p.cc.Invoke(ctx, "/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func TestInterceptor(t *testing.T) {
	conn, _ := testkit.SetupGRPCGatewayTestEnv(t,
		func(server *grpc.Server) {
		},
		func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return nil
		},
		grpc.ChainUnaryInterceptor(UnaryInterceptor))
	client := &pingServiceClient{cc: conn}
	_, err := client.Ping(context.Background(), &PingRequest{Message: "hello"})
	if err != nil {
		t.Fatal(err)
	}
}
