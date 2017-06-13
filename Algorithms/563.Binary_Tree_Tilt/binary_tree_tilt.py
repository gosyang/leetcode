#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: binary_tree_tilt.py
Author: gosyang
Date: 2017/05/25 10:24:38
"""


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def sumOfRoot(self, root):
        """
        Sum val of root tree
        这种solution 重复计算了 子树的和, 应该两个递归一起搞
        
        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0
        return self.sumOfRoot(root.left) + self.sumOfRoot(root.right) + root.val

    def sumAndTilt(self, root):
        """
        Sum val of root tree and tilt at same time
        :type root: TreeNode
        :rtype: int(sum), int(tilt)
        """
        if root is None:
            return 0, 0
        left_sum, left_tilt = self.sumAndTilt(root.left)
        right_sum, right_tilt = self.sumAndTilt(root.right)
        
        return left_sum + root.val + right_sum, abs(left_sum - right_sum) + left_tilt + right_tilt
    
    def findTilt(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        # 用 sumOfRoot 递归, 耗时1000ms以上
        # if root is None:
        #     return 0
        # return self.findTilt(root.left) + self.findTilt(root.right) \
        #        + abs(self.sumOfRoot(root.left) - self.sumOfRoot(root.right))
        
        # 用一次递归
        return self.sumAndTilt(root)[1]
    
    
                                                        


