// Package stack - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* basic - file brief introduce */
/*
modification history
-----------------------------
2020/4/17 11:40 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package main

import "fmt"

// 初始化
var queue []int
var stack []int

func main() {
	// 入队 入栈
	queue = append(queue, 1)
	stack = append(stack, 1)
	fmt.Println(queue, stack)

	queue = append(queue, 2)
	stack = append(stack, 2)

	// 出队 出栈
	queue = queue[1:]
	stack = stack[0: len(queue)]
	fmt.Println(queue, stack)
}
