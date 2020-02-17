#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: top-k-frequent-elements
Author: gosyang
Date: 2020/2/17 10:24 AM
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
"""
import collections

class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        count = collections.Counter(nums)
        return [i[0] for i in count.most_common(k)]

# 最简单的是hash表，我们这里用heap试试
# 思路也是先统计词频了，然后普通的做法

        #count = collections.Counter(nums)
        #return heapq.nlargest(k, count.keys(), key=count.get)
