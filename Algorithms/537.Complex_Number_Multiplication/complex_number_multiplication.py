#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: complex_number_multiplication.py
Author: gosyang 
Date: 2017/06/13 17:17:28
"""

class Solution(object):
    def complexNumberMultiply(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        a_l = a.split("+")
        b_l = b.split("+")
        a_shi, a_xu = int(a_l[0]), int(a_l[1][:-1])
        b_shi, b_xu = int(b_l[0]), int(b_l[1][:-1])
        r_shi = a_shi * b_shi - a_xu * b_xu
        r_xu = a_shi * b_xu + a_xu * b_shi
        return "{}+{}i".format(r_shi, r_xu)

