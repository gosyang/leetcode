#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: subtree_of_another_tree.py
Author: gosyang
Date: 2017/07/03 16:06:38
"""

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isEqualTree(self, s, t):
        """
        :type s: TreeNode
        :type t: TreeNode
        :rtype: bool
        """
        if s is None and t is None:
            return True
        if s is not None and t is not None:
            if s.val != t.val:
                return False
            return self.isEqualTree(s.left, t.left) and self.isEqualTree(s.right, t.right)
        return False

    def isSubtree(self, s, t):
        """
        :type s: TreeNode
        :type t: TreeNode
        :rtype: bool
        """
        if self.isEqualTree(s, t):
            return True

        if s is None or t is None:
            return False

        # 拿t的整棵树作比较每次
        return self.isSubtree(s.left, t) or self.isSubtree(s.right, t)
        # return self.isSubtree(s.left, t.left) or self.isSubtree(s.right, t.right)



