#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: find-mode-in-binary-search-tree
Author: gosyang
Date: 2020/2/8 11:28 AM
"""
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def findMode(self, root):
        """
        给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
        假定 BST 有如下定义：

        结点左子树中所含结点的值小于等于当前结点的值
        结点右子树中所含结点的值大于等于当前结点的值
        左子树和右子树都是二叉搜索树
        :type root: TreeNode
        :rtype: List[int]
        """
        res = []
        stack = []
        cur = root
        max_num = 1
        cur_num = 1
        pre_val = float('inf')
        while stack or cur:
            while cur:
                stack.append(cur)
                cur = cur.left
            top = stack.pop()

            # 开始
            if top.val != pre_val:
                cur_num = 1
                pre_val = top.val
            else:
                cur_num += 1

            if cur_num == max_num:
                res.append(top.val)
            elif cur_num > max_num:
                max_num = cur_num
                res = [top.val]

            # 结束

            cur = top.right
        return res

# 不使用额外的空间,
# 中序遍历后，递增数列，逐个比较，超过最大次数的，就清空之前的重新赋值，这里要注意不要丢了cur=top.right
# 易错：比较最最大值，获取更新等都要注意 52行
# 优化：初始pre是inf，这样第一个必然是不等于，少在while里写判断if len(res)==0