package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	urls  []string
	resps []string
}

func NewHttpClient(urls []string) *HttpClient {
	return &HttpClient{urls: urls}
}

func (c *HttpClient) FetchAll(workersNumber int) chan string {
	jobs := make(chan string)
	resps := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < workersNumber; i++ {
		wg.Add(1)
		go worker(&wg, jobs, resps)
	}

	go func() {
		for _, u := range c.urls {
			jobs <- u
		}
		close(jobs)
		wg.Wait()
		close(resps)
	}()

	return resps
}

func worker(wg *sync.WaitGroup, jobs <-chan string, resps chan<- string) {
	defer wg.Done()
	for j := range jobs {
		resp, _ := http.Get(j)
		time.Sleep(time.Second * 1)

		resps <- fmt.Sprintf("response from %s is: %s", j, resp)
		resp.Body.Close()
	}
}

func main() {
	urls := []string{"http://ya.ru", "http://vk.com", "http://mail.ru", "http://mail.ru"}
	client := NewHttpClient(urls)
	results := client.FetchAll(3)

	for range urls {
		fmt.Printf("%s\n", <-results)
	}
}
