package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	done := make(chan struct{}, 1)

	go func() {
		for i := 0; i < 1_000_00; i++ {
			select {
			case <-ctx.Done():
				done <- struct{}{}
				return
			default:
				fmt.Println("sleeping...")
				time.Sleep(4 * time.Second)
			}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("shutdown started")

		ctxT, cancelT := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelT()

		for {
			select {
			case <-ctxT.Done():
				fmt.Println("graceful shutdown timed out")
				os.Exit(1)
			case <-done:
				fmt.Println("graceful shutdown completed")
				return
			}
		}
	}
}
