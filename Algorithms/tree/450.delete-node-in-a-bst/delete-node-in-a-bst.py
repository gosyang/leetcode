#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: delete-node-in-a-bst
Author: gosyang
Date: 2020/2/6 4:50 PM
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
"""
# Definition for a binary tree node.


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def deleteNode(self, root, key):
        """
        :type root: TreeNode
        :type key: int
        :rtype: TreeNode
        """
        if root is None:
            return root
        if root.val > key:
            root.left = self.deleteNode(root.left, key)
        elif root.val < key:
            root.right = self.deleteNode(root.right, key)
        else:
            if root.left is None and root.right is None:
                root = None
            elif root.right is not None:
                # 如果右边有值，把右边最左边的挪过来，再删掉它
                root.val = self.successor(root.right)
                root.right = self.deleteNode(root.right, root.val)
            else:
                root.val = self.predecessor(root.left)
                root.left = self.deleteNode(root.left, root.val)
        return root

    def successor(self, root):
        """
        always left
        """
        while root.left:
            root = root.left
        return root.val

    def predecessor(self, root):
        """
        always right
        """
        while root.right:
            root = root.right
        return root.val

# 这里需要注意 self.deleteNode 是返回新的根节点，这个返回值要赋值给root
