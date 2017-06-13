#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: max_consecutive_ones.py
Author: gosyang 
Date: 2017/06/05 10:16:54
"""


class Solution(object):
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = 0
        tmp = 0
        for i in nums:
            if not i:
                res = max(res, tmp)
                tmp = 0
            else:
                tmp += 1
        return max(res, tmp)
        
