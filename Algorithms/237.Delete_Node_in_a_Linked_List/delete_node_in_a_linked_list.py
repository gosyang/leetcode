#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"
File: delete_node_in_a_linked_list.py
Author: gosyang
Date: 2017/06/25 16:06:38
"""

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def deleteNode(self, node):
        """
	Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3, the linked list should become 1 -> 2 -> 4 after calling your function.

        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
	# 异常的先判断
        if node is None:
            return 
        node.val = node.next.val
        node.next = node.next.next
