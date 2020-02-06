#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: binary-search-tree-iterator
Author: gosyang
Date: 2020/2/6 4:04 PM
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class BSTIterator1(object):
    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self.vals = self.inorder(root)
        self.index = 0  # 这次到谁了

    def inorder(self, root):
        """
        :param root:
        :return: []int 中序遍历值
        """
        if root is None:
            return []
        return self.inorder(root.left) + [root.val] + self.inorder(root.right)

    def next(self):
        """
        @return the next smallest number
        :rtype: int
        """
        # 下次是谁
        self.index += 1
        return self.vals[self.index-1]

    def hasNext(self):
        """
        @return whether we have a next smallest number
        :rtype: bool
        """
        return self.index < len(self.vals)

# Your BSTIterator object will be instantiated and called as such:
# obj = BSTIterator(root)
# param_1 = obj.next()
# param_2 = obj.hasNext()

# 由于是bst，先init中序遍历好就行了，但是空间复杂是O(N), 不是O(h)

class BSTIterator(object):
    def __init__(self, root):
        self.stack = []
        self._leftmost_inorder(root)

    def _leftmost_inorder(self, root):
        # For a given node, add all the elements in the leftmost branch of the tree
        # under it to the stack.
        while root:
            self.stack.append(root)
            root = root.left

    def next(self):
        topmost_node = self.stack.pop()

        if topmost_node.right:
            self._leftmost_inorder(topmost_node.right)

        return topmost_node.val

    def hasNext(self):
        return len(self.stack) > 0


