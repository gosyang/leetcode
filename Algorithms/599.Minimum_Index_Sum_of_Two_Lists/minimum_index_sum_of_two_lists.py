#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: minimum_index_sum_of_two_lists.py
Author: gosyang 
Date: 2017/06/12 13:41:32
"""


class Solution(object):
    def findRestaurant(self, list1, list2):
        """
        O(len(list1) + len(list2))
        :type list1: List[str]
        :type list2: List[str]
        :rtype: List[str]
        """
        list1_map = {}
        for i in range(len(list1)):
            list1_map[list1[i]] = i
        # better: list1_map = {r : ind for ind, r in enumerate(list1)}
        # print list1_map
        least = 3000
        res = []
        for i in range(len(list2)):
            if list2[i] in list1_map:
                sum = i + list1_map[list2[i]]
                if sum == least:
                    res.append(list2[i])
                elif sum < least:
                    res = [list2[i]]
                    least = sum
                    # print "find less sum [{}]: {}".format(sum, res)
        return res
