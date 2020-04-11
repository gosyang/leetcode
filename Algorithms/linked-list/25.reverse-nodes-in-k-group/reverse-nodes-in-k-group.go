// Package _5_reverse_nodes_in_k_group - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* reverse-nodes-in-k-group - file brief introduce */
/*
modification history
-----------------------------
2020/4/11 6:16 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _5_reverse_nodes_in_k_group

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a := head
	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}