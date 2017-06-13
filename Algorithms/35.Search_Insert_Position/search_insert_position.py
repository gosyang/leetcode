#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: search_insert_position.py
Author: gosyang
Date: 2017/05/22 10:52:40
"""


class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        # O(n)的 最直接的顺序遍历解法
        # for i in range(len(nums)):
        #     if target <= nums[i]:
        #         return i
        # return len(nums)
    
        # O(lgn) 的二分查找法, 直接排好序了已经
        low = 0
        high = len(nums) - 1
        if high < 0:
            return 0
        if target <= nums[low]:
            return low
        if target > nums[high]:
            return high + 1
        
        while high > low:
            mid = (high - low) / 2 + low
            if nums[mid] == target:
                return mid
            if target < nums[mid]:
                high = mid
            else:
                low = mid + 1
        return low

