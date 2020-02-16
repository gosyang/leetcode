#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: zui-xiao-de-kge-shu-lcof
Author: gosyang
Date: 2020/2/16 4:46 PM
"""
import heapq
class Solution(object):
    def getLeastNumbers(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: List[int]
        输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
        0 <= k <= arr.length <= 1000
        0 <= arr[i] <= 1000
        """
        if k == 0 or len(arr) == 0:
            return []
        q = []
        for n in arr:
            if len(q) < k:
                heapq.heappush(q, -n)
            elif -n > q[0]:
                heapq.heapreplace(q, -n)

        return [-i for i in q]

# 这里注意，k是可以为0的




