#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-tree-pruning
Author: gosyang
Date: 2020/2/7 11:06 AM
"""
from collections import deque

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def pruneTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root is None:
            return root
        root.left = self.pruneTree(root.left)
        root.right = self.pruneTree(root.right)
        if root.left is None and root.right is None:
            if root.val == 0:
                return None
            else:
                return root
        return root

    def pruneTree2(self, root):
        # 如果要用非递归的话，这就是个后序遍历！需要用两个栈来存
        que1 = deque([root],)
        que2 = deque()
        while que1:
            r = que1.pop()
            if r is None:
                continue
            que2.append(r)
            if r.left:
                que1.append(r.left)
            if r.right:
                que1.append(r.right)
        # 这里发现无法赋值
        while que2:
            r2 = que2.pop()
            # TODO 本来用的这种解法，问题是r2=None 不成功！
            # if r2.left is None and r2.right is None:
            #     if r2.val == 0:
            #         r2 = None
            if r2.left is not None and r2.left.val == -1:
                r2.left = None
            if r2.right is not None and r2.right.val == -1:
                r2.right = None
            if r2.left is None and r2.right is None and r2.val == 0:
                r2.val = -1
        return None if root.val == -1 else root


# 乍一看要判断很多，子都0才删，但是仔细发现，从下往上删就行了啊