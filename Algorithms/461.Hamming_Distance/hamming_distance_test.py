#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: hamming_distance_test.py
Author: baidu(baidu@baidu.com)
Date: 2017/05/15 16:10:09
"""
import unittest

import hamming_distance


class HammingDistanceTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = hamming_distance.Solution()

    def tearDown(self):
        self.solution = None

    def test_x_y_all_zero(self):
        self.assertEqual(self.solution.hammingDistance(0, 0), 0)

    def test_x_y_1_4(self):
        self.assertEqual(self.solution.hammingDistance(1, 4), 2)

    def test_x_y_all_max(self):
        x = 2 * 1024 * 1024 * 1024 - 1
        y = 2 * 1024 * 1024 * 1024 - 1
        self.assertEqual(self.solution.hammingDistance(x, y), 0)
        
        
if __name__ == "__main__":
    unittest.main()


