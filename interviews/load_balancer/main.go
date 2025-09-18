package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID    int
	queue chan string
}

type LoadBalancer struct {
	workers         []Worker
	workersIterator int
	wg              sync.WaitGroup
}

func NewLoadBalancer(n int) *LoadBalancer {
	lb := LoadBalancer{workersIterator: 0}

	workers := make([]Worker, 0, n)
	for i := range n {
		workers = append(workers, *NewWorker(i+1, &lb.wg))
	}

	lb.workers = workers
	return &lb
}

func (lb *LoadBalancer) Dispatch(task string) {
	worker := lb.workers[lb.workersIterator]

	lb.wg.Add(1)
	worker.queue <- task

	lb.workersIterator = (lb.workersIterator + 1) % len(lb.workers)
}

func (lb *LoadBalancer) Stop() {
	for _, w := range lb.workers {
		close(w.queue)
	}

	lb.wg.Wait()
}

func NewWorker(id int, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		ID:    id,
		queue: make(chan string),
	}

	go func() {
		for task := range worker.queue {
			fmt.Printf("Worker %d got task %s\n", worker.ID, task)
			time.Sleep(time.Second * 1)
			wg.Done()
		}
	}()

	return worker
}

// используется round robin
func main() {
	lb := NewLoadBalancer(3) // три воркера
	lb.Dispatch("task1")
	lb.Dispatch("task2")
	lb.Dispatch("task3")
	lb.Dispatch("task4")

	// ждём завершения всех задач
	lb.Stop()
}
