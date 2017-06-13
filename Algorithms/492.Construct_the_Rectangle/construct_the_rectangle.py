#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: construct_the_rectangle.py
Author: gosyang
Date: 2017/06/13 17:59:05
"""
import math


class Solution(object):
    def constructRectangle(self, area):
        """
        :type area: int
        :rtype: List[int]
        """
        # 这个方法超时了, 例如9999992
        # for l in range(int(math.sqrt(area)), area + 1):
            # 判断是否整除
            # if not area % l:
            #     w = area / l
                #  加入sqrt 取整值正好小, 如 2 【1, 2】
                # if l >= w:
                #     return [l, w]
        
        # 大到小的来遍历, 因为int取整的数字必然偏小, 上面的方法需要做判断,这个不用
        for l in range(int(math.sqrt(area)), 0, -1):
            if not area % l:
                return [area / l, l]
        
            
        

