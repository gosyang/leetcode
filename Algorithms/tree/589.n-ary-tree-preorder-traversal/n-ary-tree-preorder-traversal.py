#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: n-ary-tree-preorder-traversal
Author: gosyang
Date: 2020/2/4 3:57 PM
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
    def preorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        res = []
        stack = deque([root],)
        while stack:
            r = stack.pop()
            if r is None:
                continue
            res.append(r.val)
            for c in reversed(r.children):
                stack.append(c)
        return res

# 这里关键在r.children的如何加入，深度优先需要反着入，r.children[::-1]

