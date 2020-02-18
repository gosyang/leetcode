// Package _6_partition_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* partition-list - file brief introduce */
/*
modification history
-----------------------------
2020/2/18 11:07 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。
*/
package _6_partition_list

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var small, big, nS, nB *ListNode
	for head != nil {
		if head.Val < x {
			if small == nil {
				small = head
				nS = head
			} else {
				small.Next = head
				small = small.Next
			}
		} else {
			if big == nil {
				big = head
				nB = head
			} else {
				big.Next = head
				big = big.Next
			}
		}
		// TODO 注意这里，移动了head，必须把head的next变为nil
		tmp := head.Next
		head.Next = nil
		head = tmp
	}
	// TODO 必须考虑small是空的情况
	if small == nil {
		return nB
	}
	small.Next = nB
	return nS
}

// 新的两个指针，把cur移过来，注意扫尾，注意Next的地方是否可能为nil的引用！！
// 答案用了个哑节点，比较简便
func partition2(head *ListNode, x int) *ListNode {
	small := &ListNode{
		Val: 0,
	}
	big := &ListNode{
		Val: 0,
	}
	smallHead := small
	bigHead := big
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	// TODO 这里简化，同样的这里是最后做了一次扫尾，这里的big如果不是最后一次的head，可能指向了small，成环了
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}