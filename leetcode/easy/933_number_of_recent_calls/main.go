package main

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	// FIFO
	for len(this.requests) > 0 && t-this.Head() > 3000 {
		this.requests = this.requests[1:]
	}
	return len(this.requests)
}

func (this *RecentCounter) Head() int {
	return this.requests[0]
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
