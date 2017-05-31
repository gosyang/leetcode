#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: distribute_candies.py
Author: gosyang
Date: 2017/05/31 09:47:13
"""
import  collections


class Solution(object):
    def distributeCandies(self, candies):
        """
        :type candies: List[int]
        :rtype: int
        """
        counter = collections.Counter(candies)
        # 按照value 倒序排列, 算出谁出现最多次
        # l = sorted(counter.items(), lambda x, y: cmp(x[1], y[1]), reverse=True)
        return min(len(counter), len(candies) / 2)
    
        # 或者不算仔细直接用set
        # return min(len(set(candies)), len(candies) / 2)
    
        
        

