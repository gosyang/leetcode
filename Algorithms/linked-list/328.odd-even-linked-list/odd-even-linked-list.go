// Package _28_odd_even_linked_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* odd-even-linked-list - file brief introduce */
/*
modification history
-----------------------------
2020/4/16 8:35 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/
package _28_odd_even_linked_list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
// 自己想的，没有答案那么精炼
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := &ListNode{Val:0}
	cur := head
	evenCur := evenHead
	for cur != nil && cur.Next != nil {
		evenCur.Next = cur.Next
		evenCur = evenCur.Next
		cur.Next = cur.Next.Next
		// 如果偶数个数，cur要连后面的偶数，不能把它变成nil
		if cur.Next == nil {
			break
		}
		cur = cur.Next
		evenCur.Next = nil
	}
	cur.Next = evenHead.Next
	return head
}

func oddEvenListGood(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	even := evenHead
	odd := head
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}