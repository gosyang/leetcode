#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: kth-largest-element-in-a-stream
Author: gosyang
Date: 2020/2/16 4:18 PM
设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。

你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。

你可以假设 nums 的长度≥ k-1 且k ≥ 1。

链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
"""
import heapq

class KthLargest(object):

    def __init__(self, k, nums):
        """
        :type k: int
        :type nums: List[int]
        """
        self.k = k
        self.heap = []
        for n in nums:
            if len(self.heap) < k:
                heapq.heappush(self.heap, n)
            else:
                if n > self.heap[0]:
                    heapq.heappop(self.heap)
                    heapq.heappush(self.heap, n)

    def add(self, val):
        """
        :type val: int
        :rtype: int
        """
        # 由于一开始初始化的不一定填满，且看条件基本就是差一个，因此这里要做填的操作
        if len(self.heap) < self.k:
            heapq.heappush(self.heap, val)
            return self.heap[0]

        if val > self.heap[0]:
            heapq.heappop(self.heap)
            heapq.heappush(self.heap, val)
        return self.heap[0]


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)