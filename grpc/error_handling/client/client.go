package main

import (
	"context"
	"flag"
	"log"
	"os/user"
	"time"

	"github.com/ibakuman/go-lab/grpc/gen/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var addr = flag.String("addr", "localhost:50052", "the address to serve on")

func main() {
	flag.Parse()

	name := "unknown"
	if u, err := user.Current(); err == nil && u.Username != "" {
		name = u.Username
	}

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := common.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, reqName := range []string{"", name} {
		r, err := c.SayHello(ctx, &common.HelloRequest{Name: reqName})
		if err != nil {
			if status.Code(err) != codes.InvalidArgument {
				log.Printf("Received unexpected error: %v", err)
				continue
			}
			log.Printf("Received expected error: %v", err)
			continue
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
