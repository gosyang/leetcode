#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: n-ary-tree-postorder-traversal
Author: gosyang
Date: 2020/2/4 4:15 PM
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
    def postorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        stack, res = deque([root],), []
        while stack:
            r = stack.pop()
            if r is None:
                continue
            res.append(r.val)
            for c in r.children:
                stack.append(c)

        return res[::-1]
# 与589 中序遍历对比！！精妙