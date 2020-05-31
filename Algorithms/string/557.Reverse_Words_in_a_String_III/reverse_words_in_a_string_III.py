#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_words_in_a_string_III.py
Author: gosyang
Date: 2017/05/30 23:50:57
"""


class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return " ".join(map(lambda x: x[::-1], s.split(" ")))
        
        # others
        # return ' '.join(x[::-1] for x in s.split(' '))
        # return ' '.join(s.split(' ')[::-1])[::-1]
