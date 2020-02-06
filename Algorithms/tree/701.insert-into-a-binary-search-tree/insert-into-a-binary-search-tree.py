#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: submissions
Author: gosyang
Date: 2020/2/6 3:53 PM
"""

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def insertIntoBST(self, root, val):
        """
        :type root: TreeNode
        :type val: int
        :rtype: TreeNode
        """
        if root is None:
            return TreeNode(val)
        if root.val > val:
            root.left = self.insertIntoBST(root.left, val)
        else:
            root.right = self.insertIntoBST(root.right, val)
        return root

# 递归的插就完了