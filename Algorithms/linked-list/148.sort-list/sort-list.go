// Package _48_sort_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sort-list - file brief introduce */
/*
modification history
-----------------------------
2020/2/20 4:42 PM, by gosyang, create
*/
/*
DESCRIPTION
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
*/
package _48_sort_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)

	// 合并两个排好序的链表，背会！双dummy
	res := new(ListNode)
	h := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	}
	if right != nil {
		h.Next = right
	}
	return res.Next
}

// 用于交换两个节点上的数值
func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

// 快速排序的主体函数
func quickSort(head, end *ListNode) {
	if head == nil || head.Next == nil {
		return // 头节点已经是结束节点或要排序只有一个头节点直接返回
	}
	pivot := head.Val             // 头节点的值作为pivot
	slow, fast := head, head.Next // 定义快慢指针
	for fast != nil {             // 当快指针不等于结束指针时，不断执行以下操作
		if fast.Val <= pivot { // 如果快指针小于pivot
			slow = slow.Next // 不断向后移动慢指针
			swap(slow, fast) // 交换快慢指针指向的节点值
		}
		fast = fast.Next // 快指针向后移动一位
	}
	// 循环结束后交换头节点和慢指针指向的值,把pivot放在大小两堆数值的中间
	swap(head, slow)
	quickSort(head, slow) // 递归处理pivot左右两边的链表
	quickSort(slow.Next, end)
}
