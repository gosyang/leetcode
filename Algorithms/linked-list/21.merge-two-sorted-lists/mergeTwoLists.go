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
