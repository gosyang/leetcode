#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: print-binary-tree
Author: gosyang
Date: 2020/2/10 10:24 AM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def printTree(self, root):
        """
        输出二叉树
        :type root: TreeNode
        :rtype: List[List[str]]
        """

