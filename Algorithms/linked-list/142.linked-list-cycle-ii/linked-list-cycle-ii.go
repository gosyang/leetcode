// Package _42_linked_list_cycle_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* linked-list-cycle-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/21 3:47 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/
package _42_linked_list_cycle_ii

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	jiaodian := hasCycle(head)
	if jiaodian == nil {
		return nil
	}
	// head 和交点一起走的交点即环入口
	cur := head
	for cur != jiaodian {
		cur = cur.Next
		jiaodian = jiaodian.Next
	}
	return cur
}

func hasCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}