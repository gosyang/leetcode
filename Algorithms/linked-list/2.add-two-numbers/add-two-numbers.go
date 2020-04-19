// Package __add_two_numbers - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* add-two-numbers - file brief introduce */
/*
modification history
-----------------------------
2020/4/14 12:42 PM, by gosyang, create
*/
/*
DESCRIPTION
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package __add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	dummy := &ListNode{
		Val: -1,
	}
	cur := dummy
	for l1 != nil || l2 != nil {
		sum := add
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		add = sum / 10
		// 插入
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
	}
	if add != 0 {
		cur.Next = &ListNode{
			Val: add,
		}
	}

	return dummy.Next
}

