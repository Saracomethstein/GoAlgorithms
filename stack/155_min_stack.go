package stack

// link to the task on leetcode
// https://leetcode.com/problems/min-stack/description/

type MinStack struct {
	min int
	top *StackNode
}

type StackNode struct {
	data    int
	lastMin int
	next    *StackNode
}

var newTop *StackNode

func Constructor() MinStack {
	return MinStack{top: nil, min: 0}
}

func (this *MinStack) Push(val int) {
	if this.top == nil {
		newTop = &StackNode{data: val, next: this.top}
		this.min = val
	} else {
		newTop = &StackNode{data: val, lastMin: this.min, next: this.top}
	}

	this.top = newTop

	if this.top.data < this.min {
		this.min = this.top.data
	}
}

func (this *MinStack) Pop() {
	if this.top.next == nil {
		this.top = nil
		return
	}

	this.min = this.top.lastMin
	*this.top = *this.top.next
}

func (this *MinStack) Top() int {
	return this.top.data
}

func (this *MinStack) GetMin() int {
	return this.min
}
