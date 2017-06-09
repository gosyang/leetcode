#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: find_all_numbers_disappeared_in_an_array.py
Author: gosyang 
Date: 2017/06/09 10:47:37
"""


class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        击败 99.73 % ! O(2n)
        :type nums: List[int]
        :rtype: List[int]
        """
        res = []
        s = set(nums)
        for i in range(len(nums)):
            if i + 1 not in s:
                res.append(i + 1)
        return res
        


