// Package _06_Reverse_Linked_List - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* reverseList - file brief introduce */
/*
modification history
-----------------------------
2019/12/20 3:21 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _06_Reverse_Linked_List

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	rHead := &ListNode{
		Val: head.Val,
		Next: nil,
	}
	head = head.Next
	for head != nil {
		tmp := head.Next
		head.Next = rHead
		rHead = head
		head = tmp
	}
	return rHead
}
// 翻转链表，考虑nil，单个的情况，之后就是从第二个开始
// 完全可以用 var 生命一个空指针啊！！