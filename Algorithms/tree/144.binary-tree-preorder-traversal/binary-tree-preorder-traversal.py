#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-tree-preorder-traversal
Author: gosyang
Date: 2020/2/7 3:31 PM
"""

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def preorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        stack = [root]
        res = []
        while stack:
            r = stack.pop()
            if r is None:
                continue
            res.append(r.val)
            stack.append(r.right)
            stack.append(r.left)
        return res
