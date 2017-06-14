#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: range_addition_II_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import range_addition_II


class RangeAdditionIITestCase(unittest.TestCase):
    def setUp(self):
        self.solution = range_addition_II.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.maxCount(3, 3, [[2, 2], [3, 3]]), 4)

    def test_example(self):
        self.assertEqual(self.solution.maxCount(3, 3, [[1, 3], [3, 1]]), 1)
        
        
if __name__ == "__main__":
    unittest.main()


