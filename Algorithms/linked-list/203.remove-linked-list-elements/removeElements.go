// Package _03_remove_linked_list_elements - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* removeElements - file brief introduce */
/*
modification history
-----------------------------
2019/12/20 3:11 PM, by gosyang, create
*/
/*
DESCRIPTION
删除链表中等于给定值 val 的所有节点。

*/
package _03_remove_linked_list_elements

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}

// 递归解法一次AC
