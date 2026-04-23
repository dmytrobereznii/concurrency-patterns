// The Fan-In multiplexes multiple input channels into a single output channel.
// The Fan-Out allows to use multiple goroutines to process tasks from a single input channel.
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numChans := 5
	in := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	for v := range fanin(fanout(in, numChans)) {
		fmt.Println(v)
	}
}

func fanin(chans []chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)

			go func() {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						}
					}
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func fanout(in chan int, numChans int) []chan int {
	chans := make([]chan int, numChans)

	for i := range chans {
		chans[i] = pipeline(in)
	}

	return chans
}

func pipeline(in chan int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * rand.Intn(5)
		}
		close(out)
	}()

	return out
}
