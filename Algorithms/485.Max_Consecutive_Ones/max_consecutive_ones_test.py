#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: max_consecutive_ones_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import max_consecutive_ones


class MaxConsecutiveOnesTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = max_consecutive_ones.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.findMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]), 3)

    def test_one_1(self):
        self.assertEqual(self.solution.findMaxConsecutiveOnes([1]), 1)

    def test_one_0(self):
        self.assertEqual(self.solution.findMaxConsecutiveOnes([0]), 0)

    def test_max_1(self):
        self.assertEqual(self.solution.findMaxConsecutiveOnes([1 for i in range(10000)]), 10000)

    def test_max_0(self):
        self.assertEqual(self.solution.findMaxConsecutiveOnes([0 for i in range(10000)]), 0)
        
        
if __name__ == "__main__":
    unittest.main()


