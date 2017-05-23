#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: search_for_a_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import search_for_a_range


class SearchForARangeTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = search_for_a_range.Solution()

    def tearDown(self):
        self.solution = None

    def test_sample(self):
        self.assertItemsEqual(self.solution.searchRange([5, 7, 7, 8, 8, 10], 8), [3, 4])

    def test_empty(self):
        self.assertItemsEqual(self.solution.searchRange([], 10), [-1, -1])

    def test_single_one_right(self):
        self.assertItemsEqual(self.solution.searchRange([10], 10), [0, 0])

    def test_single_one_wrong(self):
        self.assertItemsEqual(self.solution.searchRange([1], 10), [-1, -1])

    def test_all_same(self):
        self.assertItemsEqual(self.solution.searchRange([1, 1, 1], 1), [0, 2])

    def test_single_one_low(self):
        self.assertItemsEqual(self.solution.searchRange([1, 2, 3], 1), [0, 0])

    def test_single_one_high(self):
        self.assertItemsEqual(self.solution.searchRange([1, 2, 3], 3), [2, 2])

    def test_same_middle(self):
        self.assertItemsEqual(self.solution.searchRange([1, 2, 2, 3], 2), [1, 2])

    def test_two_low(self):
        self.assertItemsEqual(self.solution.searchRange([1, 3], 1), [0, 0])
        
if __name__ == "__main__":
    unittest.main()


