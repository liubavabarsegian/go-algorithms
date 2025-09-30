package main

type MinStack struct {
	Mins  []int
	Items []int
}

func (this *MinStack) Push(val int) {
	this.Items = append(this.Items, val)

	if val < this.Mins[len(this.Mins)-1] {
		this.Mins = append(this.Mins, val)
	} else {
		this.Mins = append(this.Mins, this.Mins[len(this.Mins)-1])
	}

}

func (this *MinStack) Pop() {
	this.Mins = this.Mins[0 : len(this.Mins)-1]
	this.Items = this.Items[0 : len(this.Items)-1]
}

func (this *MinStack) Top() int {
	return this.Items[len(this.Items)-1]
}

func (this *MinStack) GetMin() int {
	return this.Mins[len(this.Mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
