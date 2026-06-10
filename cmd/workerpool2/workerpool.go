package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pool struct {
	workersNum int
	queue      chan job
	wait       *sync.WaitGroup
}

func NewPool(workersNum int) *Pool {
	return &Pool{
		workersNum: workersNum,
		queue:      make(chan job, workersNum),
		wait:       &sync.WaitGroup{},
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workersNum; i++ {
		p.wait.Add(1)
		go func() {
			defer p.wait.Done()
			for {
				select {
				case j, ok := <-p.queue:
					if !ok {
						fmt.Println("Channel is closed")
						return
					}
					time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
					fmt.Println(j.Print())
				}
			}
		}()
	}
}

func (p *Pool) Submit(j job) {
	p.queue <- j
}

func (p *Pool) Stop() {
	close(p.queue)
}

func (p *Pool) Wait() {
	p.wait.Wait()
}
