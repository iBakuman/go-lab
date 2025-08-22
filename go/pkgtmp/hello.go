package pkgtmp

import "context"

type Handler interface {
	Handle(ctx context.Context, data []byte)  error
}

func SayHello(ctx context.Context,handler Handler) error {
	return handler.Handle(ctx, []byte("hello world"))
}