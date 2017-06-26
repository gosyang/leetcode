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

