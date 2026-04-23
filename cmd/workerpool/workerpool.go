package main

import "fmt"

type Worker struct {
	id            int
	jobsCompleted int
}

var workers = []*Worker{
	{id: 0},
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
	{id: 6},
	{id: 7},
	{id: 8},
	{id: 9},
}

type Pool[T any] struct {
	pool    chan *Worker
	handler func(int, T)
}

func NewPool[T any](handler func(int, T)) *Pool[T] {
	return &Pool[T]{
		handler: handler,
		pool:    make(chan *Worker, len(workers)),
	}
}

func (p *Pool[T]) Create() {
	for _, w := range workers {
		p.pool <- w
	}
}

func (p *Pool[T]) Handle(t T) {
	w := <-p.pool
	go func() {
		p.handler(w.id, t)
		w.jobsCompleted++
		p.pool <- w
	}()
}

func (p *Pool[T]) Wait() {
	for range len(workers) {
		<-p.pool
	}
}

func (p *Pool[T]) Stats() {
	fmt.Println("_______Results:_______")
	for _, w := range workers {
		fmt.Printf("Worker %d completed %d jobs\n", w.id, w.jobsCompleted)
	}
}
