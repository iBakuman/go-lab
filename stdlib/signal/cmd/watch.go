package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ory/graceful"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use: "watch",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				shouldExists bool
				iter         = 1
			)
			for !shouldExists {
				select {
				case <-cmd.Context().Done():
					shouldExists = true
				default:
					doSomething(iter)
				}
				iter++
			}
			log.Printf("graceful shutdown...")
		},
	})

	root.AddCommand(&cobra.Command{
		Use: "signal",
		Run: func(cmd *cobra.Command, args []string) {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
			var (
				shouldExists bool
				iter         = 1
			)
			for !shouldExists {
				select {
				case <-ch:
					shouldExists = true
				default:
					doSomething(iter)
				}
				iter++
			}
			log.Printf("graceful shutdown...")
		},
	})

	root.AddCommand(&cobra.Command{
		Use: "g",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				shouldExists bool
				lock         sync.Mutex
			)
			return graceful.Graceful(func() error {
				iter := 1
				for {
					lock.Lock()
					if shouldExists {
						lock.Unlock()
						break
					}
					lock.Unlock()
					doSomething(iter)
					iter++
				}
				return nil
			}, func(ctx context.Context) error {
				log.Printf("shudown func called")
				lock.Lock()
				shouldExists = true
				lock.Unlock()
				return nil
			})
		},
	})
}

func doSomething(iter int) {
	for i := 1; i <= 4; i++ {
		log.Printf("iteration %d, progress %d", iter, i)
		time.Sleep(time.Second)
	}
}
