#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: convert-bst-to-greater-tree
Author: gosyang
Date: 2020/2/6 12:25 PM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    sum = 0
    def convertBST(self, root):
        """
        给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root is None:
            return root
        self.convertBST(root.right)
        root.val += self.sum
        self.sum = root.val
        self.convertBST(root.left)
        return root


# 1. 中序遍历一遍，得出所有的值，累加后，再遍历一遍，对应的改
# 2. 其实有个一个累加的值就行，右root左的遍历，不断加sum的值就是了
