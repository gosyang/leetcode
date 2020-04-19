// Package design - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* dui-lie-de-zui-da-zhi-lcof - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 7:59 PM, by gosyang, create
*/
/*
DESCRIPTION
阿里考过
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1
*/
package design

type MaxQueue struct {
	queue []int
	max []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		queue: []int{},
		max: []int{},
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	// 更新max递减队列，将value插入并删除比他小的值
	for i := 0; i < len(this.max); i++ {
		if value > this.max[i] {
			this.max[i] = value
			this.max = this.max[:i+1]
			return
		}
	}
	this.max = append(this.max, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	if res == this.max[0] {
		this.max = this.max[1:]
	}
	return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */