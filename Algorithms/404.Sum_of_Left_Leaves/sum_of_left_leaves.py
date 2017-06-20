#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: sum_of_left_leaves.py
Author: gosyang
Date: 2017/06/20 17:09:40
"""


# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def sumOfLeftLeaves(self, root):
        """
        Find the sum of all left leaves in a given binary tree.

        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0
        if root.left is None:
            return self.sumOfLeftLeaves(root.right)
        if root.left.left is None and root.left.right is None:
            return self.sumOfLeftLeaves(root.right) + root.left.val
        return self.sumOfLeftLeaves(root.left) + self.sumOfLeftLeaves(root.right)
