#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: minimum_index_sum_of_two_lists_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import minimum_index_sum_of_two_lists


class HammingDistanceTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = minimum_index_sum_of_two_lists.Solution()

    def tearDown(self):
        self.solution = None

    def test_example1(self):
        a = ["Shogun", "Tapioca Express", "Burger King", "KFC"]
        b = ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
        self.assertItemsEqual(self.solution.findRestaurant(a, b), ["Shogun"])

    def test_example2(self):
        a = ["Shogun", "Tapioca Express", "Burger King", "KFC"]
        b = ["KFC", "Shogun", "Burger King"]
        self.assertItemsEqual(self.solution.findRestaurant(a, b), ["Shogun"])
    
    def test_find_none(self):
        a = ["a"]
        b = ["b"]
        self.assertItemsEqual(self.solution.findRestaurant(a, b), [])

    def test_find_one(self):
        a = ["a"]
        b = ["a"]
        self.assertItemsEqual(self.solution.findRestaurant(a, b), ["a"])
        
        
if __name__ == "__main__":
    unittest.main()


