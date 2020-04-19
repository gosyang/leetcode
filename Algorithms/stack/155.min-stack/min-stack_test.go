// Package _55_min_stack - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* min-stack_test - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 5:40 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _55_min_stack

import (
	"testing"
	"fmt"
)

func TestMinStack(t *testing.T) {
	// ["MinStack","push","push","push","min","top","pop","min"]
	// [[],[-2],[0],[-1],[],[],[],[]]
	// 预期 [null,null,null,null,-2,-1,null,-2]
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.Min())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Min())

	// ["MinStack","push","push","push","min","pop","min"]
	// [[],[0],[1],[0],[],[],[]]
	obj = Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	fmt.Println(obj.Min())
	fmt.Println(obj.min, obj.stack)
	obj.Pop()
	fmt.Println(obj.min, obj.stack)
	fmt.Println(obj.Min())
}
