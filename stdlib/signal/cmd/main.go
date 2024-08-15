package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func runWithCtx() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		<-ch
		log.Printf("Interrupt signal received, shutting down...")
	}()
	if err := root.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

func runWithoutCtx() {
	if err := root.Execute(); err != nil {
		panic(err)
	}
}

func main() {
	runWithCtx()
}
