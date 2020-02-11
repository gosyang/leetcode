#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: cousins-in-binary-tree
Author: gosyang
Date: 2020/2/11 11:01 AM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isCousins(self, root, x, y):
        """
        :type root: TreeNode
        :type x: int
        :type y: int
        :rtype: bool
        """
        if root is None or root.val == x or root.val == y:
            return False

        stack = [(root, 1, 0)]
        x_height, y_height = 0, 0
        x_father, y_father = 0, 0
        while stack:
            top, height, father = stack.pop()
            if top.val == x:
                x_height, x_father = height, father
            elif top.val == y:
                y_height, y_father = height, father

            if x_father != y_father and x_height == y_height and x_height != 0:
                return True

            if top.right is not None:
                stack.append((top.right, height+1, top.val))
            if top.left is not None:
                stack.append((top.left, height+1, top.val))
        return False


# 高度相同但父节点不同, 节点值1-100之间