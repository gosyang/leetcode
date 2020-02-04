#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: univalued-binary-tree
Author: gosyang
Date: 2020/2/4 2:55 PM
"""
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
from collections import deque
class Solution(object):
    def isUnivalTree(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True
        value = root.val
        deq = deque([root],)
        while deq:
            r = deq.pop()
            if not r:
                continue
            if r.val != value:
                return False
            deq.append(r.left)
            deq.append(r.right)
        return True

