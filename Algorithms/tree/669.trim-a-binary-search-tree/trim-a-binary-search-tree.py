#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: trim-a-binary-search-tree
Author: gosyang
Date: 2020/2/6 10:59 AM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def trimBST(self, root, L, R):
        """
        给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。
        你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

        :type root: TreeNode
        :type L: int
        :type R: int
        :rtype: TreeNode
        """
        if root is None:
            return root
        if root.val < L:
            return self.trimBST(root.right, L, R)
        if root.val > R:
            return self.trimBST(root.left, L, R)
        root.left = self.trimBST(root.right, L, R)
        root.right = self.trimBST(root.right, L, R)
        return root

# 这种题乍一看好难呀还要换根，但其实如果不在范围内，那左子树或右子树就可以完全舍弃了
