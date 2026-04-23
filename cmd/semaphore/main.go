package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func cache() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	cacheCount.Add(1)
}

var cacheCount atomic.Int64
var smthCount atomic.Int64

func getSmth() {
	withSemaphore(func() {
		cache()
	})

	smthCount.Add(1)
}

var maxGoroutines = 10
var sem = make(chan struct{}, maxGoroutines)

func withSemaphore(f func()) {
	select {
	case sem <- struct{}{}:
	default:
		return
	}

	go func() {
		f()

		<-sem
	}()
}

/*
Semaphore is
*/
func main() {
	ticker := time.NewTicker(1 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

loop:
	for {
		select {
		case <-ticker.C:
			getSmth()
		case <-ctx.Done():
			break loop
		}
	}

	fmt.Println("smthCount:", smthCount.Load())
	fmt.Println("cacheCount:", cacheCount.Load())
}
