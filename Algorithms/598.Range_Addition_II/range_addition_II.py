#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: range_addition_II.py
Author: gosyang
Date: 2017/06/14 10:32:32
"""


class Solution(object):
    def maxCount(self, m, n, ops):
        """
        Given an m * n matrix M initialized with all 0's and several update operations.

        Operations are represented by a 2D array, and each operation is represented
        by an array with two positive integers a and b, which means M[i][j] should
        be added by one for all 0 <= i < a and 0 <= j < b.

        You need to count and return the number of maximum integers in the matrix
        after performing all the operations.
        
        :type m: int
        :type n: int
        :type ops: List[List[int]]
        :rtype: int
        """
        # 由于一定是从固定起始, 直接找操作的所有的 最小即可
        x_min, y_min = m, n
        for op in ops:
            x_min, y_min = min(x_min, op[0]), min(y_min, op[1])
        return x_min * y_min

