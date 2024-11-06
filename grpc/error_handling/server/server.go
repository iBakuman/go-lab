package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/ibakuman/go-lab/grpc/gen/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 50052, "the port to serve on")

type server struct {
	common.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *common.HelloRequest) (*common.HelloReply, error) {
	if in.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	}
	return &common.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	common.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
