#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: maximum-depth-of-n-ary-tree
Author: gosyang
Date: 2020/2/4 3:02 PM
"""
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
from collections import deque
class Solution(object):
    def maxDepth(self, root):
        """
        :type root: Node
        :rtype: int
        """
        res = 0
        stack = deque([(1, root)],)
        while stack:
            depth, r = stack.pop()
            if r is None:
                continue
            res = max(res, depth)
            for c in r.children:
                stack.append((depth+1, c))
        return res

# 直接非递归的搞就行，bfs和dfs都可以, 关键是记录层数
