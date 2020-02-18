// Package _34_palindrome_linked_list - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* palindrome-linked-list - file brief introduce */
/*
modification history
-----------------------------
2020/2/18 10:33 AM, by gosyang, create
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
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow在偶数时候就是一半的，奇数的就是中间那个，放在前面
	rHead := reverseList(slow.Next)
	for head != nil && rHead != nil {
		if head.Val != rHead.Val {
			return false
		}
		head = head.Next
		rHead = rHead.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
// 一次搞定，但是是看了题解了，把后面半段翻转即可，背会reverseList和熟记 快慢指针
// 快慢指针诀窍：1. 先判断0个和1个节点的情况，2. 判断next, next.next != nil，3. 结果就是slow最后会落在奇数中点和偶数n/2处

