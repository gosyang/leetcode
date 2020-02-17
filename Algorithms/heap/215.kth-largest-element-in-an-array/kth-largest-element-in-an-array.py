#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: kth-largest-element-in-an-array
Author: gosyang
Date: 2020/2/17 9:51 AM
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
"""
import heapq

class Solution(object):
    def findKthLargest(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        q = []
        for n in nums:
            if len(q) < k:
                heapq.heappush(q, n)
            elif q[0] < n:
                heapq.heapreplace(q, n)
        return heapq.heappop(q)

# heapq是小根堆，里面放的是前k大的值
