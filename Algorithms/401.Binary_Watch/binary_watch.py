#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: binary_watch.py
Author: gosyang
Date: 2017/05/19 14:17:36
"""


class Solution(object):
    def bitCount(self, num):
        res = num % 2
        while num / 2 > 0:
            num /= 2
            res += num % 2
        return res
    
    def readBinaryWatch(self, num):
        """
        :type num: int
        :rtype: List[str]
        """
        # 审题不严,如果是全灭就是0:00, 不是[]
        if num == 0:
            return ["0:00"]
        if num > 8:
            return []
        res = []
        for hour in xrange(12):
            for minute in xrange(60):
                if self.bitCount(hour) + self.bitCount(minute) == num:
                    if minute < 10:
                        minute = "0{}".format(minute)
                    res.append("{}:{}".format(hour, minute))
        return res
                    
            
        
        
            


