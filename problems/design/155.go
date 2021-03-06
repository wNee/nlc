package main

type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if x < this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data) -1 ]
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return -1
	}
	return this.min[len(this.min)-1]
}
