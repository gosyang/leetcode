// Package _290_convert_binary_number_in_a_linked_list_to_integer - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* getDecimalValue - file brief introduce */
/*
modification history
-----------------------------
2019/12/19 2:34 PM, by gosyang, create
*/
/*
DESCRIPTION
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。

请你返回该链表所表示数字的 十进制值 。
*/
package main

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := head.Val
	for head.Next != nil {
		ans <<= 1
		head = head.Next
		ans += head.Val
	}
	return ans
}

// 这种二进制十进制一定是使用位运算，链表可以一位一位的移动，几乎最优解
// 必须关注一次性通过，for的判断条件需要关注最后一个的值能否加上
