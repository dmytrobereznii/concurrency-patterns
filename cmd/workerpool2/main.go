package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type job struct {
	val int
}

func (j job) Print() string {
	return fmt.Sprintf("Job %d", j.val)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pool := NewPool(runtime.NumCPU())
	pool.Start()

	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			pool.Submit(job{val: i})
		}
	}

	pool.Stop()
	pool.Wait()
}
