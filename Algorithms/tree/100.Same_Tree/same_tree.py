#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: same_tree.py
Author: gosyang
Date: 2017/06/26 16:06:38
"""

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
from collections import deque

class Solution(object):
    def isSameTree(self, p, q):
        """
	Given two binary trees, write a function to check if they are equal or not.

	Two binary trees are considered equal if they are structurally identical and the nodes have the same value.

        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if p is not None and q is not None:
            if p.val != q.val:
                return False
            else:
                return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        elif p is None and q is None:
            return True
        else:
            return False

    def isSameTree2(self, p, q):
        """
        采用非递归的办法，使用双端队列入队列存储，比用栈简单
        """
        def check(p, q):
            if not p and not q:
                return True
            if not p or not q:
                return False
            if p.val == q.val:
                return True
            return False

        deq = deque([(p,q),])
        while deq:
            p, q = deq.popleft()
            if not check(p, q):
                return False
             
            # 如果是True要么都是NULL，要么都有，要判断下
            if p:
                deq.append([p.left, q.left])
                deq.append([p.right, q.right])
        # 不可少！！
        return True

    