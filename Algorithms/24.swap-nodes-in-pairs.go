// Package Algorithms - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* 24.swap-nodes-in-pairs - file brief introduce */
/*
modification history
-----------------------------
2019/12/19 6:39 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
package Algorithms

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs(res.Next)
	res.Next = head

	return res
}

// 递归想法，关键是处理head的前驱节点想清楚