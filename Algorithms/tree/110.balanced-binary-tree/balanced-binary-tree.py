#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: balanced-binary-tree
Author: gosyang
Date: 2020/2/8 12:00 PM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def isBalanced(self, root):
        """
        给定一个二叉树，判断它是否是高度平衡的二叉树。

        本题中，一棵高度平衡二叉树定义为：
        一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
        :type root: TreeNode
        :rtype: bool
        """
        if root is None:
            return True
        res, _ = self.helper(root)
        return res

    def helper(self, root):
        if root is None:
                return True, 0
        l, lm = self.helper(root.left)
        r, rm = self.helper(root.right)
        if not l or not r:
            return False, 0
        if abs(lm - rm) > 1:
            return False, 0
        return True, max(lm, rm)+1

# 利用了python的多返回值，如果不是，可以一个返回值返回高度，如果不行高度是-1
    def isBalanced2(self, root):
        return self.helper2(root) != -1

    def helper2(self, root):
        if root is None:
            return 0
        l = self.helper2(root.left)
        r = self.helper2(root.right)
        if l == -1 or r == -1:
            return -1
        if abs(l - r) > 1:
            return -1
        return max(l, r) + 1

