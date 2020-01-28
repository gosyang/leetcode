// Package _9_remove_nth_node_from_end_of_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* removeNthfromEnd - file brief introduce */
/*
modification history
-----------------------------
2019/12/19 3:24 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
你能尝试使用一趟扫描实现吗？
*/
package main

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	nodeSlice := []*ListNode{head}
	for head.Next != nil {
		head = head.Next
		nodeSlice = append(nodeSlice, head)
	}
	delIndex := len(nodeSlice) - n
	if delIndex == 0 {
		return nodeSlice[delIndex].Next
	}
	nodeSlice[delIndex-1].Next = nodeSlice[delIndex].Next
	return nodeSlice[0]
}

// 不错一次通过，空间换时间
// 节省空间的办法就是存两个，第一个走出去n步后再走第二个

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// n一定有效，所以只删一个也是nil
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			break
		}
	}
	// 意味着就是删第一个
	if fast == nil {
		return head.Next
	}
	// 继续往前走到底
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
