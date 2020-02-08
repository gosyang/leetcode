#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-tree-inorder-traversal
Author: gosyang
Date: 2020/2/7 2:42 PM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        res = []
        stack = []
        cur = root
        while stack or cur:
            while cur is not None:
                stack.append(cur)
                cur = cur.left
            top = stack.pop()
            res.append(top.val)
            cur = top.right
        return res

# 由于depque没有peek，这个取不删的操作，引入cur来简化代码吧
