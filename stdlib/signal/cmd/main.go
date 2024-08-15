package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch
	fmt.Println("Signal received, shutting down...")
}
