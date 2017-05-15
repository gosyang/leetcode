#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: number_of_1_bits.py
Author: gosyang
Date: 2017/05/15 18:22:40
"""
class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        # 这里容易出错, 如果ret = 0, 对于 n = 1 来说就会忽略!
        ret = n % 2
        while n / 2 != 0:
            n /= 2
            ret += n % 2
        return ret
