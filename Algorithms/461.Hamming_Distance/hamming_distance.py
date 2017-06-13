#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: Hamming_Distance.py
Author: gosyang
Date: 2017/05/15 16:06:38
"""

class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        x_mod, y_mod = x, y
        ret = 0
        for i in range(32):
            if x_mod == 0 and y_mod == 0:
                break
            if x_mod % 2 != y_mod % 2:
                ret += 1
            x_mod /= 2
            y_mod /= 2
        return ret

if __name__ == "__main__":
    print "461"

