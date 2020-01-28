// Package _76_middle_of_the_linked_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* middleNode - file brief introduce */
/*
modification history
-----------------------------
2019/12/23 11:06 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
*/
package _76_middle_of_the_linked_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && slow != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 一般这种中点问题就是快慢指针，从头开始
