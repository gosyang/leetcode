#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: product-of-the-last-k-numbers
Author: gosyang
Date: 2020/2/16 12:16 PM
"""
class ProductOfNumbers(object):

    def __init__(self):
        self.a = [1]

    def add(self, num):
        """
        :type num: int
        :rtype: None
        """
        if num == 0:
            self.a = [1]
            return
        self.a.append(num * self.a[-1])


    def getProduct(self, k):
        """
        :type k: int
        :rtype: int
        """
        if k >= len(self.a):
            return 0
        return self.a[-1] // self.a[-k-1]



# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)