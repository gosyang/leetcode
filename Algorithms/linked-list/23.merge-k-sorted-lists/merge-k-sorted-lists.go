/* Copyright
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	if listsLen == 0 {
		return nil
	} else if listsLen == 1 {
		return lists[0]
	}

	mid := listsLen / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			return l1
		} else if l2 != nil {
			return l2
		} else {
			return nil
		}
	}

	var head, tail *ListNode
	if l1.Val < l2.Val {
		head, tail = l1, l1
		l1 = l1.Next
	} else {
		head, tail = l2, l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			tail = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = l2
			l2 = l2.Next
		}
	}
	for l1 != nil {
		tail.Next = l1
		tail = l1
		l1 = l1.Next
	}
	for l2 != nil {
		tail.Next = l2
		tail = l2
		l2 = l2.Next
	}

	return head
}
