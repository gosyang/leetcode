#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: diameter_of_binary_tree.py
Author: gosyang
Date: 2017/07/02 16:06:38
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def diameterOfBinaryTreeAndLongestPath(self, root):
        """
        :type root: TreeNode
        :rtype: int, int
        """
        if root is None:
            return 0, 0
        left_diameter, left_longest = self.diameterOfBinaryTreeAndLongestPath(root.left)
        right_diameter, right_longest = self.diameterOfBinaryTreeAndLongestPath(root.right)
        return max(left_diameter, right_diameter, left_longest + right_longest + 1), max(left_longest + 1, right_longest + 1)
        
    def diameterOfBinaryTree(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        diameter, longest = self.diameterOfBinaryTreeAndLongestPath(root)
        # 题意是path的数目，不是点的，所以如果是空就不用 - 1
        if diameter == 0:
            return 0
        return diameter - 1
