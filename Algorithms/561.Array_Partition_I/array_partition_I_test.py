#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: array_partition_I.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import array_partition_I


class ArrayPartitionITestCase(unittest.TestCase):
    def setUp(self):
        self.solution = array_partition_I.Solution()

    def tearDown(self):
        self.solution = None

    def test_sample_case(self):
        self.assertEqual(self.solution.arrayPairSum([1, 4, 3, 2]), 4)

    def test_all_negative_min(self):
        l = [-10000, -1, 0, 1]
        self.assertEqual(self.solution.arrayPairSum(l), -10000)

        
if __name__ == "__main__":
    unittest.main()


