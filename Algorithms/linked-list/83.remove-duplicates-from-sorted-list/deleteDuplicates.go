// Package _3_remove_duplicates_from_sorted_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* deleteDuplicates - file brief introduce */
/*
modification history
-----------------------------
2019/12/20 10:37 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
*/
package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	for h != nil && h.Next != nil {
		// 如果重复
		if h.Val == h.Next.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return head
}

// 还是要想清楚，怎么移动，如果相同的话不能移动
