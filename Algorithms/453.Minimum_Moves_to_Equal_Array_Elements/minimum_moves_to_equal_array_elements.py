#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: minimum_moves_to_equal_array_elements.py
Author: gosyang
Date: 2017/06/18 11:10:34
"""


class Solution(object):
    def minMoves(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # 排序后, 从小到大挨着消除差异, 同是这个次数会累计到后面的差异中,因此需要累计
        nums.sort()
        sum, diff_sum = 0, 0
        lens = len(nums)
        for i in range(lens - 1):
            diff_sum += nums[i + 1] - nums[i]
            sum += diff_sum
        return sum
        
        # 数学问题, sum + m * (n - 1) = x * n, sum是当前的总和, 移动m后最后都变成x,
        # 由于 $最小的肯定每次都加$关键$,  x = minNum + m, 所以得到以下答案,
        # return sum(nums) - min(nums) * len(nums)

