#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: validate-binary-search-tree
Author: gosyang
Date: 2020/2/5 11:51 AM

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
"""

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.helper(root)

    def helper(self, root, lower=float('-inf'), upper=float('inf')):
        if root is None:
            return True
        val = root.val
        if val <= lower or val >= upper:
            return False
        if not self.helper(root.left, lower, val):
            return False
        if not self.helper(root.right, val, upper):
            return False
        return True



# 1.中序遍历出的结果如果是递增的就是，否则不是
# 2.递归判断 解法是看了答案的，不需要返回值来判断