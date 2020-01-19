// Package _41_linked_list_cycle - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* hasCycle - file brief introduce */
/*
modification history
-----------------------------
2019/12/20 2:23 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表，判断链表中是否有环。
O(1)的存储解决
*/
package _41_linked_list_cycle

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast.Next == slow.Next {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// 使用哈希表记录地址是个办法，判断是否访问过
// 如果要O(1)的空间复杂，使用一个每次走两次的fast，一个走1次的slow，追逐;
// 判断的时候fast快所以判断fast的边界nil问题

// 第二次做，精简不少
func hasCycle2(head *ListNode) bool {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}