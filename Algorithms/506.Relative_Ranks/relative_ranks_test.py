#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: relative_ranks_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import relative_ranks


class RelativeRanksTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = relative_ranks.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        example = [5, 4, 3, 2, 1]
        ans = ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
        self.assertItemsEqual(self.solution.findRelativeRanks(example), ans)

    def test_empty(self):
        example = []
        ans = []
        self.assertItemsEqual(self.solution.findRelativeRanks(example), ans)

    def test_only_gold(self):
        example = [100]
        ans = ["Gold Medal"]
        self.assertItemsEqual(self.solution.findRelativeRanks(example), ans)

    def test_ordered(self):
        example = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
        ans = ["10", "9", "8", "7", "6", "5", "4",
               "Bronze Medal", "Silver Medal", "Gold Medal"]
        self.assertItemsEqual(self.solution.findRelativeRanks(example), ans)
        
if __name__ == "__main__":
    unittest.main()


