#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-tree-postorder-traversal
Author: gosyang
Date: 2020/2/7 11:45 AM
"""
from collections import deque
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        stack = deque([root],)
        res = []
        while stack:
            r = stack.pop()
            if r is None:
                continue
            res.append(r.val)
            stack.append(r.left)
            stack.append(r.right)
        return res[::-1]

# 后序遍历的结果要先pop出来后存起来，最后倒序的输出，拿个栈存也行
# 结合814题，删掉为0的子树一样，从下找，也是后序遍历
