#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-search-tree-to-greater-sum-tree
Author: gosyang
Date: 2020/2/7 1:56 PM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def bstToGst(self, root):
        """
        给出二叉搜索树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
        :type root: TreeNode
        :rtype: TreeNode
        """
        stack1, stack2 = [], []
        # 中序遍历得到
        cur = root
        while stack1 or cur:
            while cur:
                stack1.append(cur)
                cur = cur.left
            top = stack1.pop()
            stack2.append(top)
            cur = top.right

        pre = 0
        while stack2:
            r = stack2.pop()
            r.val += pre
            pre = r.val
        return root

# 如果右根左的顺序遍历，就直接加就行
    def bstToGst2(self, root):
        stack = []
        pre = 0
        cur = root
        while stack or cur:
            while cur:
                stack.append(cur)
                cur = cur.right
            top = stack.pop()
            top.val += pre
            pre = top.val
            cur = top.left
        return root





# 思路就是非递归的方式中序遍历的压入栈，再弹出来累加