#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: comtains_duplicate.py
Author: baidu(baidu@baidu.com)
Date: 2017/05/16 16:56:26
"""
class Solution(object):
    def containsDuplicate(self, nums):
        """
        用时54ms, 基本思路是 变成set 再比较len
        :type nums: List[int]
        :rtype: bool
        """
        nums.sort()
        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                return True
        return False
