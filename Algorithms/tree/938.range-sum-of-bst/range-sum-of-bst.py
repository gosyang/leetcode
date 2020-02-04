#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: range-sum-of-bst.py
Author: gosyang
Date: 2020/02/04 16:06:38
"""
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from collections import deque

class Solution(object):
    def rangeSumBST(self, root, L, R):
        """
        :type root: TreeNode
        :type L: int
        :type R: int
        :rtype: int
        """
        res = 0
        deq = deque([root],)
        while deq:
            # root = deq.popleft()
            root = deq.pop()
            if not root:
                continue
            if root.val > R:
                deq.append(root.left)
                continue
            if root.val < L:
                deq.append(root.right)
                continue
            res += root.val
            deq.append(root.left)
            deq.append(root.right)

        return res
        
        # 非递归的话，用deq模拟栈，popleft就是宽搜了,也可以，差不多 