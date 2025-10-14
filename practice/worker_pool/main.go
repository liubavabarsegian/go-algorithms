package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, tasks <-chan string) {
	defer wg.Done()
	for t := range tasks {
		time.Sleep(time.Second * 1)
		fmt.Printf("Worker %d started task %s \n", id, t)

	}
}

func main() {
	workersNumber := 6
	tasks := make(chan string)

	var wg sync.WaitGroup
	for i := range workersNumber {
		wg.Add(1)
		go worker(&wg, i, tasks)
	}

	urls := []string{"http://ya.ru", "http://vk.com", "http://mail.ru", "http://mail.ru"}
	for _, u := range urls {
		tasks <- u
	}
	close(tasks)

	wg.Wait()
}
