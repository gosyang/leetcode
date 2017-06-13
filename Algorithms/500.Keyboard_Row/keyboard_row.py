#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: keyboard_row.py
Author: gosyang
Date: 2017/06/01 09:44:37
"""


class Solution(object):
    def findWords(self, words):
        """
        :type words: List[str]
        :rtype: List[str]
        """
        res = []
        row_q = "qwertyuiop"
        row_a = "asdfghjkl"
        
        for word in words:
            row = None
            for letter in word:
                letter = letter.lower()
                if letter in row_q:
                    row_this = "q"
                elif letter in row_a:
                    row_this = "a"
                else:
                    row_this = "z"
                if row is None:
                    row = row_this
                elif row != row_this:
                    row = -1
                    break
            if row != -1:
                res.append(word)
        return res
    
    # 何时才能pythonic?
    # row1 = set("qwertyuiopQWERTYUIOP")
    # row2 = set("asdfghjklASDFGHJKL")
    # row3 = set("zxcvbnmZXCVBNM")
    # return [word for word in words if set(word).issubset(row1) or
    #         set(word).issubset(row2) or
    #         set(word).issubset(row3)]

