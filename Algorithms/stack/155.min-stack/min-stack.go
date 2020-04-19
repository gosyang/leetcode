// Package _55_min_stack - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* min-stack - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 5:36 PM, by gosyang, create
*/
/*
DESCRIPTION
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

*/
package _55_min_stack

type MinStack struct {
	stack []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	// 这里必须是 <= ，如果是存index可以但略显繁琐用<，010，删掉0的话，还有最小值第一个0
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	if this.stack[len(this.stack)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */