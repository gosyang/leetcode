#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: minimum_absolute_difference_in_BST.py
Author: gosyang 
Date: 2017/06/15 10:06:05
"""

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def __init__(self):
        self.list = []
    
    def inorder_tree(self, root):
        if root is None:
            return
        self.inorder_tree(root.left)
        self.list.append(root.val)
        self.inorder_tree(root.right)
    
    def getMinimumDifference(self, root):
        """
        Given a binary search tree with non-negative values,
        find the minimum absolute difference between values of any two nodes.
        
        :type root: TreeNode
        :rtype: int
        """
        self.inorder_tree(root)
        # 题中给定是一个『二叉搜索树』, 左 < val < 右, 因此中序遍历完就是排好序的, 不用排否则需要排
        # self.list.sort()
        # 至少有两个, 如果不是 res = sys.maxint
        res = self.list[1] - self.list[0]
        for i in range(len(self.list) - 1):
            diff = self.list[i + 1] - self.list[i]
            if diff < res:
                res = diff
        return res

