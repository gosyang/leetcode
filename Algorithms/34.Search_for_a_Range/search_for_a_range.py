#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: search_for_a_range.py
Author: gosyang
Date: 2017/05/23 10:08:09
"""


class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        low = 0
        high = len(nums) - 1
        if high < low:
            return [-1, -1]
        
        while low < high:
            if nums[low] == target and nums[high] == target:
                return [low, high]
            mid = (high - low) / 2 + low
            if nums[mid] == target:
                # 搞了半天的原因在这儿, 一开始没有弄对输入是一个切片, 而是[low, mid] 低级错误
                # 对于切片是相对位置了!
                left = self.searchRange(nums[low: mid], target)
                right = self.searchRange(nums[mid + 1: high + 1], target)
                if left[0] == -1:
                    l = mid
                else:
                    l = low + left[0]
                if right[-1] == -1:
                    r = mid
                else:
                    r = mid + right[-1] + 1
                return [l, r]
                
            elif nums[mid] < target:
                low = mid + 1
            else:
                high = mid - 1
                
        if nums[low] != target:
            return [-1, -1]
        else:
            return [low, high]
            
            
        
        

