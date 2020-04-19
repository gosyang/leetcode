// Package _45_add_two_numbers_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* add-two-numbers-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/14 12:16 PM, by gosyang, create
*/
/*
DESCRIPTION
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7
*/
package _45_add_two_numbers_ii

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nl1 := []int{}
	nl2 := []int{}
	ll1 := -1
	for l1 != nil {
		ll1 += 1
		nl1 = append(nl1, l1.Val)
		l1 = l1.Next
	}
	ll2 := -1
	for l2 != nil {
		ll2 += 1
		nl2 = append(nl2, l2.Val)
		l2 = l2.Next
	}
	// 从尾部前面边加边组成链表
	var head *ListNode
	// 进位
	add := 0
	for ll1 >= 0 || ll2 >= 0 {
		res := add
		// 有数就加
		if ll1 >= 0 {
			res += nl1[ll1]
			ll1 -= 1
		}
		if ll2 >= 0 {
			res += nl2[ll2]
			ll2 -= 1
		}
		add = res / 10
		res = res % 10
		cur := &ListNode{
			Val: res,
			Next: head,
		}
		head = cur
	}
	if add != 0 {
		cur := &ListNode{
			Val: add,
			Next: head,
		}
		head = cur
	}
	return head
}

// 不好的办法，相加的和很大会超出范围的，
func addTwoNumbersBad(l1 *ListNode, l2 *ListNode) *ListNode {
	nl1 := []int{}
	nl2 := []int{}
	ll1 := -1
	for l1 != nil {
		ll1 += 1
		nl1 = append(nl1, l1.Val)
		l1 = l1.Next
	}
	ll2 := -1
	for l2 != nil {
		ll2 += 1
		nl2 = append(nl2, l2.Val)
		l2 = l2.Next
	}
	n1 := 0
	n2 := 0
	for i := 0; i <= ll1; i++ {
		n1 = n1*10 + nl1[i]
	}
	for i := 0; i <= ll2; i++ {
		n2 = n2*10 + nl2[i]
	}
	ans := n1 + n2
	// 这里需要特别注意
	if ans == 0 {
		return &ListNode{
			Val: 0,
		}
	}
	var head *ListNode
	for ans > 0 {
		cur := &ListNode{
			Val: ans % 10,
			Next: head,
		}
		ans = ans / 10
		head = cur
	}
	return head
}