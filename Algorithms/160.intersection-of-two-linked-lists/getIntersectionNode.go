// Package _60_intersection_of_two_linked_lists - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* getIntersectionNode - file brief introduce */
/*
modification history
-----------------------------
2019/12/20 2:36 PM, by gosyang, create
*/
/*
DESCRIPTION
编写一个程序，找到两个单链表相交的起始节点。

注意：
如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
*/
package _60_intersection_of_two_linked_lists

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		pA = pA.Next
		pB = pB.Next
		if pA == nil && pB == nil {
			return nil
		}
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
	}
	return pA
}

// O（1）时间复杂度的思想是，把链表合并，之后走完两个链表，如果都走到了头，说明不相交否则相交，且是相交点