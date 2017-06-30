#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: missing_number.py
Author: gosyang
Date: 2017/06/30 16:06:38
"""

class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # 求和， 等差数列求和公式 - 总和
        max = len(nums)
        return max * (max + 1) / 2 - sum(nums)