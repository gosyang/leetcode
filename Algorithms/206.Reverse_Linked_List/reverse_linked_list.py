#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_linked_list.py
Author: gosyang
Date: 2017/05/18 14:22:38
"""
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # 如果是None则 没有next方法, 考虑
        if head is None:
            return None
        #
        reverse_head = head
        pre_head = None
        while reverse_head.next is not None:
            tmp_node = reverse_head
            reverse_head = reverse_head.next
            tmp_node.next = pre_head
            pre_head = tmp_node
        reverse_head.next = pre_head
        return reverse_head
            
            
            

