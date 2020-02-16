#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: four
Author: gosyang
Date: 2020/2/16 2:47 PM
"""
import heapq
class Solution(object):
    def isPossible(self, target):
        """
        :type target: List[int]
        :rtype: bool
        这种是看起来最简单的，但是用到max函数要排序，每次维护这个到10^9就会失败，因此用堆维护会好点
        """
        while True:
            max_num = max(target)
            if max_num == 1:
                return True
            max_index = target.index(max_num)
            s = sum(target)
            l = max_num - (s - max_num)
            if l <= 0:
                return False
            target[max_index] = l

    # 这个用 [] + heapq 维护堆，用负数入堆，拿出来取反实现反向排序
    def isPossible2(self, target):
        q = []
        sum = 0
        for n in target:
            heapq.heappush(q, -n)
            sum += n
        while q:
            max_num = abs(heapq.heappop(q))
            if max_num == 1:
                return True
            l = max_num - (sum - max_num)
            if l <= 0:
                return False
            sum = max_num
            heapq.heappush(q, -l)
