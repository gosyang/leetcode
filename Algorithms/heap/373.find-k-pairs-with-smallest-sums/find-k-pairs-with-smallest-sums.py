#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: find-k-pairs-with-smallest-sums
Author: gosyang
Date: 2020/2/17 1:23 PM
给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。

找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
"""
import heapq

class Solution(object):
    def kSmallestPairs(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[List[int]]
        """
        q = []
        for x in nums1:
            for y in nums2:
                q += [([x, y], x+y)]
        return [p[0] for p in heapq.nsmallest(k, q, key=lambda s: s[1])]

# 这个可以扩展详细写过程，包括不满k的话填，满的话比较，replace。等等
