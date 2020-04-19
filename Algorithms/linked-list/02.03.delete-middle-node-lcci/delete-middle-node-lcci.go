// Package _2_03_delete_middle_node_lcci - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* delete-middle-node-lcci - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 5:06 PM, by gosyang, create
*/
/*
DESCRIPTION
实现一种算法，删除单向链表中间的某个节点（除了第一个和最后一个节点，不一定是中间节点），假定你只能访问该节点。

示例：
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
*/
package _2_03_delete_middle_node_lcci

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 题意就是不能访问他的前节点，做法是，把d的值付到c，再把c连到e，相当于删了d