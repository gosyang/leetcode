#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: move_zeroes.py
Author: gosyang
Date: 2017/06/12 16:56:41
"""

class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        zero_list = [index for index, value in enumerate(nums) if value == 0]
        
        for i in reversed(zero_list):
            del nums[i]
        
        nums += [0 for i in range(len(zero_list))]
            
        # 或者直接扫描, 两个指针, 快的非零值赋给慢的,最后再补0, 不要有把0移动的想法,那样太慢
        # l = 0
        # for r in range(len(nums)):
        #     if nums[r] != 0:
        #         nums[l] = nums[r]
        #         l += 1
        # for l in range(l, len(nums)):
        #     nums[l] = 0


