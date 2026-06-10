package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/**
1. Create a fixed-size pool of N workers at startup (goroutines blocking on a
  jobs channel)
2. Accept jobs via a Submit(job) method — send into the jobs channel
+ 3. Stop the jobs channel to signal workers to stop
+ 4. Wait for all in-flight workers to finish (sync.WaitGroup)
+ 5. Support context.Context cancellation — workers exit when ctx is done
*/

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
