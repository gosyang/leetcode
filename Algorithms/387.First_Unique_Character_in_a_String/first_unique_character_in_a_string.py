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
Date: 2017/06/23 16:06:38
"""

import collections


class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        map = collections.Counter(s)
        for i, v in enumerate(s):
            if map[v] == 1:
                return i
        return -1

        # best:
	# >>> import string
        # >>> string.ascii_lowercase
        # 'abcdefghijklmnopqrstuvwxyz'
        # return min([s.find(c) for c in string.ascii_lowercase if s.count(c) == 1] or [-1])
