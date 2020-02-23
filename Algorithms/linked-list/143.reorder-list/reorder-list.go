// Package _43_reorder_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* reorder-list - file brief introduce */
/*
modification history
-----------------------------
2020/2/19 8:43 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
package _43_reorder_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	other := reverseList(slow.Next)
	cur := head
	for other != nil {
		tmp1 := cur.Next
		tmp2 := other.Next
		cur.Next = other
		other.Next = tmp1
		cur = tmp1
		other = tmp2
	}
	// 这里很好，other是新的头，所以旧的要干掉； 或者在other得到后, slow.Next = nil
	if cur != nil {
		cur.Next = nil
	}
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 看起来就是拆成两半，技术的话中间那个在最后，然后后半段reverse，合并！
// 很想汇文列表啊234题
// 一次AC！关键是想清楚，优雅的互相替换，
