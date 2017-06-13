#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: number_complement.py
Author: gosyang
Date: 2017/05/26 11:03:15
"""


class Solution(object):
    def findComplement(self, num):
        """
        :type num: int
        :rtype: int
        """
        for i in range(33):
            exp = 2 ** i
            if num / exp == 0:
                return exp - 1 - num
                #return (exp - 1) ^ num
                


