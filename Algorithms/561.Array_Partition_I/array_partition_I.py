#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: array_partition_I.py
Author: gosyang
Date: 2017/05/17 14:30:17
"""


class Solution(object):
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums.sort()
        sum = 0
        for i in range(len(nums) / 2):
            sum += nums[i * 2]
        return sum

    # sum(sorted(nums)[::2]) 利用切片来实现的!

