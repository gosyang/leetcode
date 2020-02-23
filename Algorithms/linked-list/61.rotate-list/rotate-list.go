// Package _1_rotate_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* rotate-list - file brief introduce */
/*
modification history
-----------------------------
2020/2/18 10:04 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

*/
package _1_rotate_list

/*
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length += 1
	}
	last := k % length
	// 如果挪完没变
	if last == 0 {
		return head
	}
	cur = head
	end := head
	for i := 0; i < length-last; i++ {
		end = cur
		cur = cur.Next
	}
	end.Next = nil
	newHead := cur
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return newHead
}

// 如果每次走一步，显然要每次遍历到最后一个，不如先遍历一遍得到长度 再 % k，得到第几个挪过来
// TODO 需要注意的点：1. head是0，%的length不能是0，因此这里要先判断head nil； 2. 如果一看k是长度，那就没必要直接返回了，否则后面还要判断最后一个情况更麻烦

// 优化一下的话，就是第一次找长度的时候就找到尾指针，然后直接连过来，不用再有最后一步找了
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	cur := head
	end := head
	for cur != nil {
		end = cur
		cur = cur.Next
		length += 1
	}
	last := k % length
	// 如果挪完没变
	if last == 0 {
		return head
	}

	// 挪到最后一个就行了
	cur = head
	end.Next = head
	for i := 0; i < length-last-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}
