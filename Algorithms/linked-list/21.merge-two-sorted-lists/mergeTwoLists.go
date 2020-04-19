// Package _1_merge_two_sorted_lists - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* mergeTwoLists - file brief introduce */
/*
modification history
-----------------------------
2019/12/19 4:40 PM, by gosyang, create
*/
/*
DESCRIPTION
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// 递归解法，类似归并排序 时间空间O(m+n)

func mergeTwoListsGood(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 这里直接把连接连起来？
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return dummy.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	cur := dummy
	for l1 != nil || l2 != nil {
		var n int
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				n = l1.Val
				l1 = l1.Next
			} else {
				n = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			n = l1.Val
			l1 = l1.Next
		} else {
			n = l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}
	return dummy.Next
}