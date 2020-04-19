// Package _46_validate_stack_sequences - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* validate-stack-sequences - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 4:42 PM, by gosyang, create
*/
/*
DESCRIPTION
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。
*/
package _46_validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for i := 0; i < len(pushed); i++ {
		// 压入栈
		stack = append(stack, pushed[i])
		// 如果此时栈顶元素是poped的出栈元素，就弹出，并一直弹出
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}