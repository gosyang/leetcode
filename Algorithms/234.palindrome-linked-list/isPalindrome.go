// Package _34_palindrome_linked_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* isPalindrome - file brief introduce */
/*
modification history
-----------------------------
2019/12/23 11:22 AM, by gosyang, create
*/
/*
DESCRIPTION
请判断一个链表是否为回文链表。

你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
package _34_palindrome_linked_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	rhead := head
	rhead.Next = nil
	for fast.Next != nil {
		slow = slow.Next
	}
}

// 快慢指针找到中点指针，并翻转前半段，做对比