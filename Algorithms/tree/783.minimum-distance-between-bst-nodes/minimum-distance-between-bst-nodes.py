#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: minimum-distance-between-bst-nodes
Author: gosyang
Date: 2020/2/7 10:41 AM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def minDiffInBST(self, root):
        """
        给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。
        注意：二叉树的大小范围在 2 到 100。二叉树总是有效的，每个节点的值都是整数，且不重复。
        :type root: TreeNode
        :rtype: int
        """
        ans = float('inf')
        vals = self.inorder(root)
        for i in range(1, len(vals)):
            ans = min(ans, vals[i]-vals[i-1])
        return ans

    def inorder(self, root):
        if root is None:
            return []
        return self.inorder(root.left) + [root.val] + self.inorder(root.right)


# 似乎没有更好的解法，就是中序遍历，得到结果后找出相邻的最小即可
